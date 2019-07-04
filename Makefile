 # Go parameters
GOCMD=go
GINKGOCMD=ginkgo
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GINKGOCMD)
GOGET=$(GOCMD) get
BINARY_NAME=go-demo
BINARY_PATH=./bin/$(BINARY_NAME)
MAIN_PATH=./cmd/main.go

all: test build

build:
	$(GOBUILD) -o $(BINARY_PATH) -v $(MAIN_PATH)
test:
	$(GOTEST) --v -cover ./...
clean:
	@echo 'Not implemented yet'
# 	$(GOCLEAN)
# 	rm -f $(BINARY_PATH)
run:
	$(GOBUILD) -o $(BINARY_PATH) -v $(MAIN_PATH)
	./$(BINARY_PATH)
deps:
	$(GOGET) -u github.com/joho/godotenv
	$(GOGET) -u github.com/lib/pq
deps-test:
	$(GOGET) -u github.com/google/uuid
	$(GOGET) -u github.com/onsi/ginkgo
	$(GOGET) -u github.com/onsi/gomega

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_PATH) -v $(MAIN_PATH)
build-docker:
	docker run --rm --name go-demo_build -it -v "$(PWD)":/go/src/github.com/BojanKomazec/go-demo -w /go/src/github.com/BojanKomazec/go-demo golang:latest /bin/bash -c "cd cmd; go get; cd ../; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o '$(BINARY_PATH)' -v '$(MAIN_PATH)'"
