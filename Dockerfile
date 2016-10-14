FROM golang:1.7-alpine
MAINTAINER Toshiaki Maki <makingx at gmail.com>

ADD echo-server /
EXPOSE 8080

ENTRYPOINT ["/echo-server"]
