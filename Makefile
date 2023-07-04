GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell cat VERSION)
REVISION=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
TAG=$(shell git describe --tags --always --dirty)
BUILDUSER="$(shell whoami)@$(shell hostname)"
BUILDDATE=$(shell date +"%Y-%m-%d %H:%M:%S")
MODELNAME=$(subst module ,, $(shell cat go.mod | grep "module"))

run:
	echo $(GOPATH)

.PHONY: build
# build
build:
	go build -ldflags \
	"-s -w -X $(MODELNAME)/global/version.Version=$(VERSION) \
	-X $(MODELNAME)/global/version.Tag=$(TAG) \
	-X $(MODELNAME)/global/version.Revision=$(REVISION) \
	-X $(MODELNAME)/global/version.Branch=$(BRANCH) \
	-X $(MODELNAME)/global/version.BuildUser=$(BUILDUSER) \
	-X $(MODELNAME)/global/version.BuildDate=$(BUILDDATE)" \
	-trimpath \
    -o ./ $(MODELNAME)/cmd
