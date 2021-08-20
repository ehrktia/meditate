GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD)fmt
GOCLEAN=$(GOCMD) clean
BINARY_NAME?= meditate
   
.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
.PHONY: test
test:
	$(GOTEST) -race -covermode=atomic -coverprofile=coverage.out ./...
.PHONY: format
format:
	$(GOFMT) -s -w .
.PHONY: lint
lint:
	golint ./...
	golangci-lint run
.PHONY: sec
sec:
	gosec ./...
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: check
check: format lint test
