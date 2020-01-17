GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=checkout-client
    
all: dep build
dep:
	$(GOGET) google.golang.org/grpc
	$(GOGET) github.com/stretchr/testify/assert
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v 
test: 
	$(GOTEST) --tags=integration -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)