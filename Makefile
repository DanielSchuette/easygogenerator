# list everything except for vendored pkgs, lint, test, and install
CMD_DIR := cmd/gensite.go
PKGS := $(shell go list github.com/DanielSchuette/easygogenerator/... | grep -v /vendor)
EXEC := $(BIN_DIR)/gensite.go
BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(EXEC): test windows
	go install $(CMD_DIR)

.PHONY: test
test: lint
	go test $(PKGS)

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter ./... --vendor

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

# create binaries for windows, linux, and darwin and put them inside a newly created bin/
LOCAL_BIN_DIR := bin/gensite.go
VERSION ?= 0.0.1

.PHONY: windows
windows: linux darwin 
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -o $(LOCAL_BIN_DIR)-v$(VERSION)-windows-amd64 $(CMD_DIR)

.PHONY: linux
linux: darwin 
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o $(LOCAL_BIN_DIR)-v$(VERSION)-linux-amd64 $(CMD_DIR)

.PHONY: darwin
darwin: 
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build -o $(LOCAL_BIN_DIR)-v$(VERSION)-darwin-amd64 $(CMD_DIR)
