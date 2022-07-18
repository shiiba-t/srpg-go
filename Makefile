.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
