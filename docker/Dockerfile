FROM golang:1.11.2-alpine

RUN apk update && apk upgrade
RUN apk add --no-cache git

RUN go get github.com/op/go-logging
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/lib/pq
RUN go get github.com/kkdai/youtube

RUN mkdir -p /app
WORKDIR /app

ADD /src/api/. /app

RUN cd /app
RUN go build -o main
# RUN rm -rf /app/main.go

EXPOSE 22 80
CMD /app/main
