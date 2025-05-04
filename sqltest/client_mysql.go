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
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	defaultMysqlPort = 3306
)

// MySQLClient represents a client for MySQL server.
type MySQLClient struct {
	*sqlClient
}

// NewMySQLClient returns a client instance.
func NewMySQLClient() Client {
	client := &MySQLClient{
		sqlClient: newSQLClient(),
	}
	client.SetPort(defaultMysqlPort)
	return client
}

// Open opens a database specified by the internal configuration.
// nolint: gosec, exhaustruct, staticcheck
func (client *MySQLClient) Open() error {
	if client.TLSEnabled() {
		rootCertPool := x509.NewCertPool()
		pem, err := os.ReadFile(client.RootCertFile)
		if err != nil {
			return err
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			return err
		}
		clientCert := make([]tls.Certificate, 0, 1)
		certs, err := tls.LoadX509KeyPair(client.ClientCertFile, client.ClientKeyFile)
		if err != nil {
			return err
		}
		clientCert = append(clientCert, certs)
		mysql.RegisterTLSConfig("custom", &tls.Config{
			RootCAs:            rootCertPool,
			Certificates:       clientCert,
			InsecureSkipVerify: true,
		})
	}

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		client.User,
		client.Password,
		client.Host,
		client.Port,
		client.Database)

	dbURLParams := []string{
		"parseTime=true",
	}
	if client.TLSEnabled() {
		dbURLParams = append(dbURLParams, "tls=custom")
	}

	if dbURLParamCnt := len(dbURLParams); 0 < dbURLParamCnt {
		dbURL += "?" + dbURLParams[0]
		for n := 1; n < dbURLParamCnt; n++ {
			dbURL += "&" + dbURLParams[n]
		}
	}

	db, err := sql.Open("mysql", dbURL)
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

// Query executes a query that returns rows.
func (client *MySQLClient) Query(query string, args ...any) (*sql.Rows, error) {
	if client.db == nil {
		err := client.Open()
		if err != nil {
			return nil, err
		}
	}
	return client.db.Query(query, args...)
}
