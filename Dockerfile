FROM golang:1.9-alpine

RUN apk update
RUN apk add git
RUN go get -u github.com/alecthomas/gometalinter
RUN gometalinter --install

RUN mkdir -p /go/src/github.com/agrea/go-test-runner
WORKDIR /go/src/github.com/agrea/go-test-runner

COPY *.go .

RUN go build .
RUN go install .

CMD ["go-test-runner"]
