
# Projec name
PROJECT_NAME=learn-go-backend
PORT=8080

# Binary name
BINARY_NAME=main

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOBINPATH=$(GOBIN)/$(BINARY_NAME)
GOFILES=$(GOBASE)/app/cmd/*.go

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: build clean run daemon deps doctor containerize run-container 

deps:
	@echo "Ensuring dependencies are up to date..."
	@go mod tidy

build: deps
	@echo "Building..."
	@go build -o $(GOBINPATH) $(GOFILES)

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf $(GOBIN)

run: build
	@echo "Running..."
	@$(GOBINPATH)

daemon: build
	@air --build.cmd "go build -o $(GOBINPATH) $(GOFILES)" --build.bin "$(GOBINPATH)"

doctor:
	@echo "Ensuring dependencies are up to date..."
	@go mod tidy

	@echo "Formatting code..."
	@go fmt ./...

	@echo "Running go vet..."
	@go vet ./...

	@echo "Running tests..."
	@go test ./...

	@echo "Running integration tests..."
	@go test -tags integration ./...

	@echo "All checks passed!"

containerize:
	@echo "Building container..."
	@docker build -t $(PROJECT_NAME):latest .
	@echo "Container built!"


run-container: containerize
	@echo "Running container..."
	@docker run --name $(PROJECT_NAME) -d -p $(PORT):$(PORT) $(PROJECT_NAME):latest


dockerhub-push: containerize
	@echo "Pushing container to DockerHub..."
	@docker tag $(PROJECT_NAME):latest $(DOCKERHUB_USERNAME)/$(PROJECT_NAME):latest
	@docker push $(DOCKERHUB_USERNAME)/$(PROJECT_NAME):latest

# Help target
help:
	@echo "Available targets:"
	@echo "  build            	- Build the application"
	@echo "  clean            	- Remove binary and clear cache"
	@echo "  run              	- Build and run the application"
	@echo "  daemon             - Deamon running application"
	@echo "  deps             	- Ensure dependencies are up to date"
	@echo "  doctor           	- Run vet, test, and build"
	@echo "  containerize     	- Build container"
	@echo "  run-container    	- Run container"
	@echo "  dockerhub-push   	- Push container to DockerHub"
	@echo "  help             	- Print this help message"