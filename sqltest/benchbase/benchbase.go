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
	benchbaseRoot         = "BENCHBASE_ROOT"
	benchbaseConfigEnv    = "BENCHBASE_CONFIG"
	benchbaseBenchEnv     = "BENCHBASE_BENCH"
	DatabaseName          = "benchbase"
	benchbaseDefaultBench = "tpcc"
)

// Installed checks if BenchBase is properly installed and accessible.
func Installed() bool {
	root := os.Getenv(benchbaseRoot)
	if root == "" {
		return false
	}

	// Check for jar file
	jarPath := filepath.Join(root, "benchbase.jar")
	if _, err := os.Stat(jarPath); err != nil {
		matches, globErr := filepath.Glob(filepath.Join(root, "benchbase-*.jar"))
		if globErr != nil || len(matches) == 0 {
			return false
		}
		jarPath = matches[0]
	}

	// Test if java and benchbase jar are accessible
	cmd := exec.Command("java", "-jar", jarPath, "-h")
	err := cmd.Run()
	return err == nil
}

// RunWorkload runs a BenchBase benchmark using a single java -jar invocation:
//
//	java -jar benchbase.jar -b <bench> -c <config> --create=true --load=true --execute=true
//
// Environment variables:
//
//	BENCHBASE_ROOT  : Root directory where benchbase jar & config/ reside
//	BENCHBASE_CONFIG: Relative path under BENCHBASE_ROOT to config XML (defaults to config/postgres/sample_tpcc_config.xml)
//	BENCHBASE_BENCH : Benchmark name (defaults to value passed as defaultBench)
func RunWorkload(t *testing.T, defaultBench string) error {
	t.Helper()

	// Gather parameters from environment or defaults.
	root := os.Getenv(benchbaseRoot)
	if root == "" {
		return errors.New("BENCHBASE_ROOT is not specified")
	}

	bench := os.Getenv(benchbaseBenchEnv)
	if bench == "" {
		bench = defaultBench
	}

	configRel := os.Getenv(benchbaseConfigEnv)
	if configRel == "" {
		// Provide more realistic default path aligning with example command.
		configRel = "config/postgres/sample_tpcc_config.xml"
	}

	// Resolve jar: try explicit benchbase.jar, then glob benchbase-*.jar.
	jarPath := filepath.Join(root, "benchbase.jar")
	if _, err := os.Stat(jarPath); err != nil {
		matches, globErr := filepath.Glob(filepath.Join(root, "benchbase-*.jar"))
		if globErr != nil {
			return globErr
		}
		if len(matches) == 0 {
			return errors.New("benchbase jar not found (benchbase.jar or benchbase-*.jar)")
		}
		jarPath = matches[0]
	}

	configPath := filepath.Join(root, configRel)
	if _, err := os.Stat(configPath); err != nil {
		return err
	}

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
		// Return full output for debugging.
		return errors.New("benchbase execution failed: " + err.Error() + "\n" + string(out))
	}

	t.Logf("benchbase duration: %s", dur)
	t.Logf("benchbase output:\n%s", string(out))

	return nil
}
