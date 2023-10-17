FROM golang:bookworm as builder

RUN go install github.com/betats/go-vpnsocks/cmd@latest
RUN apt-get update \
 && apt-get install -y tini

FROM debian:bookworm
COPY --from=builder /go/bin/cmd /usr/local/bin/go-vpnsocks
COPY --from=builder /usr/bin/tini /usr/bin/tini
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["/usr/local/bin/entrypoint.sh"]
