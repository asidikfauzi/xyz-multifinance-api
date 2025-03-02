FROM golang:1.23-alpine AS builder

RUN apk add --no-cache make bash

WORKDIR /go/src/xyz-multifinance-api

COPY ../../../ ./

# Tambahkan skrip wait-for-db.sh
COPY build/package/docker/wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

# Jalankan skrip untuk menunggu MySQL sebelum migrasi
RUN /wait-for-db.sh && make migrate

RUN make seed
RUN make all

FROM alpine:3.18.4
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/xyz-multifinance-api/bin/multifinance ./
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /

# Set timezone dan environment
ENV TZ=Asia/Jakarta
ENV ZONEINFO=/zoneinfo.zip

# Expose port 8080
EXPOSE 8080

# Gunakan path absolut untuk entrypoint
ENTRYPOINT ["./multifinance"]
