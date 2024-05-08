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
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// PqClient represents a client for PostgreSQL server.
type PqClient struct {
	*Config
	db *sql.DB
}

// NewPqClient returns a new lib/pq client.
func NewPqClient() *PqClient {
	client := &PqClient{
		Config: NewDefaultConfig(),
		db:     nil,
	}
	client.SetPort(defaultPostgresPort)
	return client
}

// Open opens a database specified by the internal configuration.
func (client *PqClient) Open() error {
	dsParams := []string{
		"host=" + client.Host,
		"port=" + strconv.Itoa(client.Port),
		"sslmode=disable",
	}
	if 0 < len(client.User) {
		dsParams = append(dsParams, "user="+client.User)
	}
	if 0 < len(client.Password) {
		dsParams = append(dsParams, "password="+client.Password)
	}
	if 0 < len(client.Database) {
		dsParams = append(dsParams, "dbname="+client.Database)
	}
	if client.TLSConfig != nil {
		dsParams = append(dsParams, "sslmode=require")
		if 0 < len(client.TLSConfig.CertFile) {
			dsParams = append(dsParams, "sslcert="+client.TLSConfig.CertFile)
		}
		if 0 < len(client.TLSConfig.KeyFile) {
			dsParams = append(dsParams, "sslkey="+client.TLSConfig.KeyFile)
		}
		if 0 < len(client.TLSConfig.RootCert) {
			dsParams = append(dsParams, "sslrootcert="+client.TLSConfig.RootCert)
		}
	}

	dsName := strings.Join(dsParams, " ")
	db, err := sql.Open("postgres", dsName)
	if err != nil {
		return err
	}
	client.db = db
	return nil
}

// DB returns a connected database instance.
func (client *PqClient) DB() *sql.DB {
	return client.db
}

// Close closes opens a database specified by the internal configuration.
func (client *PqClient) Close() error {
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
func (client *PqClient) Ping() error {
	return client.db.Ping()
}

// Query executes a query that returns rows.
func (client *PqClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if client.db == nil {
		err := client.Open()
		if err != nil {
			return nil, err
		}
	}
	return client.db.Query(query, args...)
}

// CreateDatabase creates a specified database.
func (client *PqClient) CreateDatabase(name string) error {
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
func (client *PqClient) DropDatabase(name string) error {
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
func (client *PqClient) Use(name string) error {
	client.SetDatabase(name)
	return nil
}
