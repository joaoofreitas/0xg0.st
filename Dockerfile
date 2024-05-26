FROM golang:1-alpine AS builder

RUN mkdir /0xg0
WORKDIR /0xg0

COPY . /0xg0/

RUN go build

EXPOSE 80
VOLUME ["/storage"]

ENTRYPOINT  ["./0xg0.st", "-p=80", "-stderrthreshold=INFO"]
