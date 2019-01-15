FROM golang:1.11.2-alpine

RUN mkdir -p /app
WORKDIR /app



ADD /src/api/. /app

RUN apk add --no-cache git
RUN go get github.com/op/go-logging
RUN go get github.com/gin-gonic/gin

RUN cd /app
RUN go build -o main
# RUN rm -rf /app/main.go

EXPOSE 22 80
CMD /app/main
