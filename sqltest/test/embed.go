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
	"SmplCrudDouble":  smplCrudDouble,
	"FuncAggrInt":     funcAggrInt,
	"FuncAggrFloat":   funcAggrFloat,
	"FuncMathInt":     funcMathInt,
	"SmplAlterAdd":    smplAlterAdd,
	"SmplCrudFloat":   smplCrudFloat,
	"YcsbWorkload":    ycsbWorkload,
	"SmplLimitDouble": smplLimitDouble,
	"SmplOrderFloat":  smplOrderFloat,
	"SmplCrudInt":     smplCrudInt,
	"FuncMathDouble":  funcMathDouble,
	"SmplLimitInt":    smplLimitInt,
	"SmplLimitFloat":  smplLimitFloat,
	"FuncMathFloat":   funcMathFloat,
	"SmplOrderDouble": smplOrderDouble,
	"SmplCrudText":    smplCrudText,
	"Pgbench":         pgbench,
	"FuncAggrDouble":  funcAggrDouble,
	"SmplOrderInt":    smplOrderInt,
}

//go:embed smpl_crud_double.qst
var smplCrudDouble []byte

//go:embed func_aggr_int.qst
var funcAggrInt []byte

//go:embed func_aggr_float.qst
var funcAggrFloat []byte

//go:embed func_math_int.qst
var funcMathInt []byte

//go:embed smpl_alter_add.qst
var smplAlterAdd []byte

//go:embed smpl_crud_float.qst
var smplCrudFloat []byte

//go:embed ycsb_workload.qst
var ycsbWorkload []byte

//go:embed smpl_limit_double.qst
var smplLimitDouble []byte

//go:embed smpl_order_float.qst
var smplOrderFloat []byte

//go:embed smpl_crud_int.qst
var smplCrudInt []byte

//go:embed func_math_double.qst
var funcMathDouble []byte

//go:embed smpl_limit_int.qst
var smplLimitInt []byte

//go:embed smpl_limit_float.qst
var smplLimitFloat []byte

//go:embed func_math_float.qst
var funcMathFloat []byte

//go:embed smpl_order_double.qst
var smplOrderDouble []byte

//go:embed smpl_crud_text.qst
var smplCrudText []byte

//go:embed pgbench.qst
var pgbench []byte

//go:embed func_aggr_double.qst
var funcAggrDouble []byte

//go:embed smpl_order_int.qst
var smplOrderInt []byte
