# syntax=docker/dockerfile:1

FROM golang:1.22.5

WORKDIR /app

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GIN_MODE=release

RUN make build

CMD ["make", "start-prod"]