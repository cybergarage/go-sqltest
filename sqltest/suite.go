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
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/cybergarage/go-sqltest/sqltest/test"
	"github.com/cybergarage/go-sqltest/sqltest/util"
)

const (
	TestDBNamePrefix          = "sqltest"
	SuiteDefaultTestDirectory = "./test"
)

// SuiteOption is a function type used to configure a Suite.
type SuiteOption func(*Suite) error

// SuiteErrorHandler is a function type used to handle errors in a Suite.
type SuiteErrorHandler func(*Suite, *ScenarioRunner, error)

// WithSuiteClient returns a SuiteOption that sets a client for testing.
func WithSuiteClient(client Client) SuiteOption {
	return func(suite *Suite) error {
		suite.SetClient(client)
		return nil
	}
}

// WithSuiteDirectories returns a SuiteOption that loads scenario tests from the specified directory.
func WithSuiteDirectories(dirs ...string) SuiteOption {
	return func(suite *Suite) error {
		return suite.LoadDirectorySenarios(dirs...)
	}
}

// WithSuiteEmbeds returns a SuiteOption that loads scenario tests from the specified bytes.
func WithSuiteEmbeds(tests ...map[string][]byte) SuiteOption {
	return func(suite *Suite) error {
		if len(tests) == 0 {
			tests = []map[string][]byte{test.EmbedTests}
		}
		return suite.LoadEmbedSenarios(tests...)
	}
}

// WithSuiteRegexes returns a SuiteOption that extracts scenario tests with the specified regexes.
func WithSuiteRegexes(regexes ...string) SuiteOption {
	return func(suite *Suite) error {
		tests, err := suite.ExtractScenarioMatchingTests(regexes...)
		if err != nil {
			return err
		}
		suite.tests = tests
		return nil
	}
}

// WithSuiteErrorHandler returns a SuiteOption that sets a handler for errors in a Suite.
func WithSuiteErrorHandler(handler SuiteErrorHandler) SuiteOption {
	return func(suite *Suite) error {
		suite.errorHandler = handler
		return nil
	}
}

// WithSuiteStepHandler returns a SuiteOption that sets a handler for errors in a Suite.
func WithSuiteStepHandler(handler ScenarioStepHandler) SuiteOption {
	return func(suite *Suite) error {
		for _, test := range suite.tests {
			test.SetStepHandler(handler)
		}
		return nil
	}
}

// Suite represents a scenario test suite.
type Suite struct {
	tests        []*ScenarioRunner
	client       Client
	errorHandler SuiteErrorHandler
}

// NewSuite returns a scenario test suite instance.
func NewSuite() *Suite {
	suite := &Suite{
		tests:        make([]*ScenarioRunner, 0),
		client:       nil,
		errorHandler: nil,
	}
	return suite
}

// NewSuite returns a scenario test suite instance with the specified options.
func NewSuiteWith(opts ...SuiteOption) (*Suite, error) {
	suite := NewSuite()
	for _, opt := range opts {
		if err := opt(suite); err != nil {
			return nil, err
		}
	}
	return suite, nil
}

// NewSuiteWithDirectory returns a scenario test suite instance which loads under the specified directory.
func NewSuiteWithDirectory(dirs ...string) (*Suite, error) {
	suite := NewSuite()
	return suite, suite.LoadDirectorySenarios(dirs...)
}

// NeweEmbedSuite returns a scenario test suite instance which loads under the specified directory.
func NeweEmbedSuite(tests ...map[string][]byte) (*Suite, error) {
	suite := NewSuite()
	if len(tests) == 0 {
		tests = []map[string][]byte{test.EmbedTests}
	}
	return suite, suite.LoadEmbedSenarios(tests...)
}

// SetClient sets a client for testing.
func (suite *Suite) SetClient(c Client) {
	suite.client = c
}

// SetErrorHandler sets a handler for errors in a Suite.
func (suite *Suite) SetErrorHandler(handler SuiteErrorHandler) {
	suite.errorHandler = handler
}

// LoadDirectorySenarios loads scenario tests from the specified directories.
func (suite *Suite) LoadDirectorySenarios(dirs ...string) error {
	for _, dir := range dirs {
		re := regexp.MustCompile(".*\\." + ScenarioFileExt)
		findPath := util.NewFileWithPath(dir)
		files, err := findPath.ListFilesWithRegexp(re)
		if err != nil {
			return err
		}

		for _, file := range files {
			s, err := NewScenarioRunnerWithFile(file.Path)
			if err != nil {
				return err
			}
			suite.tests = append(suite.tests, s)
		}
	}
	return nil
}

// LoadEmbedSenarios loads scenario tests from the specified bytes.
func (suite *Suite) LoadEmbedSenarios(testMaps ...map[string][]byte) error {
	for _, testMap := range testMaps {
		for name, b := range testMap {
			s, err := NewScenarioRunnerWithBytes(name, b)
			if err != nil {
				return err
			}
			suite.tests = append(suite.tests, s)
		}
	}
	return nil
}

// ScenarioTests returns all loaded scenario tests.
func (suite *Suite) ScenarioTests() []*ScenarioRunner {
	return suite.tests
}

// ExtractScenarioTests returns scenario tests with the specified names.
func (suite *Suite) ExtractScenarioTests(regexpNames ...string) ([]*ScenarioRunner, error) {
	tests := make([]*ScenarioRunner, 0)
	for _, name := range regexpNames {
		nameRegexp := regexp.MustCompile(name)
		isFound := false
		for _, test := range suite.tests {
			if strings.EqualFold(test.Name(), name) || nameRegexp.MatchString(test.Name()) {
				tests = append(tests, test)
				isFound = true
			}
		}
		if !isFound {
			return nil, fmt.Errorf("%s is not found", name)
		}
	}
	return tests, nil
}

func (suite *Suite) ExtractScenarioMatchingTests(regexes ...string) ([]*ScenarioRunner, error) {
	tests := make([]*ScenarioRunner, 0)
	for _, test := range suite.tests {
		for _, regex := range regexes {
			re, err := regexp.Compile(regex)
			if err != nil {
				return tests, err
			}
			if re.MatchString(test.Name()) {
				tests = append(tests, test)
				continue
			}
		}
	}
	return tests, nil
}

// Run runs all loaded scenario tests. The method stops the testing when a scenario test is aborted, and the following tests are not run.
func (suite *Suite) Run() error {
	for _, test := range suite.tests {
		test.SetClient(suite.client)
		err := test.Run()
		if err != nil {
			return fmt.Errorf("%s : %w", test.Name(), err)
		}
	}
	return nil
}

// Test runs all loaded scenario tests with regular expressions for scenarios. The method stops the testing when a scenario test is aborted, and the following tests are not run.
func (suite *Suite) Test(t *testing.T) error {
	var err error
	t.Run(TestRunDescription, func(t *testing.T) {
		for _, test := range suite.tests {
			t.Run(test.Name(), func(t *testing.T) {
				err = suite.TestScenario(t, test)
			})
		}
	})
	return err
}

// RunScenarioTest runs the specified test.
func (suite *Suite) TestScenario(t *testing.T, test *ScenarioRunner) error {
	t.Helper()

	var err error
	client := suite.client

	testDBName := fmt.Sprintf("%s%d", TestDBNamePrefix, time.Now().UnixNano())

	client.SetDatabase(testDBName)

	err = client.Open()
	if err != nil {
		t.Error(err)
		return err
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
		return err
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
		if suite.errorHandler != nil {
			suite.errorHandler(suite, test, err)
		}
	}

	return err
}
