FROM mysql:latest

# Copy init script ke dalam direktori inisialisasi MySQL
COPY build/package/mysql/init.sql /docker-entrypoint-initdb.d/init.sql

# Beri izin eksekusi ke skrip
RUN chmod +x /docker-entrypoint-initdb.d/init.sql
