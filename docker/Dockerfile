FROM golang:1.11.2-alpine

RUN mkdir -p /app
WORKDIR /app

ADD /src/api/. /app
RUN go get github.com/op/go-logging
RUN go get github.com/gin-gonic/gin

RUN go build ./main.go -o main
RUN rm -rf /app/main.go

EXPOSE 22,80
ENTRYPOINT ./app
