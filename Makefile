help: ## Show help messages.
	@grep -E '^[0-9a-zA-Z_-]+:(.*?## .*)?$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run="."
dir="./..."
short="-short"
timeout=40s

.PHONY: unit_test
unit_test: ## Run unit tests. You can set: [run, timeout, short, dir, flags]. Example: make unit_test flags="-race".
	@go mod tidy; go test -trimpath --timeout=$(timeout) $(short) $(dir) -run $(run) $(flags)

.PHONY: unit_test_watch
unit_test_watch: ## Run unit tests in watch mode. You can set: [run, timeout, short, dir, flags]. Example: make unit_test flags="-race".
	@echo "running tests on $(run). waiting for changes..."
	@-zsh -c "go mod tidy; go test -trimpath --timeout=$(timeout) $(short) $(dir) -run $(run) $(flags); repeat 100 printf '#'; echo"
	@reflex -d none -r "(\.go$$)|(go.mod)" -- zsh -c "go mod tidy; go test -trimpath --timeout=$(timeout) $(short) $(dir) -run $(run) $(flags); repeat 100 printf '#'"

.PHONY: dependencies
dependencies: ## Install dependencies requried for development operations.
	@go install github.com/cespare/reflex@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
	@go get -u -d ./...
	@go mod tidy

.PHONY: lint
lint: ## Lint the code
	go fmt ./...
	go vet ./...
	golangci-lint run ./...

.PHONY: ci_tests
ci_tests: ## Run tests for CI.
	go test -trimpath --timeout=10m -failfast -v -race -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: clean
clean: ## Clean test caches and tidy up modules.
	@go clean -testcache
	@go mod tidy

.PHONY: mocks
mocks: ## Generate mocks in all packages.
	@go generate ./...

.PHONY: coverage
coverage: ## Show the test coverage on browser.
	go test -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | tail -n 1
	go tool cover -html=coverage.out
