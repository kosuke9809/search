GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=tmp
MAIN_PATH=cmd/main.go

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
	./$(BINARY_NAME)

deps:
	$(GOMOD) download

lint:
	golangci-lint run

.PHONY: all build test clean run deps lint
