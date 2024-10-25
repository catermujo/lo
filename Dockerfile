
FROM golang:1.23.1

WORKDIR /go/src/github.com/catermujo/wiz

COPY Makefile go.* ./

RUN make tools
