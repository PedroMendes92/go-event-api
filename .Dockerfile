# syntax=docker/dockerfile:1

FROM golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go-event-api

CMD ["/go-event-api"]