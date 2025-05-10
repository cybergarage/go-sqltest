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
)

// RunEmbedSuites runs the embedded test suites with the specified regular expressions.
func RunEmbedSuites(t *testing.T, client Client, regexes ...string) error {
	t.Helper()

	opts := []SuiteOption{
		WithSuiteClient(client),
		WithSuiteRegexes(regexes...),
	}

	return RunEmbedSuitesWith(t, opts...)
}

// RunEmbedSuitesWith runs the embedded test suites with the specified options.
func RunEmbedSuitesWith(t *testing.T, opts ...SuiteOption) error {
	t.Helper()

	embedOpts := []SuiteOption{
		WithSuiteEmbeds(),
	}

	embedOpts = append(embedOpts, opts...)

	suite, err := NewSuiteWith(embedOpts...)
	if err != nil {
		return err
	}

	return suite.Test(t)
}
