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
	"FuncMathFloat":     funcMathFloat,
	"SmplTxnTimestamp":  smplTxnTimestamp,
	"SelectLimitFloat":  selectLimitFloat,
	"UpdateArithInt":    updateArithInt,
	"SelectOrderFloat":  selectOrderFloat,
	"SelectOrderDouble": selectOrderDouble,
	"SmplCrudTimestamp": smplCrudTimestamp,
	"SmplCrudText":      smplCrudText,
	"UpdateArithDouble": updateArithDouble,
	"YcsbWorkload":      ycsbWorkload,
	"SelectLimitDouble": selectLimitDouble,
	"SmplTxnFloat":      smplTxnFloat,
	"SmplCrudDouble":    smplCrudDouble,
	"Pgbench":           pgbench,
	"FuncMathDouble":    funcMathDouble,
	"SelectOrderInt":    selectOrderInt,
	"FuncAggrFloat":     funcAggrFloat,
	"SmplTxnDouble":     smplTxnDouble,
	"SmplAlterAdd":      smplAlterAdd,
	"SmplTxnText":       smplTxnText,
	"UpdateArithFloat":  updateArithFloat,
	"SmplTxnInt":        smplTxnInt,
	"FuncAggrDouble":    funcAggrDouble,
	"SelectLimitInt":    selectLimitInt,
	"FuncMathInt":       funcMathInt,
	"SmplCrudInt":       smplCrudInt,
	"FuncAggrInt":       funcAggrInt,
	"SmplCrudFloat":     smplCrudFloat,
}

//go:embed func_math_float.qst
var funcMathFloat []byte

//go:embed smpl_txn_timestamp.qst
var smplTxnTimestamp []byte

//go:embed select_limit_float.qst
var selectLimitFloat []byte

//go:embed update_arith_int.qst
var updateArithInt []byte

//go:embed select_order_float.qst
var selectOrderFloat []byte

//go:embed select_order_double.qst
var selectOrderDouble []byte

//go:embed smpl_crud_timestamp.qst
var smplCrudTimestamp []byte

//go:embed smpl_crud_text.qst
var smplCrudText []byte

//go:embed update_arith_double.qst
var updateArithDouble []byte

//go:embed ycsb_workload.qst
var ycsbWorkload []byte

//go:embed select_limit_double.qst
var selectLimitDouble []byte

//go:embed smpl_txn_float.qst
var smplTxnFloat []byte

//go:embed smpl_crud_double.qst
var smplCrudDouble []byte

//go:embed pgbench.qst
var pgbench []byte

//go:embed func_math_double.qst
var funcMathDouble []byte

//go:embed select_order_int.qst
var selectOrderInt []byte

//go:embed func_aggr_float.qst
var funcAggrFloat []byte

//go:embed smpl_txn_double.qst
var smplTxnDouble []byte

//go:embed smpl_alter_add.qst
var smplAlterAdd []byte

//go:embed smpl_txn_text.qst
var smplTxnText []byte

//go:embed update_arith_float.qst
var updateArithFloat []byte

//go:embed smpl_txn_int.qst
var smplTxnInt []byte

//go:embed func_aggr_double.qst
var funcAggrDouble []byte

//go:embed select_limit_int.qst
var selectLimitInt []byte

//go:embed func_math_int.qst
var funcMathInt []byte

//go:embed smpl_crud_int.qst
var smplCrudInt []byte

//go:embed func_aggr_int.qst
var funcAggrInt []byte

//go:embed smpl_crud_float.qst
var smplCrudFloat []byte
