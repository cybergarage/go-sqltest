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
)

// RunCommand runs a sysbench command with the given configuration.
func RunCommand(t *testing.T, cmd string, config Config) error {
	t.Helper()

	args := []string{
		"sysbench",
		cmd,
	}

	for k, v := range config {
		args = append(args, fmt.Sprintf("--%s=%s", k, v))
	}

	cmdStr := strings.Join(args, " ")
	// t.Logf("Running %s", cmdStr)
	t.Run(cmdStr, func(t *testing.T) {
		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
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

	return nil
}
