// Copyright (C) 2019 The go-sqltest Authors. All rights reserved.
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
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	defaultPostgresPort = 5432
)

// PostgresClient represents a client for PostgreSQL server.
type PostgresClient struct {
	*Config
	db *sql.DB
}

// NewClient returns a client instance.
func NewPostgresClient() Client {
	client := &PostgresClient{
		Config: NewDefaultConfig(),
		db:     nil,
	}
	client.SetPort(defaultPostgresPort)
	return client
}

// Open opens a database specified by the internal configuration.
func (client *PostgresClient) Open() error {
	dsName := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", client.Host, client.Port, client.Database)
	//  user=%s password=%s
	db, err := sql.Open("postgres", dsName)
	if err != nil {
		return err
	}
	client.db = db
	return nil
}

// Close closes opens a database specified by the internal configuration.
func (client *PostgresClient) Close() error {
	if client.db == nil {
		return nil
	}
	if err := client.db.Close(); err != nil {
		return err
	}
	client.db = nil
	return nil
}

// Query executes a query that returns rows.
func (client *PostgresClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if client.db == nil {
		err := client.Open()
		if err != nil {
			return nil, err
		}
	}
	return client.db.Query(query, args...)
}

// CreateDatabase creates a specified database.
func (client *PostgresClient) CreateDatabase(name string) error {
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
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
func (client *PostgresClient) DropDatabase(name string) error {
	query := fmt.Sprintf("DROP DATABASE IF EXISTS %s", name)
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
func (client *PostgresClient) Use(name string) error {
	client.SetDatabase(name)
	return nil
}
