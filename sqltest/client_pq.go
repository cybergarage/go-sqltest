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
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// PqClient represents a client for PostgreSQL server.
type PqClient struct {
	*sqlClient
}

// NewPqClient returns a new lib/pq client.
func NewPqClient() *PqClient {
	client := &PqClient{
		sqlClient: newSQLClient(),
	}
	client.SetPort(defaultPostgresPort)
	return client
}

// Open opens a database specified by the internal configuration.
func (client *PqClient) Open() error {
	dsParams := []string{
		"host=" + client.Host,
		"port=" + strconv.Itoa(client.Port),
		"connect_timeout=" + strconv.Itoa(0),
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
	switch client.Auth {
	case AuthPlain:
		dsParams = append(dsParams, "require_auth=password")
	case AuthMD5:
		dsParams = append(dsParams, "require_auth=md5")
	case AuthSCRAMSHA256:
		dsParams = append(dsParams, "require_auth=scram-sha-256")
	}
	if client.TLSEnabled() {
		dsParams = append(dsParams, "sslmode=require")
		if 0 < len(client.ClientCertFile) {
			dsParams = append(dsParams, "sslcert="+client.ClientCertFile)
		}
		if 0 < len(client.ClientKeyFile) {
			dsParams = append(dsParams, "sslkey="+client.ClientKeyFile)
		}
		if 0 < len(client.RootCertFile) {
			dsParams = append(dsParams, "sslrootcert="+client.RootCertFile)
		}
	} else {
		dsParams = append(dsParams, "sslmode=disable")
	}

	dsName := strings.Join(dsParams, " ")
	db, err := sql.Open("postgres", dsName)
	if err != nil {
		return err
	}
	client.db = db
	return nil
}
