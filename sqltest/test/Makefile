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

embed.go : embed.pl tests $(wildcard *.qst)
	perl $< > $@

tests: ${PICT_TESTS} ${AGGR_TESTS} ${MATH_TESTS} ${YCSB_TESTS}

PICT_TESTS = \
	smpl_crud_text.qst \
	smpl_crud_int.qst \
	smpl_crud_float.qst \
	smpl_crud_double.qst

smpl_crud_text.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl text > smpl_crud_text.qst
smpl_crud_int.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl int > smpl_crud_int.qst
smpl_crud_float.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl float > smpl_crud_float.qst
smpl_crud_double.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl double > smpl_crud_double.qst
AGGR_TESTS = \
	func_aggr_int.qst \
	func_aggr_float.qst \
	func_aggr_double.qst

func_aggr_int.qst: func_aggr.pl
	perl func_aggr.pl INT > func_aggr_int.qst
func_aggr_float.qst: func_aggr.pl
	perl func_aggr.pl FLOAT > func_aggr_float.qst
func_aggr_double.qst: func_aggr.pl
	perl func_aggr.pl DOUBLE > func_aggr_double.qst
MATH_TESTS = \
	func_math_int.qst \
	func_math_float.qst \
	func_math_double.qst

func_math_int.qst: func_math.pl
	perl func_math.pl INT > func_math_int.qst
func_math_float.qst: func_math.pl
	perl func_math.pl FLOAT > func_math_float.qst
func_math_double.qst: func_math.pl
	perl func_math.pl DOUBLE > func_math_double.qst

YCSB_TESTS = \
	ycsb_workload.qst

ycsb_workload.qst: ycsb_workload.pl
	perl ycsb_workload.pl > ycsb_workload.qst
