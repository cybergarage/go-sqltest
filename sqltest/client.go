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
	"database/sql"
)

// Client represents a client interface for SQL databases.
type Client interface {
	SetHost(host string)
	SetPort(port int)
	SetDatabase(db string)
	Open() error
	Close() error
	CreateDatabase(name string) error
	DropDatabase(name string) error
	Use(name string) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
