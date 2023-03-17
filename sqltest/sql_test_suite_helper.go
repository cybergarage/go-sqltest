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
	"testing"

	"github.com/cybergarage/go-sqltest/sqltest/test"
)

const sqlTestDatabase = "tst"

func RunSQLTestSuite(t *testing.T, client Client) {
	t.Helper()

	cs, err := NeweEmbedSQLTestSuite(test.EmbedTests)
	if err != nil {
		t.Error(err)
		return
	}

	client.SetDatabase(sqlTestDatabase)

	err = client.Open()
	if err != nil {
		t.Error(err)
		return
	}

	err = client.CreateDatabase(sqlTestDatabase)
	if err != nil {
		t.Error(err)
		return
	}

	cs.SetClient(client)

	for _, test := range cs.Tests {
		t.Run(test.Name(), func(t *testing.T) {
			test.SetClient(cs.client)
			err := test.Run()
			if err != nil {
				t.Errorf("%s : %s", test.Name(), err.Error())
			}
		})
	}

	err = client.DropDatabase(sqlTestDatabase)
	if err != nil {
		t.Error(err)
	}

	err = client.Close()
	if err != nil {
		t.Error(err)
	}
}

func RunLocalSQLTestSuite(t *testing.T) {
	t.Helper()

	cs, err := NewSQLTestSuiteWithDirectory(SQLTestSuiteDefaultTestDirectory)
	if err != nil {
		t.Error(err)
		return
	}

	client := NewMySQLClient()
	client.SetDatabase(sqlTestDatabase)

	err = client.Open()
	if err != nil {
		t.Error(err)
		return
	}

	err = client.CreateDatabase(sqlTestDatabase)
	if err != nil {
		t.Error(err)
	}

	cs.SetClient(client)

	for _, test := range cs.Tests {
		t.Run(test.Name(), func(t *testing.T) {
			test.SetClient(cs.client)
			err := test.Run()
			if err != nil {
				t.Errorf("%s : %s", test.Name(), err.Error())
			}
		})
	}

	err = client.DropDatabase(sqlTestDatabase)
	if err != nil {
		t.Error(err)
	}
}
