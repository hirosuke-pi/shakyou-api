FROM golang:latest

ENV TZ /usr/share/zoneinfo/Asia/Tokyo
ENV GO111MODULE=on

WORKDIR /app
COPY ./app /app

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080
