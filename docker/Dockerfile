FROM golang:1.15-alpine

RUN apk update && apk upgrade
RUN apk add --no-cache git
RUN apk add --no-cache ffmpeg
RUN apk add build-base

RUN go get github.com/op/go-logging
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/lib/pq
RUN go get github.com/jinzhu/gorm
RUN go get github.com/stretchr/testify/assert
RUN go get github.com/rs/zerolog

#RUN go get github.com/rylio/ytdl
#WORKDIR /go/src/github.com/rylio/ytdl
#RUN git pull
#RUN git checkout tags/v0.6.2

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

RUN mkdir -p /go/main
ADD /src/api/. /go/main
ADD /src/api/. /go/src/main
WORKDIR /go/main/
RUN cd /go/main/
RUN go build -o main

EXPOSE 22 80
CMD go test
CMD /app/main
