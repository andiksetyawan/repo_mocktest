.PHONY: test unittest

test:
	go test -v -cover -covermode=atomic ./...

unittest:
	go test -short  ./...
