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

// ScenarioCaseOption represents an option function for a scenario case.
type ScenarioCaseOption func(*ScenarioCase) error

// ScenarioCase represents a scenario case.
type ScenarioCase struct {
	query    string
	bindings []any
	rows     QueryRows
}

// WithScenarioCaseQuery returns a scenario case option to set a query.
func WithScenarioCaseQuery(query string) ScenarioCaseOption {
	return func(sc *ScenarioCase) error {
		sc.query = query
		return nil
	}
}

// WithScenarioCaseBindings returns a scenario case option to set bindings.
func WithScenarioCaseBindings(bindings []any) ScenarioCaseOption {
	return func(sc *ScenarioCase) error {
		sc.bindings = bindings
		return nil
	}
}

// WithScenarioCaseRows returns a scenario case option to set rows.
func WithScenarioCaseRows(rows QueryRows) ScenarioCaseOption {
	return func(sc *ScenarioCase) error {
		sc.rows = rows
		return nil
	}
}

// NewScenarioCase returns a scenario case instance.
func NewScenarioCaseWith(opts ...ScenarioCaseOption) *ScenarioCase {
	sc := &ScenarioCase{
		query:    "",
		bindings: []any{},
		rows:     QueryRows{},
	}
	for _, opt := range opts {
		opt(sc)
	}
	return sc
}

// Query returns a query of the scenario case.
func (sc *ScenarioCase) Query() string {
	return sc.query
}

// Bindings returns bindings of the scenario case.
func (sc *ScenarioCase) Bindings() []any {
	return sc.bindings
}

// Rows returns rows of the scenario case.
func (sc *ScenarioCase) Rows() QueryRows {
	return sc.rows
}
