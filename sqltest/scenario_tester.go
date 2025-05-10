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

// ScenarioTesterOption represents an option function for a scenario tester.
type ScenarioTesterOption func(*ScenarioTester) error

// ScenarioTester represents a scenario tester.
type ScenarioTester struct {
	scenario     *Scenario
	client       Client
	stepHandler  ScenarioStepHandler
	queryDialect QueryDialect
}

// WithScenarioTesterClient returns a scenario tester option to set a client.
func WithScenarioTesterClient(client Client) ScenarioTesterOption {
	return func(tester *ScenarioTester) error {
		tester.SetClient(client)
		return nil
	}
}

// WithScenarioTesterClient returns a scenario tester option to set a client.
func WithScenarioTesterFile(filename string) ScenarioTesterOption {
	return func(tester *ScenarioTester) error {
		return tester.LoadFile(filename)
	}
}

// WithScenarioTesterBytes returns a scenario tester option to set a client.
func WithScenarioTesterBytes(name string, b []byte) ScenarioTesterOption {
	return func(tester *ScenarioTester) error {
		return tester.ParseBytes(name, b)
	}
}

// WithScenarioTesterStepHandler returns a scenario tester option to set a step handler.
func WithScenarioTesterStepHandler(handler ScenarioStepHandler) ScenarioTesterOption {
	return func(tester *ScenarioTester) error {
		tester.SetStepHandler(handler)
		return nil
	}
}

// WithScenarioTesterQueryDialect returns a scenario tester option to set a query dialect.
func WithScenarioTesterQueryDialect(dialect QueryDialect) ScenarioTesterOption {
	return func(tester *ScenarioTester) error {
		tester.queryDialect = dialect
		return nil
	}
}

// NewScenarioTester returns a scenario tester instance.
func NewScenarioTester() *ScenarioTester {
	tester := &ScenarioTester{
		scenario:     nil,
		client:       nil,
		stepHandler:  nil,
		queryDialect: QueryDialectNone,
	}
	return tester
}

// NewScenarioTesterWith returns a scenario tester instance with the specified options.
func NewScenarioTesterWith(options ...ScenarioTesterOption) (*ScenarioTester, error) {
	tester := NewScenarioTester()
	for _, option := range options {
		err := option(tester)
		if err != nil {
			return nil, err
		}
	}
	return tester, nil
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
func (tester *ScenarioTester) SetClient(c Client) {
	tester.client = c
}

// SetStepHandler sets a step handler for testing.
func (tester *ScenarioTester) SetStepHandler(handler ScenarioStepHandler) {
	tester.stepHandler = handler
}

// Name returns the loaded senario name.
func (tester *ScenarioTester) Name() string {
	return tester.scenario.Name()
}

// LoadFile loads a specified scenario test file.
func (tester *ScenarioTester) LoadFile(filename string) error {
	tester.scenario = NewScenario()
	return tester.scenario.LoadFile(filename)
}

// ParseBytes loads a specified scenario test bytes.
func (tester *ScenarioTester) ParseBytes(name string, b []byte) error {
	tester.scenario = NewScenario()
	return tester.scenario.ParseBytes(name, b)
}

// LoadFileWithBasename loads a scenario test file which has specified basename.
func (tester *ScenarioTester) LoadFileWithBasename(basename string) error {
	return tester.LoadFile(basename + "." + ScenarioFileExt)
}

// Scenario returns the loaded scenario.
func (tester *ScenarioTester) Scenario() *Scenario {
	return tester.scenario
}

// Run runs a loaded scenario test.
func (tester *ScenarioTester) Run() error {
	scenario := tester.Scenario()
	if scenario == nil {
		return nil
	}

	client := tester.client
	if client == nil {
		return errors.New(errorClientNotFound)
	}

	stepHandler := func(n int, query *Query, err error) error {
		if tester.stepHandler != nil {
			tester.stepHandler(scenario, n, query, err)
		}
		return err
	}

	errTraceMsg := func(n int) string {
		queries := scenario.Queries()
		errTraceMsg := tester.Name() + "\n"
		for i := 0; i < n; i++ {
			errTraceMsg += fmt.Sprintf(goodQueryPrefix, i, queries[i])
			errTraceMsg += "\n"
		}
		return errTraceMsg
	}

	testCases := scenario.Cases()
	for n, testCase := range testCases {
		query := testCase.Query()
		dialectQuery := query.DialectString(tester.queryDialect)
		log.Infof("[%d] %s", n, dialectQuery)
		rows, err := client.Query(dialectQuery, query.Arguments()...)
		if err != nil {
			errTraceMsg := errTraceMsg(n)
			errTraceMsg += fmt.Sprintf(errorQueryPrefix, n, dialectQuery)
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

		expectedRows := testCase.Rows()
		if len(rsRows) != len(expectedRows) {
			return stepHandler(n, query, fmt.Errorf("%s"+errorJSONResponseUnmatchedRowCount, errTraceMsg(n), n, query, rsRows, expectedRows))
		}

		for _, row := range rsRows {
			err = testCase.HasRow(row)
			if err != nil {
				return stepHandler(n, query, fmt.Errorf("%s"+errorQueryPrefix+"%w", errTraceMsg(n), n, query, err))
			}
		}

		stepHandler(n, query, nil)
	}

	return nil
}
