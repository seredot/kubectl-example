FROM golang:1.14

WORKDIR /go/src/kubectl-example
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["kubectl-example"]
