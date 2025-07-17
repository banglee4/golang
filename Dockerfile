# Stage 1: Build dengan Go 1.24
FROM golang:1.24 AS builder

WORKDIR /golang

# Salin semua file project ke container builder
COPY . .

# Download dependencies
RUN go mod download

# Build binary bernama 'app' dari package di direktori saat ini
RUN go build -o modulv .

# Stage 2: Image runtime minimal
FROM debian:bullseye-slim

WORKDIR /golang

# Copy hasil build dari builder
COPY --from=builder /golang/modul .

# Expose port 8080, sesuai aplikasi Go listen di port ini
EXPOSE 3031

# Jalankan binary aplikasi
CMD ["./modulv"]

