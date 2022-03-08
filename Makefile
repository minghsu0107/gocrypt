VERSION ?= latest
ARG ?= --config config-example.yml

##@ General

help:  ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development
fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

test: fmt vet ## Run unit tests.
	go test -v ./...

##@ Build
deps: ## Install dependencies
	go mod download

build: ## Build binary.
	go build gocrypt.go

all: deps test ## Install deps, test and build binary.
	make build

##@ Docker
docker-build: ## Build docker image.
	docker build -t minghsu0107/gocrypt .

docker-pull: ## Pull docker image.
	docker pull minghsu0107/gocrypt:$(VERSION)

docker-run: ## Run docker image.
	docker run -u $(id -u):$(id -g) -v "$(HOME):/home/appuser/" --rm -it minghsu0107/gocrypt:$(VERSION) $(ARG)
