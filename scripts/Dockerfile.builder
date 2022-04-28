FROM golang:1.17.9-alpine3.15
RUN apk -U add bash git gcc musl-dev make docker-cli curl ca-certificates
WORKDIR /go/src/github.com/rawmind0/api-test