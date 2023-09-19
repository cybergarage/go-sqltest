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
	"context"
	"database/sql"
	"fmt"
	"net"
	"strconv"

	"github.com/jackc/pgx/v5"
)

// PgxClient represents a client for PostgreSQL server.
type PgxClient struct {
	*Config
	conn *pgx.Conn
}

// NewPgxClient returns a Pgx client instance.
func NewPgxClient() Client {
	client := &PgxClient{
		Config: NewDefaultConfig(),
		conn:   nil,
	}
	client.SetPort(defaultPostgresPort)
	return client
}

// Open opens a database specified by the client.
func (client *PgxClient) Open() error { // nolint: nosprintfhostport
	url := fmt.Sprintf("postgres://%s:%s@%s/%s",
		"user",
		"password",
		net.JoinHostPort(client.Host, strconv.Itoa(client.Port)),
		client.Database)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return err
	}
	client.conn = conn
	return nil
}

// Close closes the opened database.
func (client *PgxClient) Close() error {
	if client.conn == nil {
		return nil
	}
	err := client.conn.Close(context.Background())
	if err != nil {
		return err
	}
	client.conn = nil
	return nil
}

// Conn returns a connected database connection.
func (client *PgxClient) Conn() *pgx.Conn {
	return client.conn
}

// Query executes a query that returns rows.
func (client *PgxClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

// CreateDatabase creates a specified database.
func (client *PgxClient) CreateDatabase(name string) error {
	return nil
}

// DropDatabase dtops a specified database.
func (client *PgxClient) DropDatabase(name string) error {
	return nil
}

// Use sets a target database.
func (client *PgxClient) Use(name string) error {
	return nil
}