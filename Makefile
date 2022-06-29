.PHONY: build install;

APP_NAME = webdav
APP_VERSION = 0.2.0

ifeq ($(OS), Windows_NT)
	SHELL = pwsh
	SHELLFLAGS = -Command
	EXECUTABLE ?= ${APP_NAME}.exe
	INSTALL_DIR = C:\Developer\bin
else
	SHELL = bash
	SHELLFLAGS = -c
	EXECUTABLE ?= ${APP_NAME}
	INSTALL_DIR = /usr/local/bin
endif

MAKE_VERSION := $(shell "$(MAKE)" -v | head -n 1)
GIT_COMMIT := $(shell git rev-parse --short HEAD || echo unsupported)
BUILD_TIME := $(shell date --rfc-3339 seconds  | sed -e 's/ /T/' || echo unsupported)

# go tool link --help
# The -w and -s flags reduce binary sizes by excluding unnecessary symbols and debug info
# The -buildid= flag makes builds reproducible
LDFLAGS := "\
-X 'main.GitCommit=$(GIT_COMMIT)' \
-X 'main.Version=$(APP_VERSION)' \
-X 'main.BuildTime=$(BUILD_TIME)' \
-X 'main.MakeVersion=$(MAKE_VERSION)' \
-w -s -buildid=\
"

BUILD_DIR = bin

all: build

build: go-tidy go-build

install: build
	cp $(BUILD_DIR)/$(EXECUTABLE) $(INSTALL_DIR)/$(EXECUTABLE)
	$(EXECUTABLE) -v

go-tidy:
	go mod tidy

go-build:
	go build -trimpath -ldflags $(LDFLAGS) -o $(BUILD_DIR)/$(EXECUTABLE) main.go

go-install:
	go install -trimpath -ldflags $(LDFLAGS)
	$(EXECUTABLE) -v
