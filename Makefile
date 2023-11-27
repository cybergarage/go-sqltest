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
PKG_COVER=${PKG_NAME}-cover
PKG=${MODULE_ROOT}/${PKG_ROOT}

BIN_ROOT=cmd
BIN_ID=${MODULE_ROOT}/${BIN_ROOT}
BIN_CMD=${PKG_NAME}
BIN_CMD_ID=${BIN_ID}/${BIN_CMD}
BIN_SRCS=\
        ${BIN_ROOT}/${BIN_CMD}
BINS=\
        ${BIN_CMD_ID}

TEST_ROOT=${PKG_ROOT}/test
TEST_PKG=${MODULE_ROOT}/${TEST_ROOT}
TEST_HELPER=${TEST_ROOT}/embed
TEST_MAKEFILE=${TEST_ROOT}/Makefile

.PHONY: test format vet lint clean

all: test

scenarios: ${TEST_MAKEFILE} ${TEST_HELPER}.go

${TEST_MAKEFILE} : ${TEST_MAKEFILE}.pl $(wildcard ${TEST_ROOT}/data/*.pict)
	perl $< > $@

${TEST_HELPER}.go : ${TEST_HELPER}.pl ${TEST_MAKEFILE} $(wildcard ${TEST_ROOT}/*.qst)
	perl $< > $@

format:
	gofmt -s -w ${PKG_ROOT} ${BIN_ROOT} ${TEST_ROOT}

vet: format
	go vet ${PKG}

lint: format
	golangci-lint run ${PKG_ROOT}/... ${BIN_ROOT}/...

test: scenarios lint
	go test -v -p 1 -timeout 10m -cover -coverpkg=${PKG} -coverprofile=${PKG_COVER}.out ${PKG}/.. ${TEST_PKG}/..

build:
	go build -v -gcflags=${GCFLAGS} ${BINS}

install:
	go install -v -gcflags=${GCFLAGS} ${BINS}

clean:
	go clean -i ${PKG}

