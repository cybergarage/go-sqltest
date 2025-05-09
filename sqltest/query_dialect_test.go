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
	"fmt"
	"testing"
)

func TestNewQueryDataTypeFor(t *testing.T) {
	tests := []struct {
		dt       string
		to       QueryDialect
		expected string
	}{
		{"VARCHAR", QueryDialectMySQL, "VARCHAR"},
		{"VARCHAR", QueryDialectPostgreSQL, "VARCHAR"},
		{"INT", QueryDialectMySQL, "INT"},
		{"INT", QueryDialectPostgreSQL, "INTEGER"},
		{"FLOAT", QueryDialectMySQL, "FLOAT"},
		{"FLOAT", QueryDialectPostgreSQL, "REAL"},
		{"DOUBLE", QueryDialectPostgreSQL, "DOUBLE PRECISION"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s->%s", tt.dt, tt.to), func(t *testing.T) {
			got, err := NewQueryDataTypeFor(tt.dt, tt.to)
			if err != nil {
				t.Fatalf("NewQueryDataTypeFor(%q, %d) = _, %v; want _, nil", tt.dt, tt.to, err)
			}
			if got != tt.expected {
				t.Errorf("NewQueryDataTypeFor(%q, %d) = %q; want %q", tt.dt, tt.to, got, tt.expected)
			}
		})
	}
}
