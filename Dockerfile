FROM golang:1.18-alpine

ENV GIN_MODE=release
ENV PORT=3004

WORKDIR /app/calling-bill

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE $PORT