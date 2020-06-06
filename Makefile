default: build

build:
	go build ./...

run: build
	./sample po

build-docker:
	docker build -t sample .

run-docker:
	docker run --rm -i -t sample pods
	