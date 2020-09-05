FROM golang:buster

RUN go get github.com/betats/go-vpnsocks/cmd \
 && go build github.com/betats/go-vpnsocks/cmd


