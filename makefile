GOLANGCI_LINT_VERSION = v1.57.2
MOCKGEN ?= go.uber.org/mock/mockgen@latest


# 
# CI
# 
.PHONY: ci
ci: lint unit_test

#
# mock files gen
# 
.PHONY: gen
gen: mockgen

.PHONY: mockgen
mockgen:
	go run $(MOCKGEN) -package=mocks -destination mock/httpadapter.go -source adapters/interface.go /

# 
# lint
# 
.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run --allow-parallel-runners

# 
# tests
# 
.PHONY: test
test: unit_test

.PHONY: unit_test
unit_test:
	go test -parallel 6 -race -count=1 -coverpkg=./... -coverprofile=unit_coverage.out -v `go list ./... | grep -v /test/`

