FROM golang:1.11.2 as builder

WORKDIR /go/src/github.com/redbubble/hog
COPY . .

ARG VERSION

RUN go get -d -v ./...
RUN go install -v ./...

FROM debian:jessie-slim

ENTRYPOINT ["/usr/local/bin/hog"]

COPY --from=builder /go/bin/hog /usr/local/bin/hog

