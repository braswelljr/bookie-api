HASGOCILINT := $(shell which golangci-lint 2> /dev/null)
ifdef HASGOCILINT
    GOLINT=golangci-lint
else
    GOLINT=bin/golangci-lint
endif

.PHONY: run
run:
	encore run

.PHONY: debug
debug: ## Debug encore by running encore daemon
	encore daemon -f

.PHONY: download
download:
	go mod download
	
.PHONY: test
test:
	encore test -race ./...

.PHONY: deploy
deploy:
	git push encore

.PHONY: fix
fix: ## Fix lint violations
	gofmt -s -w .
	goimports -w $$(find . -type f -name '*.go' -not -path "*/vendor/*")

.PHONY: check-makefile
check-makefile:
	cat -e -t -v Makefile

.PHONY: lint
lint: ## Run linters
	$(GOLINT) run