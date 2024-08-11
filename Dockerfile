FROM --platform=${TARGETPLATFORM:-linux/amd64} golang:1.22-alpine3.19 AS builder

ADD go.mod go.sum /app/
ADD go-hook.go /app/
WORKDIR /app
RUN go mod download
RUN go build -o go-hook go-hook.go

FROM ghcr.io/flant/shell-operator:latest

COPY --from=builder /app/go-hook /
WORKDIR /
ADD ./go-hook /hooks