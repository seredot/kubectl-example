default: build

build:
	go build ./...

run: build
	./sample po

docker-build:
	docker build -t sample .

docker-run:
	docker run --rm -i -t sample pods
	