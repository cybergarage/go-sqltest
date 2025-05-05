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
	"fmt"
)

type sqlClient struct {
	*Config
	db *sql.DB
}

func newSQLClient() *sqlClient {
	client := &sqlClient{
		Config: NewDefaultConfig(),
		db:     nil,
	}
	return client
}

// Open opens a database specified by the internal configuration.
// nolint: gosec, exhaustruct, staticcheck
func (client *sqlClient) Open() error {
	return ErrNotImplemented
}

// Close closes opens a database specified by the internal configuration.
func (client *sqlClient) Close() error {
	if client.db == nil {
		return nil
	}
	if err := client.db.Close(); err != nil {
		return err
	}
	client.db = nil
	return nil
}

// Ping pings the opened database.
func (client *sqlClient) Ping() error {
	if client.db == nil {
		err := client.Open()
		if err != nil {
			return err
		}
	}
	return client.db.Ping()
}

// Query executes a query that returns rows.
func (client *sqlClient) Query(query string, args ...any) (*sql.Rows, error) {
	if client.db == nil {
		err := client.Open()
		if err != nil {
			return nil, err
		}
	}
	if client.IsPreparedStatementEnabled() {
		return client.db.Query(query, args...)
	}
	stmt := NewStatement(query)
	query, err := stmt.Bind(args...)
	if err != nil {
		return nil, err
	}
	return client.db.Query(query)
}

// CreateDatabase creates a specified database.
func (client *sqlClient) CreateDatabase(name string) error {
	query := fmt.Sprintf("CREATE DATABASE %s", name)
	rows, err := client.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// DropDatabase dtops a specified database.
func (client *sqlClient) DropDatabase(name string) error {
	query := fmt.Sprintf("DROP DATABASE %s", name)
	rows, err := client.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// Use sets a target database.
func (client *sqlClient) Use(name string) error {
	client.Database = name
	return nil
}
