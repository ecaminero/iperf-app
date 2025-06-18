GOOSE_BIN=goose
UNAME_S=$(shell uname -s)

.PHONY: logging
logging:
	@echo "UNAME_S: $(UNAME_S)"
	@echo "GOOSE_BIN: $(GOOSE_BIN)"

build:
	GOOS=linux GOARCH=amd64 go build -o dist/main_linux cmd/main.go
