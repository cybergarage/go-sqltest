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

// IsInstalled checks if sysbench is installed on the system.
func IsInstalled() bool {
	_, err := exec.LookPath(program)
	return err == nil
}

// GenerateTempDBName returns a temporary database name.
func GenerateTempDBName() string {
	return fmt.Sprintf("%s%d", dbNamePrefix, time.Now().UnixNano())
}

// User returns the default user for sysbench.
func User() string {
	return "sbuser"
}

// Password returns the default password for sysbench.
func Password() string {
	return "sbpassword"
}

// RunCommand runs a sysbench command with the given configuration.
func RunCommand(t *testing.T, cmd string, config *Config) error {
	t.Helper()

	path, err := exec.LookPath(program)
	if err != nil {
		t.Skipf("%s is not installed: %v", program, err)
		return nil
	}
	t.Logf("%s is installed at %s", program, path)

	toCommandParam := func(k, v string) string {
		return fmt.Sprintf("--%s=%s", k, v)
	}

	toCommandLineArgs := func(config *Config) []string {
		args := []string{}
		for _, key := range config.Keys() {
			val, ok := config.Value(key)
			if !ok {
				continue
			}
			args = append(args, toCommandParam(key, val))
		}
		return args
	}

	toCommandLine := func(prgram string, args []string) string {
		cmdLine := []string{prgram}
		cmdLine = append(cmdLine, args...)
		return strings.Join(cmdLine, " ")
	}

	t.Run(fmt.Sprintf("%s (%s)", program, cmd), func(t *testing.T) {
		subCmds := []string{
			"prepare",
			"run",
			"cleanup",
		}
		for _, subCmd := range subCmds {
			t.Run(subCmd, func(t *testing.T) {
				args := toCommandLineArgs(config)
				args = append(args, cmd)
				args = append(args, subCmd)
				t.Log(toCommandLine(program, args))
				out, err := exec.Command(program, args...).CombinedOutput()
				outStr := string(out)
				if err != nil {
					if config.SkipOnError() {
						t.Skip(toCommandLine(program, args))
						t.Skip(err)
						t.Skip(outStr)
					} else {
						t.Error(toCommandLine(program, args))
						t.Error(err)
						t.Error(outStr)
					}
					return
				}
				if strings.Contains(outStr, "FAILED") {
					if config.SkipOnError() {
						t.Skip(toCommandLine(program, args))
						t.Skip(err)
					} else {
						t.Error(toCommandLine(program, args))
						t.Error(outStr)
					}
					return
				}
				t.Log(outStr)
				if t.Skipped() {
					t.Logf("Test skipped: %s", subCmd)
				}
				if t.Failed() {
					t.Logf("Test failed: %s", subCmd)
				}
			})
		}
	})

	return nil
}
