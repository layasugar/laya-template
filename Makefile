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


.PHONY: gen-env
# docker容器调试确保环境一样, 制作调试镜像, 思路如下
# 一般我们在docker容器里面进行调试, 这里生产用什么系统这里就用什么系统, 这里我们使用ubuntu22.04
#
gen-env:
	docker run -itd --name gogenenv golang:1.20


.PHONY: debug
# debug
debug:
	docker