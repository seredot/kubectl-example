default: build

test:
	tests/validate-resources.sh

build:
	go build ./...

run: build
	./kubectl-sample po

build-docker:
	docker build -t kubectl-sample .

run-docker:
	docker run --rm -i -t kubectl-sample pods
	
release: test
	goreleaser release --rm-dist

snapshot: test
	goreleaser release --snapshot --rm-dist