FROM golang:buster as builder

RUN go get github.com/betats/go-vpnsocks/cmd

FROM debian:buster
COPY --from=builder /go/bin/cmd /usr/local/bin/go-vpnsocks
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
