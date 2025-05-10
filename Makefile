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

TEST_SRC_ROOT=test
TEST_PKG=${MODULE_ROOT}/${PKG_NAME}

SCENARIO_ROOT=${PKG_ROOT}/scenarios
SCENARIO_HELPER=${SCENARIO_ROOT}/embed
SCENARIO_MAKEFILE=${SCENARIO_ROOT}/Makefile

.PHONY: test format vet lint clean
.IGNORE: lint

all: test

scenarios: ${SCENARIO_MAKEFILE} ${SCENARIO_HELPER}.go

${SCENARIO_MAKEFILE} : ${SCENARIO_MAKEFILE}.pl $(wildcard ${SCENARIO_ROOT}/data/*.pict)
	perl $< > $@

${SCENARIO_HELPER}.go : ${SCENARIO_HELPER}.pl ${SCENARIO_MAKEFILE} $(wildcard ${SCENARIO_ROOT}/*.qst)
	perl $< > $@

version:
	@pushd ${PKG_ROOT} && ./version.gen > version.go && popd
	-git commit ${PKG_ROOT}/version.go -m "Update version"

format: version
	gofmt -s -w ${PKG_ROOT} ${BIN_ROOT} ${SCENARIO_ROOT}

vet: format
	go vet ${PKG}

lint: format
	golangci-lint run ${PKG_ROOT}/... ${BIN_ROOT}/...

test: scenarios lint
	go test -v -p 1 -timeout 10m -cover -coverpkg=${PKG}/... -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...

test_only:
	go test -v -p 1 -timeout 10m -cover -coverpkg=${PKG}/... -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...

build:
	go build -v -gcflags=${GCFLAGS} ${BINS}

install:
	go install -v -gcflags=${GCFLAGS} ${BINS}

%.md : %.adoc
	asciidoctor -b docbook -a leveloffset=+1 -o - $< | pandoc -t markdown_strict --wrap=none -f docbook > $@

csvs := $(wildcard doc/*/*.csv) $(wildcard ${PKG_ROOT}/data/*.csv)

docs := $(patsubst %.adoc,%.md,$(wildcard doc/*.adoc))

doc_touch: $(csvs)
	touch doc/*.adoc

doc: doc_touch $(docs)

clean:
	go clean -i ${PKG}
