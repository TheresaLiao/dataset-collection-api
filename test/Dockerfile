FROM golang:1.11.2-alpine
WORKDIR /helloworld
ADD main.go /helloworld/
RUN cd /helloworld && go build -o test
RUN rm -rf main.go
EXPOSE 8080
ENTRYPOINT ./test
