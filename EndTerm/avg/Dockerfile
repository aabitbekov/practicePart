FROM golang:1.15.6 AS build-env

COPY . /app

WORKDIR /app


RUN go mod download
RUN go build -o avg ./avg


CMD ["./avg"]
