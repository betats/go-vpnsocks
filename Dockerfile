FROM golang:buster as builder

RUN go get github.com/betats/go-vpnsocks/cmd
RUN apt-get update \
 && apt-get install -y tini

FROM debian:buster
COPY --from=builder /go/bin/cmd /usr/local/bin/go-vpnsocks
COPY --from=builder /usr/bin/tini /usr/bin/tini
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["/usr/local/bin/entrypoint.sh"]
