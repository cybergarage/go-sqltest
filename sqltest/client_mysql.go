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
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultMysqlPort = 3306
)

// MySQLClient represents a client for MySQL server.
type MySQLClient struct {
	*Config
	db *sql.DB
}

// NewMySQLClient returns a client instance.
func NewMySQLClient() Client {
	client := &MySQLClient{
		Config: NewDefaultConfig(),
		db:     nil,
	}
	client.SetPort(defaultMysqlPort)
	return client
}

// Open opens a database specified by the internal configuration.
func (client *MySQLClient) Open() error {
	dsName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		client.User,
		client.Password,
		client.Host,
		client.Port,
		client.Database)
	db, err := sql.Open("mysql", dsName)
	if err != nil {
		return err
	}

	// See: https://github.com/go-sql-driver/mysql
	// Important settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	client.db = db
	return nil
}

// Close closes opens a database specified by the internal configuration.
func (client *MySQLClient) Close() error {
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
func (client *MySQLClient) Ping() error {
	if client.db == nil {
		err := client.Open()
		if err != nil {
			return err
		}
	}
	return client.db.Ping()
}

// Query executes a query that returns rows.
func (client *MySQLClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if client.db == nil {
		err := client.Open()
		if err != nil {
			return nil, err
		}
	}
	return client.db.Query(query, args...)
}

// CreateDatabase creates a specified database.
func (client *MySQLClient) CreateDatabase(name string) error {
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
func (client *MySQLClient) DropDatabase(name string) error {
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
func (client *MySQLClient) Use(name string) error {
	client.Database = name
	return nil
}
