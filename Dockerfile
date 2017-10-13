FROM golang:1.9

RUN go get -u gopkg.in/alecthomas/gometalinter.v1 && \
    gometalinter.v1 --install

RUN mkdir -p /go/src/github.com/agrea/go-test-runner
WORKDIR /go/src/github.com/agrea/go-test-runner

COPY *.go .

RUN go build . && go install .

CMD ["go-test-runner"]
