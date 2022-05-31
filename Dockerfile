FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go .
COPY ./server ./server

RUN go build -v -o /usr/local/bin/app