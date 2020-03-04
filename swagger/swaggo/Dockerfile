FROM golang:1.11.2-alpine

MAINTAINER a60144

RUN apk update && apk upgrade
RUN apk add --no-cache git

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN mkdir -p /app
WORKDIR /app

CMD swag init