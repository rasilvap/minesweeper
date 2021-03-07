## This is a self-documented Makefile. For usage information, run `make help`:
##
## For more information, refer to https://suva.sh/posts/well-documented-makefiles/

-include local/Makefile

.PHONY: 

GO = GO111MODULE=on go
GO_FILES ?= ./pkg/...
SH_FILES ?= $(shell find ./scripts -name *.sh)

all: deps build

##@ Dependencies

deps-go: ## Install backend dependencies.
	$(GO) run build.go setup


##@ Building

build-go: ## Build all Go binaries.
	@echo "build go files"
	$(GO) run build.go build

build-server: ## Build obarra server.
	@echo "build server"
	$(GO) run build.go build-server


scripts/go/bin/bra: scripts/go/go.mod
	@cd scripts/go; \
	$(GO) build -o ./bin/bra github.com/unknwon/bra

run: scripts/go/bin/bra ## Build and run web server on filesystem changes.
	@GO111MODULE=on scripts/go/bin/bra run


##@ Testing

test: ## Run tests for backend.
	@echo "test backend"
	$(GO) test -v ./pkg/...

##@ Linting

scripts/go/bin/revive: scripts/go/go.mod
	@cd scripts/go; \
	$(GO) build -o ./bin/revive github.com/mgechev/revive

revive: scripts/go/bin/revive
	@echo "lint via revive"
	@scripts/go/bin/revive \
		-formatter stylish \
		-config ./scripts/go/configs/revive.toml \
		$(GO_FILES)

revive-strict: scripts/go/bin/revive
	@echo "lint via revive (strict)"
	@scripts/revive-strict scripts/go/bin/revive

scripts/go/bin/golangci-lint: scripts/go/go.mod
	@cd scripts/go; \
	$(GO) build -o ./bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

golangci-lint: scripts/go/bin/golangci-lint
	@echo "lint via golangci-lint"
	@scripts/go/bin/golangci-lint run \
		--config ./scripts/go/configs/.golangci.toml \
		$(GO_FILES)

lint: golangci-lint revive revive-strict # Run all code checks for backend.


##@ Docker

build-docker-dev: ## Build Docker image for development (fast).
	@echo "build development container"
	@echo "\033[92mInfo:\033[0m the frontend code is expected to be built already."
	$(GO) run build.go -goos linux -pkg-arch amd64 ${OPT} build pkg-archive latest
	cp dist/obarra-latest.linux-x64.tar.gz packaging/docker
	cd packaging/docker && docker build --tag obarra/obarra:dev .

build-docker-full: ## Build Docker image for development.
	@echo "build docker container"
	docker build --tag obarra/obarra:dev .
