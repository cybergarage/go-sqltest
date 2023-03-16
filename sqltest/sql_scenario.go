// Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
)

// SQLScenario represents a scenario.
type SQLScenario struct {
	*SQLTestFile
	Filename string
	Queries  []string
	Results  []*SQLResponse
}

// NewSQLScenario return a scenario instance.
func NewSQLScenario() *SQLScenario {
	file := &SQLScenario{
		SQLTestFile: NewSQLTestFile(),
	}
	return file
}

// NewSQLScenarioWithFile return a scenario instance for the specified test scenario file.
func NewSQLScenarioWithFile(filename string) (*SQLScenario, error) {
	scn := NewSQLScenario()
	err := scn.LoadFile(filename)
	return scn, err
}

// NewSQLScenarioWithBytes return a scenario instance for the specified test scenario bytes.
func NewSQLScenarioWithBytes(name string, b []byte) (*SQLScenario, error) {
	scn := NewSQLScenario()
	err := scn.ParseBytes(name, b)
	return scn, err
}

// Name returns the loaded scenario file name.
func (scn *SQLScenario) Name() string {
	return scn.Filename
}

// IsValid checks whether the loaded scenario is available.
func (scn *SQLScenario) IsValid() error {
	if len(scn.Queries) != len(scn.Results) {
		return fmt.Errorf(errorInvalidScenarioCases, len(scn.Queries), len(scn.Results))
	}
	return nil
}

// LoadFile loads the specified scenario.
func (scn *SQLScenario) LoadFile(filename string) error {
	lines, err := scn.SQLTestFile.LoadFile(filename)
	if err != nil {
		return err
	}

	err = scn.ParseLineStrings(lines)
	if err != nil {
		return fmt.Errorf("%s : %w", filename, err)
	}

	scn.Filename = filename

	return nil
}

// ParseBytes parses the specified scenario bytes.
func (scn *SQLScenario) ParseBytes(name string, b []byte) error {
	lines := strings.Split(string(b), "\r\n")
	err := scn.ParseLineStrings(lines)
	if err != nil {
		return fmt.Errorf("%s : %w", name, err)
	}
	scn.Filename = name
	return nil
}

// ParseLineStrings parses the specified scenario line strings.
func (scn *SQLScenario) ParseLineStrings(lines []string) error {
	var queryStr, resultStr string
	scn.Queries = make([]string, 0)
	scn.Results = make([]*SQLResponse, 0)

	appendQuery := func() {
		if len(queryStr) <= 0 {
			return
		}
		scn.Queries = append(scn.Queries, strings.TrimSpace(queryStr))
		queryStr = ""
	}
	appendResult := func() error {
		if len(resultStr) <= 0 {
			return nil
		}
		result, err := NewSQLResponseWithString(strings.TrimSpace(resultStr))
		if err != nil {
			return err
		}
		scn.Results = append(scn.Results, result)
		resultStr = ""
		return nil
	}

	inJSON := false
	for _, line := range lines {
		if inJSON {
			if strings.HasPrefix(line, "}") {
				resultStr += line
				err := appendResult()
				if err != nil {
					return err
				}
				inJSON = false
				continue
			}
			resultStr += " " + strings.TrimSpace(line)
		} else {
			if strings.HasPrefix(line, "{") {
				appendQuery()
				resultStr = line
				inJSON = true
				continue
			}
			queryStr += " " + strings.TrimSpace(line)
		}
	}

	return scn.IsValid()
}

// String returns the string representation.
func (scn *SQLScenario) String() string {
	var str string
	nResults := len(scn.Results)
	for n, query := range scn.Queries {
		str += query + "\n"
		if n < nResults {
			str += scn.Results[n].String() + "\n"
		}
	}
	return str
}
