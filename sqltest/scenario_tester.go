// Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqltest

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cybergarage/go-logger/log"
)

const (
	ScenarioFileExt = "qst"
)

// ScenarioTesterOption represents an option function for a scenario runner.
type ScenarioTesterOption func(*ScenarioTester) error

// ScenarioTester represents a scenario runner.
type ScenarioTester struct {
	scenario    *Scenario
	client      Client
	stepHandler ScenarioStepHandler
}

// WithScenarioTesterClient returns a scenario tester option to set a client.
func WithScenarioTesterClient(client Client) ScenarioTesterOption {
	return func(runner *ScenarioTester) error {
		runner.SetClient(client)
		return nil
	}
}

// WithScenarioTesterClient returns a scenario tester option to set a client.
func WithScenarioTesterFile(filename string) ScenarioTesterOption {
	return func(runner *ScenarioTester) error {
		return runner.LoadFile(filename)
	}
}

// WithScenarioTesterBytes returns a scenario tester option to set a client.
func WithScenarioTesterBytes(name string, b []byte) ScenarioTesterOption {
	return func(runner *ScenarioTester) error {
		return runner.ParseBytes(name, b)
	}
}

// WithScenarioTesterStepHandler returns a scenario tester option to set a step handler.
func WithScenarioTesterStepHandler(handler ScenarioStepHandler) ScenarioTesterOption {
	return func(runner *ScenarioTester) error {
		runner.SetStepHandler(handler)
		return nil
	}
}

// NewScenarioTester returns a scenario tester instance.
func NewScenarioTester() *ScenarioTester {
	runner := &ScenarioTester{
		scenario:    nil,
		client:      nil,
		stepHandler: nil,
	}
	return runner
}

// NewScenarioTesterWith returns a scenario tester instance with the specified options.
func NewScenarioTesterWith(options ...ScenarioTesterOption) (*ScenarioTester, error) {
	runner := NewScenarioTester()
	for _, option := range options {
		err := option(runner)
		if err != nil {
			return nil, err
		}
	}
	return runner, nil
}

// NewScenarioTesterWithFile return a scenario test instance for the specified test scenario file.
func NewScenarioTesterWithFile(filename string) (*ScenarioTester, error) {
	return NewScenarioTesterWith(WithScenarioTesterFile(filename))
}

// NewScenarioTesterWithBytes return a scenario test instance for the specified test scenario bytes.
func NewScenarioTesterWithBytes(name string, b []byte) (*ScenarioTester, error) {
	return NewScenarioTesterWith(WithScenarioTesterBytes(name, b))
}

// SetClient sets a client for testing.
func (runner *ScenarioTester) SetClient(c Client) {
	runner.client = c
}

// SetStepHandler sets a step handler for testing.
func (runner *ScenarioTester) SetStepHandler(handler ScenarioStepHandler) {
	runner.stepHandler = handler
}

// Name returns the loaded senario name.
func (runner *ScenarioTester) Name() string {
	return runner.scenario.Name()
}

// LoadFile loads a specified scenario test file.
func (runner *ScenarioTester) LoadFile(filename string) error {
	runner.scenario = NewScenario()
	return runner.scenario.LoadFile(filename)
}

// ParseBytes loads a specified scenario test bytes.
func (runner *ScenarioTester) ParseBytes(name string, b []byte) error {
	runner.scenario = NewScenario()
	return runner.scenario.ParseBytes(name, b)
}

// LoadFileWithBasename loads a scenario test file which has specified basename.
func (runner *ScenarioTester) LoadFileWithBasename(basename string) error {
	return runner.LoadFile(basename + "." + ScenarioFileExt)
}

// Scenario returns the loaded scenario.
func (runner *ScenarioTester) Scenario() *Scenario {
	return runner.scenario
}

// Run runs a loaded scenario test.
func (runner *ScenarioTester) Run() error {
	scenario := runner.Scenario()
	if scenario == nil {
		return nil
	}

	queries := scenario.Queries()

	client := runner.client
	if client == nil {
		return errors.New(errorClientNotFound)
	}

	stepHandler := func(n int, query string, err error) error {
		if runner.stepHandler != nil {
			runner.stepHandler(scenario, n, query, err)
		}
		return err
	}

	errTraceMsg := func(n int) string {
		errTraceMsg := runner.Name() + "\n"
		for i := 0; i < n; i++ {
			errTraceMsg += fmt.Sprintf(goodQueryPrefix, i, queries[i])
			errTraceMsg += "\n"
		}
		return errTraceMsg
	}

	for n, query := range queries {
		log.Infof("[%d] %s", n, query)
		rows, err := client.Query(query)
		if err != nil {
			errTraceMsg := errTraceMsg(n)
			errTraceMsg += fmt.Sprintf(errorQueryPrefix, n, queries[n])
			errTraceMsg += "\n"
			return stepHandler(n, query, fmt.Errorf("%s%w", errTraceMsg, err))
		}
		err = rows.Err()
		if err != nil {
			return stepHandler(n, query, err)
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			return stepHandler(n, query, err)
		}
		columnCnt := len(columns)

		columnTypes, err := rows.ColumnTypes()
		if err != nil {
			return stepHandler(n, query, err)
		}

		// NOTE: Run() supports only the following standard column types yet.
		values := make([]any, columnCnt)
		for n, columnType := range columnTypes {
			s := strings.ToUpper(columnType.DatabaseTypeName())
			switch {
			case (0 <= strings.Index(s, "INT")):
				var v int
				values[n] = &v
			case strings.HasPrefix(s, "FLOAT") || strings.HasPrefix(s, "DOUBLE"):
				var v float64
				values[n] = &v
			case strings.HasPrefix(s, "TEXT") || (0 <= strings.Index(s, "VARCHAR")):
				var v string
				values[n] = &v
			case strings.HasPrefix(s, "BLOB") || strings.HasPrefix(s, "BINARY"):
				var v []byte
				values[n] = &v
			case strings.HasPrefix(s, "TIMESTAMP") || strings.HasPrefix(s, "DATETIME"):
				var v time.Time
				values[n] = &v
			default:
				var v any
				values[n] = &v
			}
		}

		rsRows := make([]any, 0)
		for rows.Next() {
			err = rows.Scan(values...)
			if err != nil {
				return stepHandler(n, query, err)
			}

			row := map[string]any{}
			for i := 0; i < columnCnt; i++ {
				switch v := values[i].(type) {
				case *int:
					row[columns[i]] = *v
				case *float64:
					row[columns[i]] = *v
				case *string:
					row[columns[i]] = *v
				case *time.Time:
					row[columns[i]] = *v
				case *any:
					row[columns[i]] = *v
				default:
					row[columns[i]] = values[i]
				}
			}

			rsRows = append(rsRows, row)
		}

		expectedRes := scenario.contents[n]
		expectedRows, ok := expectedRes.Rows()
		if !ok {
			if len(rsRows) != 0 {
				return stepHandler(n, query, fmt.Errorf("%s"+errorJSONResponseHasUnexpectedRows, errTraceMsg(n), n, query, rsRows))
			}
		} else {
			if len(rsRows) != len(expectedRows) {
				return stepHandler(n, query, fmt.Errorf("%s"+errorJSONResponseUnmatchedRowCount, errTraceMsg(n), n, query, rsRows, expectedRows))
			}
		}

		for _, row := range rsRows {
			err = expectedRes.HasRow(row)
			if err != nil {
				return stepHandler(n, query, fmt.Errorf("%s"+errorQueryPrefix+"%w", errTraceMsg(n), n, query, err))
			}
		}

		stepHandler(n, query, nil)
	}

	return nil
}
