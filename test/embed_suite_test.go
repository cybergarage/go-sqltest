// Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
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
	"testing"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-sqltest/sqltest"
)

func TestEmbedSuitesForDebug(t *testing.T) {
	log.EnableStdoutDebug(true)

	// server := server.NewServer()
	// err := server.Start()
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// defer server.Stop()

	client := sqltest.NewPostgresClient()
	client.SetPreparedStatementEnabled(false)
	client.SetUser("postgres")
	client.SetPassword("passwd")

	testRegexes := []string{
		// "SmplTxnText",
		// "SmplCrudInt",
		// "SmplCrud.*",
	}

	opts := []sqltest.SuiteOption{
		sqltest.WithSuiteClient(client),
		sqltest.WithSuiteRegexes(testRegexes...),
		sqltest.WithSuiteQueryDialect(sqltest.QueryDialectPostgreSQL),
	}

	if err := sqltest.RunEmbedSuitesWith(t, opts...); err != nil {
		t.Error(err)
	}
}
