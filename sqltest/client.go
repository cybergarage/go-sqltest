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
type Client interface { //nolint: interfacebloat
	// SetHost sets a host name.
	SetHost(host string)
	// SetPort sets a port number.
	SetPort(port int)
	// SetUser sets a user name.
	SetUser(user string)
	// SetPassword sets a password.
	SetPassword(passwd string)
	// SetDatabase sets a database name.
	SetDatabase(db string)
	// Open opens a database specified by the internal configuration.
	Open() error
	// Close closes the opened database.
	Close() error
	// Use uses a database.
	Use(name string) error
	// Ping pings the opened database.
	Ping() error
	// CreateDatabase creates a database.
	CreateDatabase(name string) error
	// DropDatabase drops a database.
	DropDatabase(name string) error
	// Query executes a query.
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
