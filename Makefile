GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell cat VERSION)
REVISION=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
TAG=$(shell git describe --tags --always --dirty)
BUILDUSER="$(shell whoami)@$(shell hostname)"
BUILDDATE=$(shell date +"%Y%m%d-%H:%M.%S")
PNAME=$(shell basename $(shell pwd))
APIPATH=../../api/$(PNAME)

ifeq ($(GOHOSTOS), windows)
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find ./internal -name *.proto)
	API_PROTO_FILES=$(shell find $(APIPATH) -name *.proto)
endif

.PHONY: build
# build
build:
	go build -ldflags \
	"-s -w -X github.com/layasugar/laya-template/global/version.Version=$(VERSION) \
	-X github.com/layasugar/laya-template/global/version.Tag=$(VCSTAG) \
	-X github.com/layasugar/laya-template/global/version.Revision=$(VCSREVISION) \
	-X github.com/layasugar/laya-template/global/version.Branch=$(VCSBRANCH) \
	-X github.com/layasugar/laya-template/global/version.BuildUser=$(BUILDUSER) \
	-X github.com/layasugar/laya-template/global/version.BuildDate=$(BUILDDATE)" \
	-trimpath \
    -o ./ ./cmd/$(PNAME)


.PHONY: build
build:
	make build

.PHONY: build-test
build-test:
	make build
