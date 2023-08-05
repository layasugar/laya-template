GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell cat VERSION)
REVISION=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
TAG=$(shell git describe --tags --always --dirty)
BUILDUSER="$(shell whoami)@$(shell hostname)"
BUILDDATE=$(shell date +"%Y-%m-%d^%H:%M:%S")
MODELNAME=$(subst module ,, $(shell cat go.mod | grep "module"))
PROJECTNAME=$(notdir $(MODELNAME))

run:
	echo $(PROJECTNAME)

.PHONY: build
# build
build:
	go build -ldflags "-s -w -X $(MODELNAME)/pkg/version.Version=$(VERSION) \
		-X $(MODELNAME)/pkg/version.Tag=$(TAG) \
		-X $(MODELNAME)/pkg/version.Revision=$(REVISION) \
		-X $(MODELNAME)/pkg/version.Branch=$(BRANCH) \
		-X $(MODELNAME)/pkg/version.BuildUser=$(BUILDUSER) \
		-X $(MODELNAME)/pkg/version.BuildDate=$(BUILDDATE)" \
	-trimpath \
    -o $(PROJECTNAME) $(MODELNAME)/cmd
