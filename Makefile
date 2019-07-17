 # Go parameters
GOCMD=go
GINKGOCMD=ginkgo
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GINKGOTEST=$(GINKGOCMD)
GOGET=$(GOCMD) get
BINARY_NAME=go-demo
BINARY_PATH=./bin/$(BINARY_NAME)
MAIN_PATH=./cmd/main.go

all: test build

build:
	$(GOBUILD) -o $(BINARY_PATH) -v $(MAIN_PATH)
test:
	$(GINKGOTEST) --v ./...
test-cover:
	$(GINKGOTEST) --v -cover ./...
test-bench:
	# To save results for benchcmp tool, stram the output to .txt files:
	#    $ make test-bench > [old,new].txt
	$(GOTEST) --v -run=NONE -bench=. ./...
test-bench-cpuprofile:
	# This command performs CPU profilng for a specified package.
	# It creates CPU profiling report file (named here "cpu.out").
	# pkg is relative path to the package
	# It is set via command line argument. Example:
	#    $ make test-bench-cpuprofile pkg="./internal/pkg/datatypesdemo/"
ifeq ($(pkg), )
	@echo
	@echo ERROR: Package path not specified. Use relative path to cwd.
else
	$(GOTEST) -run=NONE -bench=. -cpuprofile=cpu.out $(pkg)
endif

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
