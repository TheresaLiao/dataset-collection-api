FROM golang:1.11.2-alpine
WORKDIR /apiserver
ADD /golang /apiserver/
RUN cd /apiserver/ && go build -o main
RUN rm -rf /apiserver/
EXPOSE 8080
ENTRYPOINT ./apiserver
