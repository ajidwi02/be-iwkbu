# Gunakan base image Go
FROM golang:1.21-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Buat direktori kerja dalam container
WORKDIR /app

# Salin file go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependensi
RUN go mod download

# Salin semua file source code
COPY . .

# Build aplikasi
RUN go build -o main .

# Gunakan image kecil untuk runtime
FROM alpine:latest  
WORKDIR /root/
COPY --from=0 /app/main .

# Port aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
