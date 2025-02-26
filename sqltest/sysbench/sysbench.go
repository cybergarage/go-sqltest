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
	"os/exec"
	"strings"
	"testing"

	"github.com/cybergarage/go-logger/log"
)

func Run(t *testing.T) error {
	t.Helper()

	args := []string{
		"sysbench",
	}

	cmd := strings.Join(args, " ")
	log.Debugf("%v", cmd)
	t.Logf("%v", cmd)

	t.Run(cmd, func(t *testing.T) {
		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
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

	return nil
}
