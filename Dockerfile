FROM golang:1.12.9-buster
LABEL maintainer="yauritux@gmail.com"

ENV hour 1

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o bin/ticking_clock cmd/main.go

CMD ["sh", "-c", "/app/bin/ticking_clock ${hour}"]