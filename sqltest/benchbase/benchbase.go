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

package benchbase

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cybergarage/go-logger/log"
)

const (
	benchbaseRoot         = "BENCHBASE_ROOT"
	benchbaseConfigEnv    = "BENCHBASE_CONFIG"
	benchbaseBenchEnv     = "BENCHBASE_BENCH"
	DatabaseName          = "benchbase"
	benchbaseDefaultBench = "tpcc"
)

func SetUpQueries() []string {
	return []string{
		"CREATE DATABASE IF NOT EXISTS benchbase;",
	}
}

func RunWorkload(t *testing.T, defaultBench string) error {
	t.Helper()
	outputBenchbaseParams := func(t *testing.T, benchbaseEnvs []string, benchbaseParams []string) {
		t.Helper()
		for n, benchbaseEnv := range benchbaseEnvs {
			t.Logf("%s = %s", benchbaseEnv, benchbaseParams[n])
		}
	}

	benchbaseEnvs := []string{
		benchbaseRoot,
		benchbaseConfigEnv,
		benchbaseBenchEnv,
	}

	benchbaseParams := []string{
		"",
		"config.xml",
		defaultBench,
	}

	for n, benchbaseEnv := range benchbaseEnvs {
		if v, ok := os.LookupEnv(benchbaseEnv); ok {
			benchbaseParams[n] = v
		}
		if len(benchbaseParams[n]) == 0 {
			outputBenchbaseParams(t, benchbaseEnvs, benchbaseParams)
			t.Skipf("%s is not specified", benchbaseEnv)
			return nil
		}
	}

	outputBenchbaseParams(t, benchbaseEnvs, benchbaseParams)

	benchbasePath := benchbaseParams[0]
	benchbaseCmd := filepath.Join(benchbasePath, "bin/benchbase")
	_, err := os.Stat(benchbaseCmd)
	if err != nil {
		t.Skip(err)
		return err
	}

	bench := benchbaseParams[2]
	configFile := benchbaseParams[1]
	configPath := filepath.Join(benchbasePath, "config", configFile)

	benchbaseArgs := []string{
		benchbaseCmd,
		"",
		"--bench",
		bench,
		"--config",
		configPath,
	}

	benchbaseWorkloadCmds := []string{
		"load",
		"execute",
	}

	for _, benchbaseWorkloadCmd := range benchbaseWorkloadCmds {
		t.Run(benchbaseWorkloadCmd, func(t *testing.T) {
			benchbaseArgs[1] = benchbaseWorkloadCmd
			cmdStr := strings.Join(benchbaseArgs, " ")
			log.Debugf("%v", cmdStr)
			t.Logf("%v", cmdStr)
			out, err := exec.Command(benchbaseCmd, benchbaseArgs[1:]...).CombinedOutput()
			if err != nil {
				t.Error(err)
				return
			}
			outStr := string(out)
			if strings.Contains(outStr, "FAILED") || strings.Contains(outStr, "ERROR") {
				t.Errorf("%s", outStr)
				return
			}
			t.Logf("%s", outStr)
		})
	}

	return nil
}
