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
	"strings"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

// Statement represents a prepared statement.
type Statement struct {
	stmt string
}

// NewStatement returns a statement instance.
func NewStatement(stmt string) *Statement {
	return &Statement{
		stmt: stmt,
	}
}

// Bind binds the specified arguments to the statement.
func (s *Statement) Bind(args ...any) (string, error) {
	query := s.stmt
	for _, arg := range args {
		var str string
		if err := safecast.ToString(arg, &str); err != nil {
			return "", err
		}
		switch arg.(type) {
		case string, time.Time:
			str = "'" + str + "'"
		}
		query = strings.Replace(query, "?", str, 1)
	}
	return query, nil
}

// String returns the statement string.
func (s *Statement) String() string {
	return s.stmt
}
