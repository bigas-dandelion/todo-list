FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY db ./db

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o db-service ./db/cmd/main.go

FROM alpine:3.18

RUN apk --no-cache add postgresql-client

COPY db/wait-for-postgres.sh /usr/local/bin/wait-for-postgres.sh
RUN chmod +x /usr/local/bin/wait-for-postgres.sh

COPY --from=builder /app/db-service /app/db-service

COPY db/.env /app/db/.env

WORKDIR /app

CMD ["wait-for-postgres.sh", "postgres", "./db-service"]