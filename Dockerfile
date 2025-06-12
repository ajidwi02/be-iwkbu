# Gunakan base image Golang untuk build
FROM golang:1.21 AS builder

# Set work directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependencies terlebih dahulu
RUN go mod download

# Salin seluruh kode ke image
COPY . .

# Build aplikasi Go
RUN go build -o app .

# Gunakan base image kecil untuk run app
FROM debian:bullseye-slim

# Set working directory
WORKDIR /app

# Salin hasil build dari stage sebelumnya
COPY --from=builder /app/app .

# Expose port aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./app"]