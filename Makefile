# Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
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

PREFIX?=$(shell pwd)

GIT_ROOT=github.com/cybergarage/
PRODUCT_NAME=go-sqltest
PKG_NAME=sqltest

MODULE_ROOT=${PKG_NAME}
MODULE_PKG_ROOT=${GIT_ROOT}${PRODUCT_NAME}/${MODULE_ROOT}
MODULE_SRC_DIR=${PKG_NAME}
MODULE_SRCS=\
	${MODULE_SRC_DIR}
MODULE_PKGS=\
	${MODULE_PKG_ROOT}

ALL_ROOTS=\
	${MODULE_ROOT}

ALL_SRCS=\
	${MODULE_SRCS}

ALL_PKGS=\
	${MODULE_PKGS}

.PHONY: clean test

all: test

format:
	gofmt -s -w ${ALL_ROOTS}

vet: format
	go vet ${ALL_PKGS}

lint: vet
	golangci-lint run ${ALL_SRCS}

build: vet
	go build -v ${MODULE_PKGS}

test: lint
	go test -v -cover -p=1 ${ALL_PKGS}

clean:
	go clean -i ${ALL_PKGS}
