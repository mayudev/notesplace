FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go .
COPY ./server ./server

RUN go build -v -o /usr/local/bin/app

CMD [ "/usr/local/bin/app" ]

