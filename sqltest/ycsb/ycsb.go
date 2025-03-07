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

package ycsb

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cybergarage/go-logger/log"
)

const (
	ycsbRoot            = "YCSB_ROOT"
	ycsbDBPropertiesEnv = "YCSB_DB_PROPERTIES"
	ycsbWorkloadEnv     = "YCSB_WORKLOAD"
	DatabaseName        = "ycsb"
	ycsbDefaultWorkload = "workloada"
)

func SetUpQueries() []string {
	return []string{
		"CREATE TABLE IF NOT EXISTS usertable (YCSB_KEY VARCHAR(255) PRIMARY KEY, FIELD0 TEXT, FIELD1 TEXT, FIELD2 TEXT, FIELD3 TEXT, FIELD4 TEXT, FIELD5 TEXT, FIELD6 TEXT, FIELD7 TEXT, FIELD8 TEXT, FIELD9 TEXT);",
	}
}

func RunWorkload(t *testing.T, defaultWorkload string) error {
	t.Helper()
	outputYcsbParams := func(t *testing.T, ycsbEnvs []string, ycsbParams []string) {
		t.Helper()
		for n, ycsbEnv := range ycsbEnvs {
			t.Logf("%s = %s", ycsbEnv, ycsbParams[n])
		}
	}

	ycsbEnvs := []string{
		ycsbRoot,
		ycsbDBPropertiesEnv,
		ycsbWorkloadEnv,
	}

	ycsbParams := []string{
		"",
		"db.properties",
		defaultWorkload,
	}

	for n, ycsbEnv := range ycsbEnvs {
		if v, ok := os.LookupEnv(ycsbEnv); ok {
			ycsbParams[n] = v
		}
		if len(ycsbParams[n]) == 0 {
			outputYcsbParams(t, ycsbEnvs, ycsbParams)
			t.Skipf("%s is not specified", ycsbEnv)
			return nil
		}
	}

	outputYcsbParams(t, ycsbEnvs, ycsbParams)

	ycsbPath := ycsbParams[0]
	ycsbCmd := filepath.Join(ycsbPath, "bin/ycsb.sh")
	_, err := os.Stat(ycsbCmd)
	if err != nil {
		t.Skip(err)
		return err
	}

	workload := ycsbParams[2]
	workloadDir := filepath.Join(ycsbPath, "workloads")
	workloadFile := filepath.Join(workloadDir, workload)

	ycsbArgs := []string{
		ycsbCmd,
		"",
		"jdbc",
		"-P",
		workloadFile,
		"-P",
		ycsbParams[1],
	}

	ycsbWorkloadCmds := []string{
		"load",
		"run",
	}

	for _, ycsbWorkloadCmd := range ycsbWorkloadCmds {
		t.Run(ycsbWorkloadCmd, func(t *testing.T) {
			ycsbArgs[1] = ycsbWorkloadCmd
			cmdStr := strings.Join(ycsbArgs, " ")
			log.Debugf("%v", cmdStr)
			t.Logf("%v", cmdStr)
			out, err := exec.Command(ycsbCmd, ycsbArgs[1:]...).CombinedOutput()
			if err != nil {
				t.Error(err)
				return
			}
			outStr := string(out)
			if strings.Contains(outStr, "FAILED") {
				t.Errorf("%s", outStr)
				return
			}
			t.Logf("%s", outStr)
		})
	}

	return nil
}
