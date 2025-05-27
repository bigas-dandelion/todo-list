FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY api/cmd ./cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-service ./cmd/main.go

FROM scratch

WORKDIR /

COPY --from=builder /app/api-service /api-service

CMD ["/api-service"]