# Stage 1: Build dengan Go 1.24
FROM golang:1.24 AS builder

WORKDIR /app

# Salin semua file project ke container builder
COPY . .

# Download dependencies
RUN go mod download

# Build binary bernama 'app' dari package di direktori saat ini
RUN go build -o app .

# Stage 2: Image runtime minimal
FROM debian:bullseye-slim

WORKDIR /app

# Copy hasil build dari builder
COPY --from=builder /app/app .

# Expose port 8080, sesuai aplikasi Go listen di port ini
EXPOSE 8080

# Jalankan binary aplikasi
CMD ["./app"]

