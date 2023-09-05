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

PICT_TESTS = \
	smpl_crud_text.qst \
	smpl_crud_int.qst \
	smpl_crud_float.qst \
	smpl_crud_double.qst

tests: ${PICT_TESTS}

smpl_crud_text.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl text > smpl_crud_text.qst
	git add smpl_crud_text.qst
	git commit -m "Update smpl_crud_text.qst" smpl_crud_text.qst

smpl_crud_int.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl int > smpl_crud_int.qst
	git add smpl_crud_int.qst
	git commit -m "Update smpl_crud_int.qst" smpl_crud_int.qst

smpl_crud_float.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl float > smpl_crud_float.qst
	git add smpl_crud_float.qst
	git commit -m "Update smpl_crud_float.qst" smpl_crud_float.qst

smpl_crud_double.qst: smpl_crud.pl data/data_type.pict
	perl smpl_crud.pl double > smpl_crud_double.qst
	git add smpl_crud_double.qst
	git commit -m "Update smpl_crud_double.qst" smpl_crud_double.qst

