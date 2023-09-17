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
	"testing"
	"time"

	"github.com/cybergarage/go-sqltest/sqltest/test"
)

const TestDBNamePrefix = "sqltest"

// RunScenarioTest runs the specified test.
func RunScenarioTest(t *testing.T, client Client, test *ScenarioTest) {
	t.Helper()

	var err error
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
}

// RunEmbedSuites runs the embedded test suites.
func RunEmbedSuites(t *testing.T, client Client, testNames ...string) error {
	t.Helper()

	es, err := NeweEmbedSuite(test.EmbedTests)
	if err != nil {
		t.Error(err)
		return err
	}

	tests := es.ScenarioTests()
	if 0 < len(testNames) {
		tests, err = es.ExtractScenarioTests(testNames...)
		if err != nil {
			t.Error(err)
			return err
		}
	}

	for _, test := range tests {
		t.Run(test.Name(), func(t *testing.T) {
			RunScenarioTest(t, client, test)
		})
	}

	return nil
}

// RunEmbedSuitesWithRegex runs the embedded test suites with the specified regular expressions.
func RunEmbedSuitesWithRegex(t *testing.T, client Client, regexes ...string) error {
	t.Helper()

	es, err := NeweEmbedSuite(test.EmbedTests)
	if err != nil {
		t.Error(err)
		return err
	}

	tests, err := es.ExtractScenarioMatchingTests(regexes...)
	if err != nil {
		t.Error(err)
		return err
	}

	for _, test := range tests {
		t.Run(test.Name(), func(t *testing.T) {
			RunScenarioTest(t, client, test)
		})
	}

	return nil
}

// RunLocalSuite runs the local test suite.
func RunLocalSuite(t *testing.T, client Client) error {
	t.Helper()

	cs, err := NewSuiteWithDirectory(SuiteDefaultTestDirectory)
	if err != nil {
		t.Error(err)
		return err
	}

	for _, test := range cs.tests {
		t.Run(test.Name(), func(t *testing.T) {
			RunScenarioTest(t, client, test)
		})
	}
	return nil
}
