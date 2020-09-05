FROM golang:buster as builder

RUN go get github.com/betats/go-vpnsocks/cmd \
 && go build github.com/betats/go-vpnsocks/cmd

FROM debian:buster
COPY --from=builder /go/bin/go-vpnsocks /usr/local/bin/
