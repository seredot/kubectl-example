default: build

test:
	tests/validate-resources.sh

build:
	go build ./...

run: build
	./kubectl-example po

build-docker:
	docker build -t kubectl-example .

run-docker:
	docker run --rm -i -t kubectl-example pods
	
release: test
	goreleaser release --rm-dist

snapshot: test
	goreleaser release --snapshot --rm-dist