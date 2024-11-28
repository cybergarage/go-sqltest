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
	perl smpl_crud.pl text > smpl_crud_text.qst
smpl_crud_int.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl int > smpl_crud_int.qst
smpl_crud_float.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl float > smpl_crud_float.qst
smpl_crud_double.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl double > smpl_crud_double.qst
smpl_crud_datetime.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl datetime > smpl_crud_datetime.qst

TXN_TESTS = \
	smpl_txn_text.qst \
	smpl_txn_int.qst \
	smpl_txn_float.qst \
	smpl_txn_double.qst \
	smpl_txn_datetime.qst

smpl_txn_text.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl text > smpl_txn_text.qst
smpl_txn_int.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl int > smpl_txn_int.qst
smpl_txn_float.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl float > smpl_txn_float.qst
smpl_txn_double.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl double > smpl_txn_double.qst
smpl_txn_datetime.qst: smpl_txn.pl data/data_type.pict
	perl smpl_txn.pl datetime > smpl_txn_datetime.qst

ORDER_TESTS = \
	select_order_int.qst \
	select_order_float.qst \
	select_order_double.qst

select_order_int.qst : select_order.pl data/data_type.pict
	perl select_order.pl int > select_order_int.qst 
select_order_float.qst : select_order.pl data/data_type.pict
	perl select_order.pl float > select_order_float.qst 
select_order_double.qst : select_order.pl data/data_type.pict
	perl select_order.pl double > select_order_double.qst 

LIMIT_TESTS = \
	select_limit_int.qst \
	select_limit_float.qst \
	select_limit_double.qst

select_limit_int.qst : select_limit.pl data/data_type.pict
	perl select_limit.pl int > select_limit_int.qst 
select_limit_float.qst : select_limit.pl data/data_type.pict
	perl select_limit.pl float > select_limit_float.qst 
select_limit_double.qst : select_limit.pl data/data_type.pict
	perl select_limit.pl double > select_limit_double.qst 

AGGR_TESTS = \
	func_aggr_int.qst \
	func_aggr_float.qst \
	func_aggr_double.qst

func_aggr_int.qst : func_aggr.pl data/data_type.pict
	perl func_aggr.pl int > func_aggr_int.qst 
func_aggr_float.qst : func_aggr.pl data/data_type.pict
	perl func_aggr.pl float > func_aggr_float.qst 
func_aggr_double.qst : func_aggr.pl data/data_type.pict
	perl func_aggr.pl double > func_aggr_double.qst 

MATH_TESTS = \
	func_math_int.qst \
	func_math_float.qst \
	func_math_double.qst

func_math_int.qst : func_math.pl data/data_type.pict
	perl func_math.pl int > func_math_int.qst 
func_math_float.qst : func_math.pl data/data_type.pict
	perl func_math.pl float > func_math_float.qst 
func_math_double.qst : func_math.pl data/data_type.pict
	perl func_math.pl double > func_math_double.qst 

ARITH_TESTS = \
	update_arith_int.qst \
	update_arith_float.qst \
	update_arith_double.qst

update_arith_int.qst : update_arith.pl data/data_type.pict
	perl update_arith.pl int > update_arith_int.qst 
update_arith_float.qst : update_arith.pl data/data_type.pict
	perl update_arith.pl float > update_arith_float.qst 
update_arith_double.qst : update_arith.pl data/data_type.pict
	perl update_arith.pl double > update_arith_double.qst 

YCSB_TESTS = \
	ycsb_workload.qst

ycsb_workload.qst: ycsb_workload.pl
	perl ycsb_workload.pl > ycsb_workload.qst
