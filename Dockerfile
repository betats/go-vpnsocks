FROM golang:buster as builder

RUN go get github.com/betats/go-vpnsocks/cmd

FROM debian:buster
COPY --from=builder /go/bin/cmd /usr/local/bin/go-vpnsocks

ENTRYPOINT ["/usr/local/bin/go-vpnsocks"]
