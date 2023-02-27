VERSION=1.0.0
SERVER_NAME=appCenter-api
ENV=dev

os_name:=$(shell uname -s | tr A-Z a-z)
os_arch:=$(subst x86_64,amd64,$(shell arch | tr A-Z a-z))

mkdir=$(shell mkdir ./build/)
cp_tool=$(shell cp ../../tool/dlv ./build/)

# build with verison infos
constantDir="qbq-open-platform/common/compile"
versionDir="qbq-open-platform/common"
gitTag=$(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T)

ifeq ($(ENV), prod)
	LD_FLAGS=-ldflags "-s -w -X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${constantDir}.ServerName=${SERVER_NAME}"
else
	LD_FLAGS=-ldflags "-X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${constantDir}.ServerName=${SERVER_NAME}"
	GC_FLAGS=-gcflags="all=-N -l"
endif

.PHONY: build
build:
	@echo $(mkdir)
	CGO_ENABLED=0 GOOS=${os_name} GOARCH=${os_arch} go build $(LD_FLAGS) $(GC_FLAGS) -o build/$(SERVER_NAME)
	@echo $(cp_tool)

#.PHONY: build-plugins
#build-plugins:
#	@echo $(mkdir)
#	GOOS=${os_name} GOARCH=${os_arch} go build $(GC_FLAGS) -buildmode=plugin -o build/$(SERVER_NAME)-plugin.so plugins/service.go

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	go clean
	rm -rf build