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
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cybergarage/go-sqltest/sqltest"
)

func TestQueryContext(t *testing.T) {
	testJSONStrs := []string{
		"{}",
		"{ \"rows\" : [ { \"k\" : \"0\", \"v1\" : \"0\", \"v2\" : \"0\" } ]}",
		"{ \"rows\" : [ { \"k\" : \"0\", \"v1\" : \"0\", \"v2\" : \"0\" }, { \"k\" : \"1\", \"v1\" : \"1\", \"v2\" : \"1\" } ]}",
	}

	res := sqltest.NewQueryContext()
	for _, jsonStr := range testJSONStrs {
		err := res.ParseString(jsonStr)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestQueryContextBindings(t *testing.T) {
	testBindings := sqltest.QueryBindings{
		"0",
		"1",
		"2",
	}
	testResData := sqltest.QueryContextData{
		sqltest.QueryContextBindingsKey: testBindings,
	}

	jsonStr, err := json.Marshal(testResData)
	if err != nil {
		t.Error(err)
		return
	}

	res := sqltest.NewQueryContext()
	err = res.ParseString(string(jsonStr))
	if err != nil {
		t.Error(err)
		return
	}

	bindings, ok := res.Bindings()
	if !ok {
		t.Error("Failed to get bindings")
		return
	}

	if !reflect.DeepEqual(bindings, testBindings) {
		t.Error("Failed to get bindings")
		return
	}
}

func TestQueryContextRows(t *testing.T) {
	testRows := sqltest.QueryRows{
		sqltest.QueryResponseRow{
			"k":  0,
			"v1": 0,
			"v2": 0,
		},
		sqltest.QueryResponseRow{
			"k":  1,
			"v1": 1,
			"v2": 1,
		},
	}
	testResData := sqltest.QueryContextData{
		sqltest.QueryContextRowsKey: testRows,
	}

	jsonStr, err := json.Marshal(testResData)
	if err != nil {
		t.Error(err)
		return
	}

	res := sqltest.NewQueryContext()
	err = res.ParseString(string(jsonStr))
	if err != nil {
		t.Error(err)
		return
	}

	_, ok := res.Rows()
	if !ok {
		t.Error("Failed to get rows")
		return
	}
}
