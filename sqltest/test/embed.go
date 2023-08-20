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

import (
	_ "embed"
)

// EmbedTests is a map of test names and test queries.
var EmbedTests = map[string][]byte{
	"SimpAlterAdd":     simpAlterAdd,
	"YcsbWorkload":     ycsbWorkload,
	"SimpInsertSelect": simpInsertSelect,
	"Pgbench":          pgbench,
}

//go:embed simp_alter_add.qst
var simpAlterAdd []byte

//go:embed ycsb_workload.qst
var ycsbWorkload []byte

//go:embed simp_insert_select.qst
var simpInsertSelect []byte

//go:embed pgbench.qst
var pgbench []byte
