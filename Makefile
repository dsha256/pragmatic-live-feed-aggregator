.PHONY: swag-gen, mock-gen unit-tests, run build-race, build, setenv, help

.DEFAULT_GOAL = help

MAIN_GO="cmd/api/main.go"
MAIN_BIN="bin/main"

swag-gen: ## Generate swagger files.
	swag init -q -g $(MAIN_GO) -o docs/swagger

mock-gen: ## Generate mocked code.
	mockgen -source internal/pragmaticlivefeed/pragmatic_live_feed.go -destination internal/mock/pragmatic_live_feed.go -package mock
	mockgen -source internal/pragmaticlivefeed/repo.go -destination internal/mock/repo.go -package mock

unit-tests: ## Run unit tests in verbose mode.
	go test -v -cover -race ./...

run: build  ## Run the built app's binary file.
	./$(MAIN_BIN)

build-race: ## Build with race detector turned on.
	go build -race -o $(MAIN_BIN) $(MAIN_GO)

build: ## Build.
	go build -race -o $(MAIN_BIN) $(MAIN_GO)

setenv: ## Export all the ENV variables defined in the .env file.
	export $(grep -v '^#' .env | xargs -d '\n')

help: ## Prints this message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
