FROM golang:1.23-alpine AS build

RUN go install github.com/air-verse/air@latest
RUN go clean -cache
WORKDIR /go/src

