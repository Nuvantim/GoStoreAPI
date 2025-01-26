# Stage 1: Build
FROM golang:1.23-alpine as builder

# Set environment variable
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Buat direktori kerja di container
WORKDIR /app

# Copy go.mod dan go.sum, lalu install dependensi
COPY go.mod go.sum ./
RUN go mod tidy

# Copy seluruh file project ke container
COPY . .

# Build aplikasi
RUN go build -trimpath -ldflags="-s -w" -o /app/bin/main cli/main.go

# Stage 2: Runtime
FROM alpine:latest

# Set direktori kerja
WORKDIR /app

# Salin file .env dari build stage
COPY --from=builder /app/.env /app/.env

# Salin hasil build dari stage builder
COPY --from=builder /app/bin/main /app/main

# Expose port untuk aplikasi
EXPOSE 7373

# Jalankan aplikasi
CMD ["/app/main"]

