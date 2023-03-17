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

PKG_NAME=sqltest

MODULE_ROOT=github.com/cybergarage/go-sqltest

PKG_ROOT=${PKG_NAME}
PKG=\
	${MODULE_ROOT}/${PKG_ROOT}/...

TEST_DIR=test
TEST_HELPER_NAME=embed
TEST_HELPER=\
	${PKG_ROOT}/${TEST_DIR}/${TEST_HELPER_NAME}.go

.PHONY: test format vet lint clean

all: test

${TEST_HELPER} : ${PKG_ROOT}/${TEST_DIR}/${TEST_HELPER_NAME}.pl $(wildcard ${PKG_ROOT}/${TEST_DIR}/*.qst)
	perl $< > $@

format: ${TEST_HELPER}
	gofmt -s -w ${PKG_ROOT}

vet: format
	go vet ${PKG}

lint: format
	golangci-lint run ${PKG_ROOT}/...

test: lint
	go test -v -cover -timeout 60s ${PKG}/...

clean:
	go clean -i ${PKG}