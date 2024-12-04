# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

all: embed.go

embed.go : embed.pl $(wildcard *.qst)
	perl $< > $@

tests: ${CRUD_TESTS} ${ORDER_TESTS} ${LIMIT_TESTS} ${AGGR_TESTS} ${MATH_TESTS} ${YCSB_TESTS} ${ARITH_TESTS}

CRUD_TESTS = \
	smpl_crud_text.qst \
	smpl_crud_int.qst \
	smpl_crud_float.qst \
	smpl_crud_double.qst \
	smpl_crud_datetime.qst

smpl_crud_text.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl text int float double datetime > smpl_crud_text.qst
smpl_crud_int.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl int float double datetime text > smpl_crud_int.qst
smpl_crud_float.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl float double datetime text int > smpl_crud_float.qst
smpl_crud_double.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl double datetime text int float > smpl_crud_double.qst
smpl_crud_datetime.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl datetime text int float double > smpl_crud_datetime.qst

TXN_TESTS = \
	smpl_txn_text.qst \
	smpl_txn_int.qst \
	smpl_txn_float.qst \
	smpl_txn_double.qst \
	smpl_txn_datetime.qst

smpl_txn_text.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl text int float double datetime > smpl_txn_text.qst
smpl_txn_int.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl int float double datetime text > smpl_txn_int.qst
smpl_txn_float.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl float double datetime text int > smpl_txn_float.qst
smpl_txn_double.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl double datetime text int float > smpl_txn_double.qst
smpl_txn_datetime.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl datetime text int float double > smpl_txn_datetime.qst

INDEX_TESTS = \
	smpl_index_text.qst \
	smpl_index_int.qst \
	smpl_index_float.qst \
	smpl_index_double.qst \
	smpl_index_datetime.qst

smpl_index_text.qst: smpl_index.pl data/data_type.pict
	perl smpl_index.pl text int float double datetime > smpl_index_text.qst
smpl_index_int.qst: smpl_index.pl data/data_type.pict
	perl smpl_index.pl int float double datetime text > smpl_index_int.qst
smpl_index_float.qst: smpl_index.pl data/data_type.pict
	perl smpl_index.pl float double datetime text int > smpl_index_float.qst
smpl_index_double.qst: smpl_index.pl data/data_type.pict
	perl smpl_index.pl double datetime text int float > smpl_index_double.qst
smpl_index_datetime.qst: smpl_index.pl data/data_type.pict
	perl smpl_index.pl datetime text int float double > smpl_index_datetime.qst

ORDER_TESTS = \
	select_order_INT.qst \
	select_order_FLOAT.qst \
	select_order_DOUBLE.qst

select_order_INT.qst : select_order.pl data/data_type.pict
	perl select_order.pl INT > select_order_INT.qst 
select_order_FLOAT.qst : select_order.pl data/data_type.pict
	perl select_order.pl FLOAT > select_order_FLOAT.qst 
select_order_DOUBLE.qst : select_order.pl data/data_type.pict
	perl select_order.pl DOUBLE > select_order_DOUBLE.qst 

LIMIT_TESTS = \
	select_limit_INT.qst \
	select_limit_FLOAT.qst \
	select_limit_DOUBLE.qst

select_limit_INT.qst : select_limit.pl data/data_type.pict
	perl select_limit.pl INT > select_limit_INT.qst 
select_limit_FLOAT.qst : select_limit.pl data/data_type.pict
	perl select_limit.pl FLOAT > select_limit_FLOAT.qst 
select_limit_DOUBLE.qst : select_limit.pl data/data_type.pict
	perl select_limit.pl DOUBLE > select_limit_DOUBLE.qst 

AGGR_TESTS = \
	func_aggr_INT.qst \
	func_aggr_FLOAT.qst \
	func_aggr_DOUBLE.qst

func_aggr_INT.qst : func_aggr.pl data/data_type.pict
	perl func_aggr.pl INT > func_aggr_INT.qst 
func_aggr_FLOAT.qst : func_aggr.pl data/data_type.pict
	perl func_aggr.pl FLOAT > func_aggr_FLOAT.qst 
func_aggr_DOUBLE.qst : func_aggr.pl data/data_type.pict
	perl func_aggr.pl DOUBLE > func_aggr_DOUBLE.qst 

MATH_TESTS = \
	func_math_INT.qst \
	func_math_FLOAT.qst \
	func_math_DOUBLE.qst

func_math_INT.qst : func_math.pl data/data_type.pict
	perl func_math.pl INT > func_math_INT.qst 
func_math_FLOAT.qst : func_math.pl data/data_type.pict
	perl func_math.pl FLOAT > func_math_FLOAT.qst 
func_math_DOUBLE.qst : func_math.pl data/data_type.pict
	perl func_math.pl DOUBLE > func_math_DOUBLE.qst 

ARITH_TESTS = \
	update_arith_INT.qst \
	update_arith_FLOAT.qst \
	update_arith_DOUBLE.qst

update_arith_INT.qst : update_arith.pl data/data_type.pict
	perl update_arith.pl INT > update_arith_INT.qst 
update_arith_FLOAT.qst : update_arith.pl data/data_type.pict
	perl update_arith.pl FLOAT > update_arith_FLOAT.qst 
update_arith_DOUBLE.qst : update_arith.pl data/data_type.pict
	perl update_arith.pl DOUBLE > update_arith_DOUBLE.qst 

YCSB_TESTS = \
	ycsb_workload.qst

ycsb_workload.qst: ycsb_workload.pl
	perl ycsb_workload.pl > ycsb_workload.qst

clean:
	rm *.qst
