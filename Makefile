GIT_SHA=$(shell git rev-parse --short --verify HEAD)
GOBUILD=go build -o bin/klusternaut -ldflags "-X github.com/plombardi89/klusternaut/pkg/version.GitSha=${GIT_SHA} -X github.com/plombardi89/klusternaut/pkg/version.KlusternautVersion=${VERSION}"
SHELL_IMAGE=golang:1.10.2

default: compile

all: default install

compile:
	${GOBUILD} main.go

install:
	install -m 0755 bin/klusternaut $(GOPATH)/bin/klusternaut

shell:
	docker run \
	-it \
	-w /go/src/github.com/plombardi89/klusternaut \
	-v ${PWD}:/go/src/github.com/plombardi89/klusternaut \
	--rm ${SHELL_IMAGE} /bin/bash