GOFILES := $(shell find . -name "*.go")
PROJECTS := algorithms data_structures

# Check if any .go files need to be reformatted.
.PHONY: fmt-check
fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "$${diff}"; \
		exit 1; \
	fi;

# Run a large number of linters of various types and purposes across all go source files.
.PHONY: lint-check-source
lint-check-source:
	@if [ ! -f .golangci.yml ]; then \
		echo "Missing .golangci.yml"; \
		exit 1; \
	fi; \
	golangci-lint run --skip-files ".*_test.*";

# Run a large number of linters of various types and purposes across all go files (including test files).
.PHONY: lint-check-all
lint-check-all:
	@if [ ! -f .golangci.yml ]; then \
		echo "Missing .golangci.yml"; \
		exit 1; \
	fi; \
	golangci-lint run;

# Run the tests on every package. If any of the tests fail, then we'll exit with a status of 1. We don't want to exit at
# the first failure, though, because we want all failures to be logged together.
.PHONY: test
test:
	@failed=0; \
	for project in $(PROJECTS); do \
		echo "> $$project"; \
		cd $$project; \
		for package in *; do \
			if [ -d $$package ]; then \
				echo ">> Running tests for $$package"; \
				cd $$package; \
				go test -v ./... || failed=1; \
				cd ..; \
				echo; \
			fi; \
		done; \
		cd ..; \
	done; \
	if [ $$failed -ne 0 ]; then \
		echo "Failed test"; \
		exit 1; \
	fi;
