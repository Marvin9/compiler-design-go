.PHONY: test coverage

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.txt -covermode=atomic