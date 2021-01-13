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

# Run golint across all .go files. A confidence interval of 0.3 will not error out when files in the package don't have
# a standard package header comment. If any of the files fail the lint test, then we'll exit with a status of 1. We
# don't want to exit at the first failure, though, because we want all failures to be logged together.
.PHONY: lint-check
lint-check:
	@failed=0; \
	for file in $(GOFILES); do \
		echo $$file; \
		golint -min_confidence 0.3 -set_exit_status $$file || failed=1; \
	done; \
	if [ $$failed -ne 0 ]; then \
		exit 1; \
	fi;

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
		exit 1; \
	fi;
