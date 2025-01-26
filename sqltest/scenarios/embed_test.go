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

package test

import "testing"

func TestEmbedScenarios(t *testing.T) {
	if len(EmbedScenarios) == 0 {
		t.Errorf("EmbedTests: %v", len(EmbedScenarios))
		return
	}
	for name, bytes := range EmbedScenarios {
		if len(bytes) == 0 {
			t.Errorf("%s bytes: %d", name, len(bytes))
		}
	}
}
