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
	"fmt"
	"strings"
	"time"

	"github.com/cybergarage/go-logger/log"
)

const (
	ScenarioFileExt = "qst"
)

// ScenarioRunner represents a scenario runner.
type ScenarioRunner struct {
	Scenario *Scenario
	client   Client
}

// NewScenarioRunner returns a scenario runner instance.
func NewScenarioRunner() *ScenarioRunner {
	runner := &ScenarioRunner{
		Scenario: nil,
		client:   nil,
	}
	return runner
}

// NewScenarioRunnerWithFile return a scenario test instance for the specified test scenario file.
func NewScenarioRunnerWithFile(filename string) (*ScenarioRunner, error) {
	runner := NewScenarioRunner()
	err := runner.LoadFile(filename)
	return runner, err
}

// NewScenarioRunnerWithBytes return a scenario test instance for the specified test scenario bytes.
func NewScenarioRunnerWithBytes(name string, b []byte) (*ScenarioRunner, error) {
	runner := NewScenarioRunner()
	err := runner.ParseBytes(name, b)
	return runner, err
}

// SetClient sets a client for testing.
func (runner *ScenarioRunner) SetClient(c Client) {
	runner.client = c
}

// Name returns the loaded senario name.
func (runner *ScenarioRunner) Name() string {
	return runner.Scenario.Name()
}

// LoadFile loads a specified scenario test file.
func (runner *ScenarioRunner) LoadFile(filename string) error {
	runner.Scenario = NewScenario()
	return runner.Scenario.LoadFile(filename)
}

// ParseBytes loads a specified scenario test bytes.
func (runner *ScenarioRunner) ParseBytes(name string, b []byte) error {
	runner.Scenario = NewScenario()
	return runner.Scenario.ParseBytes(name, b)
}

// LoadFileWithBasename loads a scenario test file which has specified basename.
func (runner *ScenarioRunner) LoadFileWithBasename(basename string) error {
	return runner.LoadFile(basename + "." + ScenarioFileExt)
}

// Run runs a loaded scenario test.
func (runner *ScenarioRunner) Run() error {
	scenario := runner.Scenario
	if scenario == nil {
		return nil
	}

	err := scenario.IsValid()
	if err != nil {
		return err
	}

	client := runner.client
	if client == nil {
		return fmt.Errorf(errorClientNotFound)
	}

	errTraceMsg := func(n int) string {
		errTraceMsg := runner.Name() + "\n"
		for i := 0; i < n; i++ {
			errTraceMsg += fmt.Sprintf(goodQueryPrefix, i, scenario.Queries[i])
			errTraceMsg += "\n"
		}
		return errTraceMsg
	}

	for n, query := range scenario.Queries {
		log.Infof("[%d] %s", n, query)
		rows, err := client.Query(query)
		if err != nil {
			errTraceMsg := errTraceMsg(n)
			errTraceMsg += fmt.Sprintf(errorQueryPrefix, n, scenario.Queries[n])
			errTraceMsg += "\n"
			return fmt.Errorf("%s%w", errTraceMsg, err)
		}
		err = rows.Err()
		if err != nil {
			return err
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			return err
		}
		columnCnt := len(columns)

		columnTypes, err := rows.ColumnTypes()
		if err != nil {
			return err
		}

		// NOTE: Run() supports only the following standard column types yet.
		values := make([]interface{}, columnCnt)
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

		rsRows := make([]interface{}, 0)
		for rows.Next() {
			err = rows.Scan(values...)
			if err != nil {
				return err
			}

			row := map[string]interface{}{}
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
				case *interface{}:
					row[columns[i]] = *v
				default:
					row[columns[i]] = values[i]
				}
			}

			rsRows = append(rsRows, row)
		}

		expectedRes := scenario.Expecteds[n]
		expectedRows, err := expectedRes.Rows()
		if err != nil {
			if len(rsRows) != 0 {
				return fmt.Errorf("%s"+errorJSONResponseHasUnexpectedRows, errTraceMsg(n), n, query, rsRows)
			}
		} else {
			if len(rsRows) != len(expectedRows) {
				return fmt.Errorf("%s"+errorJSONResponseUnmatchedRowCount, errTraceMsg(n), n, query, rsRows, expectedRows)
			}
		}

		for _, row := range rsRows {
			err = expectedRes.HasRow(row)
			if err != nil {
				return fmt.Errorf("%s"+errorQueryPrefix+"%w", errTraceMsg(n), n, query, err)
			}
		}
	}

	return nil
}
