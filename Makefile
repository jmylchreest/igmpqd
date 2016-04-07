GITCOMMIT?=$(shell git rev-parse HEAD)
GITDESCRIBE?=$(shell git describe --tags 2>/dev/null || echo nightly)
BUILD_EPOCH?=$(shell date +%s)

all:
	gox -arch="386 amd64 arm" -os="linux darwin windows" \
	-ldflags="-X main.GitCommit=$(GITCOMMIT) -X main.GitDescribe=$(GITDESCRIBE) -X main.BuildTime=$(BUILD_EPOCH)" \
	-output="bin/{{.OS}}/{{.Arch}}/{{.Dir}}"
