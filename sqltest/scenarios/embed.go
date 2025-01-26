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
var EmbedScenarios = map[string][]byte{
	"UpdateArithInt":    updateArithInt,
	"SmplTxnText":       smplTxnText,
	"SmplCrudDouble":    smplCrudDouble,
	"FuncAggrInt":       funcAggrInt,
	"SelectOrderFloat":  selectOrderFloat,
	"SmplTxnDouble":     smplTxnDouble,
	"FuncAggrFloat":     funcAggrFloat,
	"SelectLimitFloat":  selectLimitFloat,
	"FuncMathInt":       funcMathInt,
	"SmplAlterAdd":      smplAlterAdd,
	"SmplCrudFloat":     smplCrudFloat,
	"UpdateArithDouble": updateArithDouble,
	"SmplIndexText":     smplIndexText,
	"SmplTxnFloat":      smplTxnFloat,
	"SelectOrderInt":    selectOrderInt,
	"SmplTxnInt":        smplTxnInt,
	"YcsbWorkload":      ycsbWorkload,
	"SmplCrudDatetime":  smplCrudDatetime,
	"SelectLimitInt":    selectLimitInt,
	"SelectLimitDouble": selectLimitDouble,
	"SmplIndexInt":      smplIndexInt,
	"SmplCrudInt":       smplCrudInt,
	"SmplIndexFloat":    smplIndexFloat,
	"FuncMathDouble":    funcMathDouble,
	"SmplIndexDouble":   smplIndexDouble,
	"SmplIndexDatetime": smplIndexDatetime,
	"FuncMathFloat":     funcMathFloat,
	"SelectOrderDouble": selectOrderDouble,
	"UpdateArithFloat":  updateArithFloat,
	"SmplCrudText":      smplCrudText,
	"Pgbench":           pgbench,
	"FuncAggrDouble":    funcAggrDouble,
	"SmplTxnDatetime":   smplTxnDatetime,
}

//go:embed update_arith_int.qst
var updateArithInt []byte

//go:embed smpl_txn_text.qst
var smplTxnText []byte

//go:embed smpl_crud_double.qst
var smplCrudDouble []byte

//go:embed func_aggr_int.qst
var funcAggrInt []byte

//go:embed select_order_float.qst
var selectOrderFloat []byte

//go:embed smpl_txn_double.qst
var smplTxnDouble []byte

//go:embed func_aggr_float.qst
var funcAggrFloat []byte

//go:embed select_limit_float.qst
var selectLimitFloat []byte

//go:embed func_math_int.qst
var funcMathInt []byte

//go:embed smpl_alter_add.qst
var smplAlterAdd []byte

//go:embed smpl_crud_float.qst
var smplCrudFloat []byte

//go:embed update_arith_double.qst
var updateArithDouble []byte

//go:embed smpl_index_text.qst
var smplIndexText []byte

//go:embed smpl_txn_float.qst
var smplTxnFloat []byte

//go:embed select_order_int.qst
var selectOrderInt []byte

//go:embed smpl_txn_int.qst
var smplTxnInt []byte

//go:embed ycsb_workload.qst
var ycsbWorkload []byte

//go:embed smpl_crud_datetime.qst
var smplCrudDatetime []byte

//go:embed select_limit_int.qst
var selectLimitInt []byte

//go:embed select_limit_double.qst
var selectLimitDouble []byte

//go:embed smpl_index_int.qst
var smplIndexInt []byte

//go:embed smpl_crud_int.qst
var smplCrudInt []byte

//go:embed smpl_index_float.qst
var smplIndexFloat []byte

//go:embed func_math_double.qst
var funcMathDouble []byte

//go:embed smpl_index_double.qst
var smplIndexDouble []byte

//go:embed smpl_index_datetime.qst
var smplIndexDatetime []byte

//go:embed func_math_float.qst
var funcMathFloat []byte

//go:embed select_order_double.qst
var selectOrderDouble []byte

//go:embed update_arith_float.qst
var updateArithFloat []byte

//go:embed smpl_crud_text.qst
var smplCrudText []byte

//go:embed pgbench.qst
var pgbench []byte

//go:embed func_aggr_double.qst
var funcAggrDouble []byte

//go:embed smpl_txn_datetime.qst
var smplTxnDatetime []byte
