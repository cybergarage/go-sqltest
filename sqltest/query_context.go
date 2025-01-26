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
	"encoding/json"
	"fmt"
)

const (
	QueryContextRowsKey     = "rows"
	QueryContextBindingsKey = "bindings"
)

// QueryContextData defines a JSON response data type.
type QueryContextData = map[string]any

// QueryRows defines a JSON response rows type.
type QueryRows = []any

// QueryBindings defines a JSON response bindings type.
type QueryBindings = []any

// QueryResponseRow defines a JSON response row type.
type QueryResponseRow = map[string]any

// QueryContext represents a response of a query.
type QueryContext struct {
	Data QueryContextData
}

// NewQueryContext returns a response instance.
func NewQueryContext() *QueryContext {
	res := &QueryContext{
		Data: QueryContextData{},
	}
	return res
}

// NewQueryContextWithString returns a response instance of the specified JSON response.
func NewQueryContextWithString(json string) (*QueryContext, error) {
	res := NewQueryContext()
	err := res.ParseString(json)
	return res, err
}

// ParseString parses a specified string response as a JSON data.
func (res *QueryContext) ParseString(jsonStr string) error {
	var rootObj any
	err := json.Unmarshal([]byte(jsonStr), &rootObj)
	if err != nil {
		return err
	}

	var ok bool
	res.Data, ok = rootObj.(QueryContextData)
	if !ok {
		return fmt.Errorf(errorInvalidJSONResponse, rootObj)
	}

	return nil
}

// Bindings returns response bindings.
func (res *QueryContext) Bindings() (QueryBindings, bool) {
	if res.Data == nil {
		return nil, false
	}

	bindingsData, ok := res.Data[QueryContextBindingsKey]
	if !ok {
		return nil, false
	}

	bindings, ok := bindingsData.(QueryBindings)
	if !ok {
		return nil, false
	}

	return bindings, true
}

// Rows returns response rows with true when the response has any rows, otherwise nil and false.
func (res *QueryContext) Rows() (QueryRows, bool) {
	if res.Data == nil {
		return nil, false
	}

	rowsData, ok := res.Data[QueryContextRowsKey]
	if !ok {
		return nil, false
	}

	rows, ok := rowsData.(QueryRows)
	if !ok {
		return nil, false
	}

	return rows, true
}

// String returns the string representation.
func (res *QueryContext) String() string {
	return fmt.Sprintf("%v", res.Data)
}
