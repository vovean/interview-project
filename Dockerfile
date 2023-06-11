FROM golang:1.18-buster

COPY . .

ENTRYPOINT go run .