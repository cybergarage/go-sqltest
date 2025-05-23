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

// ClientConfig represents a client configuration interface.
type ClientConfig interface {
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
}

// ClientTLSConfig represents a client TLS configuration interface.
type ClientTLSConfig interface {
	// TLSEnabled returns true if TLS is enabled.
	TLSEnabled() bool
	// SetClientKeyFile sets a SSL client key file.
	SetClientKeyFile(file string)
	// SetClientCertFile sets a SSL client certificate file.
	SetClientCertFile(file string)
	// SetRootCertFile sets a SSL root certificate file.
	SetRootCertFile(file string)
}

// ClientQueryConfig represents a client query configuration interface.
type ClientQueryConfig interface {
	// SetPreparedStatementEnabled sets the prepared statement enabled flag.
	SetPreparedStatementEnabled(enabled bool)
	// IsPreparedStatementEnabled returns true if prepared statements are enabled.
	IsPreparedStatementEnabled() bool
}

// Client represents a client interface for SQL databases.
type Client interface {
	// ClientConfig represents a client configuration interface.
	ClientConfig
	// ClientTLSConfig represents a client TLS configuration interface.
	ClientTLSConfig
	// ClientQueryConfig represents a client query configuration interface.
	ClientQueryConfig
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
	Query(query string, args ...any) (*sql.Rows, error)
}
