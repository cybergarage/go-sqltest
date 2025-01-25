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
type ScenarioStepHandler func(*Scenario, int, string, error)

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
	Filename  string
	Queries   []string
	Expecteds []*QueryContext
}

// NewScenario return a scenario instance.
func NewScenario() *Scenario {
	file := &Scenario{
		Filename:  "",
		Queries:   []string{},
		Expecteds: []*QueryContext{},
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
	return scn.Filename
}

// IsValid checks whether the loaded scenario is available.
func (scn *Scenario) IsValid() error {
	if len(scn.Queries) != len(scn.Expecteds) {
		return fmt.Errorf(errorInvalidScenarioCases, len(scn.Queries), len(scn.Expecteds))
	}
	return nil
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

	scn.Filename = filename

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

	scn.Filename = name
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
	scn.Queries = make([]string, 0)
	scn.Expecteds = make([]*QueryContext, 0)

	appendQuery := func() {
		if len(queryStr) == 0 {
			return
		}
		scn.Queries = append(scn.Queries, strings.TrimSpace(queryStr))
		queryStr = ""
	}
	appendResult := func() error {
		if len(resultStr) == 0 {
			return nil
		}
		result, err := NewQueryContextWithString(strings.TrimSpace(resultStr))
		if err != nil {
			return err
		}
		scn.Expecteds = append(scn.Expecteds, result)
		resultStr = ""
		return nil
	}

	inJSON := false
	for n, line := range lines {
		if inJSON {
			if strings.HasPrefix(line, "}") {
				resultStr += line
				err := appendResult()
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

	return scn.IsValid()
}

// String returns the string representation.
func (scn *Scenario) String() string {
	var str string
	nResults := len(scn.Expecteds)
	for n, query := range scn.Queries {
		str += query + "\n"
		if n < nResults {
			str += scn.Expecteds[n].String() + "\n"
		}
	}
	return str
}
