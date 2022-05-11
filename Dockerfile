FROM golang:latest

ENV TZ /usr/share/zoneinfo/Asia/Tokyo
ENV GO111MODULE=on

WORKDIR /app
COPY ./app /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

EXPOSE 8080
