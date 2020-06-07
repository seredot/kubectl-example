FROM golang:1.14

WORKDIR /go/src/kubectl-sample
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["kubectl-sample"]
