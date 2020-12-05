FROM golang:1.14-alpine

COPY data.sql /docker-entrypoint-initdb.d/

RUN mkdir -p go/src/app
WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -v