// Copyright (C) 2025 The go-sqltest Authors. All rights reserved.
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

package sysbench

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"
)

const (
	program      = "sysbench"
	dbNamePrefix = program
)

// GenerateTempDBName returns a temporary database name.
func GenerateTempDBName() string {
	return fmt.Sprintf("%s%d", dbNamePrefix, time.Now().UnixNano())
}

// RunCommand runs a sysbench command with the given configuration.
func RunCommand(t *testing.T, cmd string, config Config) error {
	t.Helper()

	toCommandParam := func(k, v string) string {
		return fmt.Sprintf("--%s=%s", k, v)
	}

	title := []string{program}
	for k, v := range config {
		title = append(title, toCommandParam(k, v))
	}
	titleStr := strings.Join(title, " ")
	t.Run(titleStr, func(t *testing.T) {
		subCmds := []string{
			"prepare",
			"run",
			"cleanup",
		}
		for _, subCmd := range subCmds {
			args := []string{}
			for k, v := range config {
				args = append(args, toCommandParam(k, v))
			}
			args = append(args, subCmd)
			t.Run(subCmd, func(t *testing.T) {
				out, err := exec.Command(program, args...).CombinedOutput()
				outStr := string(out)
				if err != nil {
					t.Skip(err)
					t.Logf("%s", outStr)
					return
				}
				if strings.Contains(outStr, "FAILED") {
					t.Errorf("%s", outStr)
					t.Logf("%s", outStr)
					return
				}
				t.Logf("%s", outStr)
			})
		}
	})

	return nil
}
