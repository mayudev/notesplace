FROM node:20-alpine AS frontend
WORKDIR /app
COPY ./web/package*.json ./
RUN npm install
COPY ./web .
RUN npm run build

FROM golang:latest AS backend

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go .
COPY ./server ./server
COPY --from=frontend /app/build ./web

RUN go build -v -o /usr/local/bin/app

EXPOSE 8080
CMD [ "app" ]

