FROM golang:1.12.9-buster
LABEL maintainer="yauritux@gmail.com"

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o bin/ticking_clock cmd/main.go

CMD ["/app/bin/ticking_clock"]