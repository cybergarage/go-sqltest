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
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

const (
	QueryResponseRowsKey = "rows"
)

// QueryResponseData defines a JSON response data type.
type QueryResponseData = map[string]interface{}

// QueryResponseRows defines a JSON response rows type.
type QueryResponseRows = []interface{}

// QueryResponseRow defines a JSON response row type.
type QueryResponseRow = map[string]interface{}

// QueryResponse represents a response of a query.
type QueryResponse struct {
	Data QueryResponseData
}

// NewQueryResponse returns a response instance.
func NewQueryResponse() *QueryResponse {
	res := &QueryResponse{}
	return res
}

// NewQueryResponseWithString returns a response instance of the specified JSON response.
func NewQueryResponseWithString(json string) (*QueryResponse, error) {
	res := NewQueryResponse()
	err := res.ParseString(json)
	return res, err
}

// ParseString parses a specified string response as a JSON data.
func (res *QueryResponse) ParseString(jsonStr string) error {
	var rootObj interface{}
	err := json.Unmarshal([]byte(jsonStr), &rootObj)
	if err != nil {
		return err
	}

	var ok bool
	res.Data, ok = rootObj.(QueryResponseData)
	if !ok {
		return fmt.Errorf(errorInvalidJSONResponse, rootObj)
	}

	return nil
}

// Rows returns response rows with true when the response has any rows, otherwise nil and false.
func (res *QueryResponse) Rows() (QueryResponseRows, error) {
	if res.Data == nil {
		return nil, fmt.Errorf(errorJSONResponseNotFound)
	}

	rowsData, ok := res.Data[QueryResponseRowsKey]
	if !ok {
		return nil, fmt.Errorf(errorJSONResponseRowsNotFound, res.Data, QueryResponseRowsKey)
	}

	rows, ok := rowsData.(QueryResponseRows)
	if !ok {
		return nil, fmt.Errorf(errorJSONResponseRowsNotFound, res.Data, QueryResponseRowsKey)
	}

	return rows, nil
}

// HasRow returns true when the response has a specified row, otherwise false.
// nolint: gocyclo
func (res *QueryResponse) HasRow(row interface{}) error {
	rowMap, ok := row.(QueryResponseRow)
	if !ok {
		return fmt.Errorf(errorJSONResponseHasNoRow, row, rowMap)
	}

	resRows, err := res.Rows()
	if err != nil {
		return err
	}

	var deepEqual func(iv1 interface{}, iv2 interface{}) bool
	deepEqual = func(iv1 interface{}, iv2 interface{}) bool {
		if reflect.DeepEqual(iv1, iv2) {
			return true
		}

		trimString := func(s string) string {
			return strings.Trim(s, "\"'")
		}

		// NOTE: DeepEqual checks the types and values strictly.
		// Therefore, support other types if needed.
		// log.Debugf("deepEqual: %v (%T) != %v (%T)", iv1, iv1, iv2, iv2)

		uint8sToString := func(ui8s []uint8) string {
			bytesLen := len(ui8s)
			bytes := make([]byte, bytesLen)
			for n := 0; n < bytesLen; n++ {
				bytes[n] = ui8s[n]
			}
			return string(bytes)
		}

		integerStringEqual := func(s1, s2 string) bool {
			var i1 int
			err := safecast.ToInt(s1, &i1)
			if err != nil {
				return false
			}
			var i2 int
			err = safecast.ToInt(s2, &i2)
			if err != nil {
				return false
			}
			return i1 == i2
		}

		realStringEqual := func(s1, s2 string) bool {
			var f1 float64
			err := safecast.ToFloat64(s1, &f1)
			if err != nil {
				return false
			}
			var f2 float64
			err = safecast.ToFloat64(s2, &f2)
			if err != nil {
				return false
			}
			return f1 == f2
		}

		almostEqual := func(a, b float64) bool {
			diff := math.Abs(a - b)
			avg := (math.Abs(a) + math.Abs(b)) / 2
			return diff/avg <= 1e-6
		}

		switch v1 := iv1.(type) {
		case string:
			var v2 string
			err := safecast.ToString(iv2, &v2)
			if err != nil {
				return false
			}
			sv1 := trimString(v1)
			sv2 := trimString(v2)
			if sv1 == sv2 {
				return true
			}
			if integerStringEqual(sv1, sv2) {
				return true
			}
			if realStringEqual(sv1, sv2) {
				return true
			}
		case int:
			var v2 int
			err := safecast.ToInt(iv2, &v2)
			if err != nil {
				return false
			}
			if v1 == v2 {
				return true
			}
		case float64:
			var v2 float64
			err := safecast.ToFloat64(iv2, &v2)
			if err != nil {
				return false
			}
			if v1 == v2 {
				return true
			}
			if almostEqual(v1, v2) {
				return true
			}
		case []uint8:
			sv1 := uint8sToString(v1)
			return deepEqual(sv1, iv2)
		case time.Time:
			tv1 := v1.Format(time.DateTime)
			tv2 := fmt.Sprintf("%v", iv2)
			return deepEqual(tv1, tv2)
		}

		sv1 := fmt.Sprintf("%v", iv1)
		sv2 := fmt.Sprintf("%v", iv2)
		return sv1 == sv2
	}

	for _, resRow := range resRows {
		resMap, ok := resRow.(QueryResponseRow)
		if !ok {
			continue
		}

		if reflect.DeepEqual(rowMap, resMap) {
			return nil
		}

		hasAllColumn := true
		for rowKey, rowData := range rowMap {
			resData, ok := resMap[rowKey]
			if !ok {
				hasAllColumn = false
				break
			}

			if !deepEqual(resData, rowData) && !deepEqual(rowData, resData) {
				hasAllColumn = false
				break
			}
		}

		if hasAllColumn {
			return nil
		}
	}

	return fmt.Errorf(errorJSONResponseHasNoRow, row, res.Data)
}

// String returns the string representation.
func (res *QueryResponse) String() string {
	return fmt.Sprintf("%v", res.Data)
}
