# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy ทั้งโปรเจกต์
COPY . .

# ใช้ go mod tidy แทน download
RUN go mod tidy

# Build binary
RUN go build -o app ./cmd/api/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 5000

CMD ["./app"]
