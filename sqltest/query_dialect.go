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
	_ "embed"
	"strings"
)

//go:embed data/query_data_types.csv
var queryDataTypeBytes []byte

var queryDataTypes []string
var queryDataTypeMap map[string]map[QueryDialect]string

// QueryDialect represents the SQL dialect used in a query.
type QueryDialect int

const (
	// QueryDialectNone represents no specific SQL dialect.
	QueryDialectNone QueryDialect = iota
	// QueryDialectMySQL represents the MySQL SQL dialect.
	QueryDialectMySQL
	// QueryDialectSQLite represents the SQLite SQL dialect.
	QueryDialectPostgreSQL
)

func init() {
	queryDataTypeMap = make(map[string]map[QueryDialect]string)

	csvLines := strings.Split(string(queryDataTypeBytes), "\n")
	headers := strings.Split(csvLines[0], ",")
	for _, line := range csvLines[1:] {
		fields := strings.Split(line, ",")
		dataType := fields[0]
		if strings.TrimSpace(dataType) == "" {
			continue
		}
		queryDataTypes = append(queryDataTypes, dataType)
		for i := 1; i < len(fields); i++ {
			if strings.TrimSpace(fields[i]) == "" {
				continue
			}
			dialect := headers[i]
			mappedType := fields[i]
			var dialectEnum QueryDialect
			switch dialect {
			case "MySQL":
				dialectEnum = QueryDialectMySQL
			case "PostgreSQL":
				dialectEnum = QueryDialectPostgreSQL
			default:
				continue
			}
			if _, exists := queryDataTypeMap[dataType]; !exists {
				queryDataTypeMap[dataType] = make(map[QueryDialect]string)
			}
			queryDataTypeMap[dataType][dialectEnum] = mappedType
		}
	}
}

// ListQueryDataTypes returns a list of supported query data types.
func ListQueryDataTypes() []string {
	return queryDataTypes
}

// NewQueryDialect returns a new QueryDialect instance.
func NewQueryDataTypeFor(dt string, to QueryDialect) (string, error) {
	dt = strings.ToUpper(strings.TrimSpace(dt))
	switch to {
	case QueryDialectMySQL:
		if mappedType, exists := queryDataTypeMap[dt][QueryDialectMySQL]; exists {
			return mappedType, nil
		}
		return dt, nil
	case QueryDialectPostgreSQL:
		if mappedType, exists := queryDataTypeMap[dt][QueryDialectPostgreSQL]; exists {
			return mappedType, nil
		}
		return dt, nil
	default:
		return dt, nil
	}
}

// String returns the string representation of the QueryDialect.
func (to QueryDialect) String() string {
	switch to {
	case QueryDialectNone:
		return "none"
	case QueryDialectMySQL:
		return "mysql"
	case QueryDialectPostgreSQL:
		return "postgresql"
	default:
		return "unknown"
	}
}
