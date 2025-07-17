FROM golang:1.24 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app .

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3031
CMD ["./app"]
