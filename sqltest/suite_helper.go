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
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/cybergarage/go-sqltest/sqltest/test"
)

const TestDBNamePrefix = "sqltest"

// RunEmbedSuites runs the embedded test suites.
func RunEmbedSuites(t *testing.T, client Client, testNames ...string) error {
	t.Helper()

	cs, err := NeweEmbedSuite(test.EmbedTests)
	if err != nil {
		t.Error(err)
		return err
	}

	for _, test := range cs.ScenarioTests() {
		if 0 < len(testNames) {
			found := false
			for _, testName := range testNames {
				if strings.EqualFold(test.Name(), testName) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		var err error
		t.Run(test.Name(), func(t *testing.T) {
			testDBName := fmt.Sprintf("%s%d", TestDBNamePrefix, time.Now().UnixNano())

			client.SetDatabase(testDBName)

			err = client.Open()
			if err != nil {
				t.Error(err)
				return
			}

			defer func() {
				err := client.Close()
				if err != nil {
					t.Error(err)
				}
			}()

			err = client.CreateDatabase(testDBName)
			if err != nil {
				t.Error(err)
				return
			}

			defer func() {
				err := client.DropDatabase(testDBName)
				if err != nil {
					t.Error(err)
				}
			}()

			test.SetClient(client)

			err = test.Run()
			if err != nil {
				t.Errorf("%s : %s", test.Name(), err.Error())
			}
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func RunLocalSuite(t *testing.T) {
	t.Helper()

	cs, err := NewSuiteWithDirectory(SuiteDefaultTestDirectory)
	if err != nil {
		t.Error(err)
		return
	}

	for _, test := range cs.tests {
		t.Run(test.Name(), func(t *testing.T) {
			testDBName := fmt.Sprintf("%s%d", TestDBNamePrefix, time.Now().UnixNano())

			client := NewMySQLClient()
			client.SetDatabase(testDBName)

			err = client.Open()
			if err != nil {
				t.Error(err)
				return
			}

			defer func() {
				err := client.Close()
				if err != nil {
					t.Error(err)
				}
			}()

			err = client.CreateDatabase(testDBName)
			if err != nil {
				t.Error(err)
			}

			defer func() {
				err := client.DropDatabase(testDBName)
				if err != nil {
					t.Error(err)
				}
			}()
			test.SetClient(client)
			err := test.Run()
			if err != nil {
				t.Errorf("%s : %s", test.Name(), err.Error())
			}
		})
	}
}
