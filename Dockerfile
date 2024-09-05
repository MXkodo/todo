FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh

RUN go build -o main ./cmd

RUN apk add --no-cache bash

CMD ["bash", "-c", "wait-for-it.sh db:5432 -- ./main"]

EXPOSE 9090
