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

var QueryScenarios = map[string][]byte{
	"MysqlBasicQuery":   mysqlBasicQuery,
	"MysqlYcsbWorkload": mysqlYcsbWorkload,
}

//go:embed mysql_basic_query.qst
var mysqlBasicQuery []byte

//go:embed mysql_ycsb_workload.qst
var mysqlYcsbWorkload []byte
