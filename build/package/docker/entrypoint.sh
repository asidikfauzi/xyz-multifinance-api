#!/bin/sh
set -e  # Berhenti jika ada error

echo "📌 Menunggu database siap..."
/wait-for-db.sh

echo "📌 Menjalankan migrasi database..."
make migrate
make seed

echo "📌 Menjalankan aplikasi dengan Air..."
exec air -c .air.toml
