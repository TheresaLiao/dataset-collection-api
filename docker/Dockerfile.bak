
FROM golang:1.11.2-alpine

RUN apk update && apk upgrade
RUN apk add --no-cache git
RUN apk add --no-cache ffmpeg

RUN go get github.com/op/go-logging
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/lib/pq
RUN go get -u github.com/jinzhu/gorm

RUN go get github.com/rylio/ytdl
WORKDIR /go/src/github.com/rylio/ytdl
RUN git pull
RUN git checkout tags/v0.6.1

RUN go get github.com/iawia002/annie
WORKDIR /go/src/github.com/iawia002/annie
RUN git pull
RUN git checkout tags/0.9.8

WORKDIR /
RUN mkdir -p /app
WORKDIR /app

ADD /src/api/. /app

RUN cd /app
RUN go build -o main
RUN rm -rf /app/main.go

EXPOSE 22 80
CMD /app/main
