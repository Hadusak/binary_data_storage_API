FROM golang:1.15.3-alpine3.12 as build-env

WORKDIR /go/src/app

ADD . .

WORKDIR /go/src/app/cmd

RUN GOOS=linux C60_ENABLED=0 go build -o /go/bin/main

FROM alpine:3.12

WORKDIR /app
COPY --from=build-env /go/bin/main /app/

ENTRYPOINT ["/app/main"]