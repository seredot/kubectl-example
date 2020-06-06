default: build

build:
	go build ./...

run: build
	./sample po
