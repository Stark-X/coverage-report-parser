GREEN := \033[32m
REST := \033[0m

APP_NAME := coverage-report-parse

help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary
	go build -o dist/$(APP_NAME)
	@echo "$(GREEN)Build complete, find the binary in dist/$(APP_NAME) $(REST)"

clean: ## clean up the dist directory
	@echo "$(GREEN)Cleaning up...$(REST)"
	rm -rf dist
	@echo "$(GREEN)Cleaned up$(REST)"

.PHONY: build clean help
