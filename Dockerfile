# Stage 1: Build dengan Go 1.24
FROM golang:1.24 as builder

WORKDIR /app

# Salin dan build
COPY . .
RUN go mod download
RUN go build -o app .

# Stage 2: Image runtime minimal
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 3001
CMD ["./app"]
