help: ## Show help messages.
	@grep -E '^[0-9a-zA-Z_-]+:(.*?## .*)?$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run="."
dir="./..."
short="-short"
tag="v0.0.1"


.PHONY: unittest
unittest: ## Run unit tests.
	@echo "running tests on $(run)."
	@go test --timeout=30s --failfast $(short) $(dir) -run $(run)


.PHONY: unittest_watch
unittest_watch: ## Run unit tests in watch mode.
	@echo "running tests on $(run). waiting for changes..."
	@-zsh -c "go test --timeout=30s $(short) $(dir) -run $(run); repeat 100 printf '#'; echo"
	@reflex -d none -r "(\.go$$)|(go.mod)" -- zsh -c "go test --timeout=30s --failfast $(short) $(dir) -run $(run); repeat 100 printf '#'"


.PHONY: unittest_race
unittest_race: ## Run unit tests with race flag.
	@echo "running tests on $(run)."
	@go test --timeout=30s $(short) $(dir) -run $(run)

.PHONY: unittest_race_watch
unittest_race_watch: ## Run unit test with race flag in watch mode.
	@echo "running tests on $(run) with race flag. waiting for changes..."
	@-zsh -c "go test --timeout=30s --failfast $(short) $(dir) -race -run $(run); repeat 100 printf '#'; echo"
	@reflex -d none -r "(\.go$$)|(go.mod)" -- zsh -c "go test --timeout=30s --failfast $(short) $(dir) -race -run $(run); repeat 100 printf '#'"


.PHONY: dependencies
dependencies: ## Install dependencies requried for development operations.
	@go get -u github.com/cespare/reflex
	@go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
	@go get github.com/stretchr/testify/mock
	@go get github.com/vektra/mockery/.../
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
	@go mod tidy


.PHONY: changelog
changelog: ## Update the changelog.
	@git-chglog > CHANGELOG.md
	@echo "Changelog has been updated."


.PHONY: changelog_release
changelog_release: ## Update the changelog with a release tag.
	@git-chglog --next-tag $(tag) > CHANGELOG.md
	@echo "Changelog has been updated."


.PHONY: clean
clean: ## Clean test caches and tidy up modules.
	@go clean -testcache
	@go mod tidy


.PHONY: mocks
mocks: ## Generate mocks in all packages.
	@go generate ./...
