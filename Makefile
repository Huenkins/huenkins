NAME=$(shell basename $(CURDIR))
# DIRS := ${sort ${dir $(shell find . | grep "./huenkins/*\\.go")}}
DIRS := ${sort ${dir $(shell find ./huenkins/ | grep "*.go")}}
SUBPACKAGES := ${sort ${dir $(shell find ./huenkins | grep "/*\\.go")}}

LAST_COMMIT := $(shell git log --format="%H" -n 1)
LAST_DATE_COMMIT := $(shell git log -1 --format=%cd)
BUILT_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%S)
VERSION := $(shell cat ./version.txt)
BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)

APP_NAME := huenkins
FULL_NAME := huenkins-huenkins
ARCHIVE_NAME := $(APP_NAME)-$(VERSION).tar.gz
ARCHIVE_FULL_NAME := $(CURDIR)/build/$(ARCHIVE_NAME)

LDFLAGS := '-X "config.CommitDate=$(LAST_DATE_COMMIT)" \
	-X "config.CommitID=$(LAST_COMMIT)" \
	-X "config.Version=$(VERSION)" \
	-X "config.BranchName=$(BRANCH_NAME)" \
	-X "config.BuiltDate=$(BUILT_DATE)" \
	-X "config.AppName=$(FULL_NAME)" \
	-X "config.JenkinsBuildNumber=$(BUILD_NUMBER)" \
	-X "config.JenkinsJobName=$(JOB_NAME)"'

GOBIN := $(CURDIR)/huenkins/go/bin/
GOPATH_ENV := $(GOPATH):$(CURDIR)/huenkins/go/

export GOPATH=$(GOPATH_ENV)

ENV:=GOPATH=$(GOPATH_ENV) GOBIN=$(GOBIN)

##
## List of commands:
##

## default:
all: clean deps fmt lint test build

log:
	@echo "======================================================================"
	@echo 'SUBPACKAGES: ' $(SUBPACKAGES)
	@echo 'DIRS: ' $(DIRS)
	@echo 'GOROOT: ' $(GOROOT)
	@echo 'GOBIN: ' $(GOBIN)
	@echo 'GOPATH: ' $(GOPATH)
	@echo 'ENV: ' $(ENV)


# Remove build and vendor directories
clean:
	@echo "======================================================================"
	@echo 'MAKE: clean: build...'
	@rm -rf build build_name.txt app_name.txt


# Installing build dependencies. You will need to run this once manually when you clone the repo
deps:
	@echo "======================================================================"
	@echo 'MAKE: install...'
	@mkdir -p $(GOBIN)
	$(ENV) go get -v github.com/golang/lint/golint


# Run server on port 8080.
run:
	@echo "======================================================================"
	@echo 'MAKE: run...'
	@$(ENV) go run  -ldflags $(LDFLAGS) ./huenkins/huenkins.go


# Build exe file and suppoting files
build: clean
	@echo "======================================================================"
	@echo 'MAKE: build...'
	$(ENV) CGO_ENABLED=0 GOOS=linux go build -ldflags $(LDFLAGS) -o build/$(APP_NAME) ./huenkins.go
	@chmod +x build/*

# Build exe file for local testing
build-mac: clean
	@echo "======================================================================"
	@echo 'MAKE: build...'
	$(ENV) CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o build/$(APP_NAME) ./huenkins.go
	@chmod 777 build/*

# Full tests
tests: fmt lint test

test:
	@echo "======================================================================"
	@echo "Run race test for " $(SUBPACKAGES)
	@go test  -cover -ldflags $(LDFLAGS) -race $(SUBPACKAGES)

lint:
	@echo "======================================================================"
	@echo "Run golint..."
	for dir in $(SUBPACKAGES); do \
		echo "golint " $$dir; \
		$(GOBIN)golint $$dir/; \
	done
	$(GOBIN)golint ./*.go

fmt:
	@echo "======================================================================"
	@go fmt ./*.go
	for dir in $(SUBPACKAGES); do \
		echo "go fmt " $$dir; \
		go fmt $$dir/*.go; \
	done
