FROM golang:1.14

WORKDIR /go/src/sample
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["sample"]
