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
	
release:
	goreleaser release --rm-dist

snapshot:
	goreleaser release --snapshot --rm-dist