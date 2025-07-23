# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go build -o site-monitor

# Runtime stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/site-monitor .

CMD ["./site-monitor"]
