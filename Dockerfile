FROM golang:1.11.4

WORKDIR /go/src/app
COPY . .

ENV GO111MODULE=on

RUN git config --global http.sslVerify false
RUN git config --global http.postBuffer 1048576000
