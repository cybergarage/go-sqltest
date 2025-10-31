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
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

const (
	benchbaseRoot        = "BENCHBASE_ROOT"
	benchbaseConfigEnv   = "BENCHBASE_CONFIG"
	benchbaseBenchEnv    = "BENCHBASE_BENCH"
	DatabaseName         = "benchbase"
	defaultBenchbaseRoot = "./"
	defaultBench         = "tpcc"
	defaultConfig        = "tpcc_config.xml"
	benchbaseJarFile     = "benchbase.jar"
)

// IsInstalled checks if BenchBase is properly installed and accessible.
func IsInstalled() bool {
	root := os.Getenv(benchbaseRoot)
	if root == "" {
		root = defaultBenchbaseRoot
	}
	jarPath := filepath.Join(root, benchbaseJarFile)
	cmd := exec.Command("java", "-jar", jarPath, "-h")
	err := cmd.Run()
	return err == nil
}

// RunWorkload runs a BenchBase benchmark using a single java -jar invocation:
// Environment variables:
//
//	BENCHBASE_ROOT  : Root directory where benchbase jar & config/ reside
//	BENCHBASE_CONFIG: Relative path under BENCHBASE_ROOT to config XML (defaults to tpcc_config.xml)
//	BENCHBASE_BENCH : Benchmark name (defaults to value passed as tpcc)
func RunWorkload(t *testing.T, benches ...string) error {
	t.Helper()

	// Gather parameters from environment or defaults.
	root := os.Getenv(benchbaseRoot)
	if root == "" {
		root = defaultBenchbaseRoot
	}

	if len(benches) == 0 {
		bench := os.Getenv(benchbaseBenchEnv)
		if bench == "" {
			bench = defaultBench
		}
		benches = []string{bench}
	}

	configRel := os.Getenv(benchbaseConfigEnv)
	if configRel == "" {
		configRel = defaultConfig
	}

	jarPath := filepath.Join(root, benchbaseJarFile)
	configPath := filepath.Join(root, configRel)

	for _, bench := range benches {
		args := []string{
			"-jar", jarPath,
			"-b", bench,
			"-c", configPath,
			"--create=true",
			"--load=true",
			"--execute=true",
		}

		t.Logf("benchbase root    : %s", root)
		t.Logf("benchbase jar     : %s", jarPath)
		t.Logf("benchbase bench   : %s", bench)
		t.Logf("benchbase config  : %s", configPath)
		t.Logf("benchbase command : java %v", args)

		start := time.Now()
		cmd := exec.Command("java", args...)
		out, err := cmd.CombinedOutput()
		dur := time.Since(start)

		if err != nil {
			err := errors.New("benchbase execution failed: " + err.Error() + "\n" + string(out))
<<<<<<< HEAD
			t.Logf("error: \n%s", err.Error())
=======
			t.Logf("error: \n%s", string(out))
>>>>>>> ad8660e (test: Update benchbase.RunWorkload())
			t.Skip(err)
			return err
		}

		t.Logf("benchbase duration: %s", dur)
		t.Logf("benchbase output:\n%s", string(out))
	}
	return nil
}
