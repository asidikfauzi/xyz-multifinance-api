FROM golang:1.23-alpine

WORKDIR /app

# Install dependency dalam satu perintah RUN untuk mengurangi layer image
RUN apk add --no-cache make && \
    go install github.com/air-verse/air@latest

# Copy go.mod dan go.sum lalu install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Tambahkan skrip entrypoint dan beri izin eksekusi
COPY build/package/docker/wait-for-db.sh /wait-for-db.sh
COPY build/package/docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /wait-for-db.sh /entrypoint.sh

# Gunakan exec form pada CMD
CMD ["/entrypoint.sh"]
