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

// Query represents a SQL query with its arguments.
type Query struct {
	query string
	args  []any
}

// NewQueryWith creates a new Query instance with the given SQL query and arguments.
func NewQueryWith(query string, args ...any) *Query {
	return &Query{
		query: query,
		args:  args,
	}
}

// String returns the SQL query string.
func (q *Query) String() string {
	return q.query
}

// DialectString returns the SQL query string for the specified dialect.
func (q *Query) DialectString(dialect QueryDialect) string {
	switch dialect {
	case QueryDialectMySQL:
		return q.query
	case QueryDialectPostgreSQL:
		return q.query
	default:
		return q.query
	}
}

// Aarguments returns the arguments for the SQL query.
func (q *Query) Arguments() []any {
	return q.args
}
