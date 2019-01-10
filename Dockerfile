FROM golang:1.11.2-alpine
WORKDIR /api
ADD /golang /api/
RUN cd /api/golang/ && go build -o main
RUN rm -rf /api/golang
EXPOSE 8080
ENTRYPOINT ./api
