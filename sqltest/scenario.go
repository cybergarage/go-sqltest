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
	"os"
	"strings"
)

// Line represents a line of a scenario test file.
type Line = string

// ScenarioOption represents an option function for a scenario.
type ScenarioOption func(*Scenario) error

// ScenarioStepHandler represents a scenario step handler.
type ScenarioStepHandler func(*Scenario, int, *Query, error)

// WithScenarioFile returns a scenario option to load the specified scenario file.
func WithScenarioFile(filename string) ScenarioOption {
	return func(scn *Scenario) error {
		return scn.LoadFile(filename)
	}
}

// WithScenarioBytes returns a scenario option to load the specified scenario bytes.
func WithScenarioBytes(name string, b []byte) ScenarioOption {
	return func(scn *Scenario) error {
		return scn.ParseBytes(name, b)
	}
}

// Scenario represents a scenario.
type Scenario struct {
	filename string
	queries  []string
	contents []*QueryContext
	cases    []*ScenarioCase
}

// NewScenario return a scenario instance.
func NewScenario() *Scenario {
	file := &Scenario{
		filename: "",
		queries:  []string{},
		contents: []*QueryContext{},
		cases:    []*ScenarioCase{},
	}
	return file
}

// NewScenarioWith returns a scenario instance with the specified options.
func NewScenarioWith(opts ...ScenarioOption) (*Scenario, error) {
	scn := NewScenario()
	for _, opt := range opts {
		err := opt(scn)
		if err != nil {
			return nil, err
		}
	}
	return scn, nil
}

// NewScenarioWithFile return a scenario instance for the specified test scenario file.
func NewScenarioWithFile(filename string) (*Scenario, error) {
	return NewScenarioWith(WithScenarioFile(filename))
}

// NewScenarioWithBytes return a scenario instance for the specified test scenario bytes.
func NewScenarioWithBytes(name string, b []byte) (*Scenario, error) {
	return NewScenarioWith(WithScenarioBytes(name, b))
}

// Name returns the loaded scenario file name.
func (scn *Scenario) Name() string {
	return scn.filename
}

// Queries returns the loaded scenario queries.
func (scn *Scenario) Queries() []string {
	return scn.queries
}

// Bindings returns the loaded scenario bindings.
func (scn *Scenario) Bindings() []QueryBindings {
	bindings := make([]QueryBindings, 0)
	for _, content := range scn.contents {
		v, ok := content.Bindings()
		if !ok {
			bindings = append(bindings, []any{}) // empty bindings
			continue
		}
		bindings = append(bindings, v)
	}
	return bindings
}

// Cases returns the loaded scenario cases.
func (scn *Scenario) Cases() []*ScenarioCase {
	return scn.cases
}

// ExpectedRows returns the loaded scenario expected rows.
func (scn *Scenario) ExpectedRows() ([]QueryRows, error) {
	rows := make([]QueryRows, 0)
	for _, content := range scn.contents {
		v, ok := content.Rows()
		if !ok {
			v = []any{} // empty rows
		}
		rows = append(rows, v)
	}
	return rows, nil
}

// LoadFile loads the specified scenario.
func (scn *Scenario) LoadFile(filename string) error {
	lines, err := scn.loadFileLines(filename)
	if err != nil {
		return err
	}

	err = scn.ParseLineStrings(lines)
	if err != nil {
		return fmt.Errorf("%s : %w", filename, err)
	}

	scn.filename = filename

	return nil
}

func (scn *Scenario) loadFileLines(filename string) ([]Line, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return scn.parseByteLines(fileBytes)
}

// ParseBytes parses the specified scenario bytes.
func (scn *Scenario) ParseBytes(name string, b []byte) error {
	lines, err := scn.parseByteLines(b)
	if err != nil {
		return err
	}

	err = scn.ParseLineStrings(lines)
	if err != nil {
		return fmt.Errorf("%s : %w", name, err)
	}

	scn.filename = name
	return nil
}

func (scn *Scenario) parseByteLines(fileBytes []byte) ([]Line, error) {
	lines := make([]Line, 0)
	for _, line := range strings.Split(string(fileBytes), "\n") {
		// Skip blank or comment lines
		if len(line) == 0 || strings.HasPrefix(line, "-") {
			continue
		}
		lines = append(lines, line)
	}
	return lines, nil
}

// ParseLineStrings parses the specified scenario line strings.
func (scn *Scenario) ParseLineStrings(lines []string) error {
	var queryStr, resultStr string
	var err error

	scn.queries = make([]string, 0)
	scn.contents = make([]*QueryContext, 0)
	scn.cases = make([]*ScenarioCase, 0)

	// Read all lines.

	appendQuery := func() {
		if len(queryStr) == 0 {
			return
		}
		scn.queries = append(scn.queries, strings.TrimSpace(queryStr))
		queryStr = ""
	}

	appendContext := func() error {
		if len(resultStr) == 0 {
			return nil
		}
		result, err := NewQueryContextWithString(strings.TrimSpace(resultStr))
		if err != nil {
			return err
		}
		scn.contents = append(scn.contents, result)
		resultStr = ""
		return nil
	}

	inJSON := false
	for n, line := range lines {
		if inJSON {
			if strings.HasPrefix(line, "}") {
				resultStr += line
				err := appendContext()
				if err != nil {
					return fmt.Errorf("line [%d] : %w (%v)", n, err, line)
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

	// Separate to cases.

	generateCases := func(scn *Scenario) ([]*ScenarioCase, error) {
		scnCases := make([]*ScenarioCase, 0)
		queries := scn.queries
		contents := scn.contents
		if len(queries) != len(contents) {
			return nil, fmt.Errorf(errorInvalidScenarioCases, len(queries), len(contents))
		}
		for n, query := range queries {
			content := contents[n]
			bindings, ok := content.Bindings()
			if !ok {
				bindings = []any{} // empty bindings
			}
			rows, ok := content.Rows()
			if !ok {
				rows = []any{} // empty rows
			}
			scnCase := NewScenarioCaseWith(
				WithScenarioCaseQuery(NewQueryWith(query, bindings)),
				WithScenarioCaseRows(rows),
			)
			scnCases = append(scnCases, scnCase)
		}
		return scnCases, nil
	}

	scn.cases, err = generateCases(scn)
	if err != nil {
		return err
	}

	return nil
}

// String returns the string representation.
func (scn *Scenario) String() string {
	var str string
	nResults := len(scn.contents)
	for n, query := range scn.queries {
		str += query + "\n"
		if n < nResults {
			str += scn.contents[n].String() + "\n"
		}
	}
	return str
}
