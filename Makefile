OS = $(shell uname | tr A-Z a-z)
export PATH := $(abspath bin/):${PATH}

# Build variables
BUILD_DIR ?= build
#VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
#COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)

#VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || git rev-parse --abbrev-ref HEAD)
COMMIT_COUNT ?= $(shell git rev-list HEAD --count)
COMMIT_HASH ?= $(shell git describe --long --always --dirty 2>/dev/null)
COMMIT_MSG ?= $(shell git log -1 --pretty=%s | head)
AUTHOR ?= $(shell git log --pretty=%an -1 | head)
BRANCH ?= $(shell git name-rev --name-only HEAD | sed -e "s/^remotes\/origin\///")
VERSION ?= ${COMMIT_HASH}.${COMMIT_COUNT}


LDFLAGS += -X 'github.com/wonktnodi/go-services-base/pkg.Version=${VERSION}'
LDFLAGS += -X 'github.com/wonktnodi/go-services-base/pkg.CommitHash=${COMMIT_HASH}'
LDFLAGS += -X 'github.com/wonktnodi/go-services-base/pkg.BuildDate=${BUILD_DATE}'
LDFLAGS += -X 'github.com/wonktnodi/go-services-base/pkg.CommitMsg=${COMMIT_MSG}'
LDFLAGS += -X 'github.com/wonktnodi/go-services-base/pkg.Author=${AUTHOR}'

VERBOSE = 1

export CGO_ENABLED ?= 1
#export GOOS=linux
#export GOARCH=amd64

ifeq (${VERBOSE}, 1)
ifeq ($(filter -v,${GOARGS}),)
	GOARGS += -v
endif
TEST_FORMAT = short-verbose
endif

# Docker variables
DOCKER_TAG ?= ${COMMIT_HASH}-${BRANCH}

# docker registry
DOCKER_REGISTRY ?= "skl.io"

# Dependency versions
GOLANG_VERSION = 1.13

.PHONY: goversion
goversion:
ifneq (${IGNORE_GOLANG_VERSION_REQ}, 1)
	@printf "${GOLANG_VERSION}\n$$(go version | awk '{sub(/^go/, "", $$3);print $$3}')" | sort -t '.' -k 1,1 -k 2,2 -k 3,3 -g | head -1 | grep -q -E "^${GOLANG_VERSION}$$" || (printf "Required Go version is ${GOLANG_VERSION}\nInstalled: `go version`" && exit 1)
endif

.PHONY: build-%
build-%: goversion
ifeq (${VERBOSE}, 1)
	go env
endif
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/$* ./cmd/$*

.PHONY: pack-%
pack-%:
	docker build -t ${DOCKER_REGISTRY}/skl/$*:${DOCKER_TAG} -f ./cmd/$*/Dockerfile .
	docker push ${DOCKER_REGISTRY}/skl/$*:${DOCKER_TAG}

.PHONY: test
test:
	echo ${VERSION}
	echo ${COMMIT_COUNT}
	echo ${COMMIT_HASH}
	echo ${COMMIT_MSG}
	echo ${AUTHOR}
	echo ${BRANCH}

.PHONY: clean
clean: ## Clean builds
	rm -rf ${BUILD_DIR}/

.PHONY: clean-image
clean-image: ## Clean builds
	docker rmi `docker images -f "dangling=true" -q`