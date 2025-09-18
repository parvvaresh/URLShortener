FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o url-shortener ./cmd/server

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/url-shortener .
COPY .env .

EXPOSE 8080

CMD ["./url-shortener"]
