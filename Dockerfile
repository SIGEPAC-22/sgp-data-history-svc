FROM golang:latest as builder
RUN apt-get update
WORKDIR /go/src
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch
COPY --from=builder /go/src .

ENTRYPOINT  ["/main"]
