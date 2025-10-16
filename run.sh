#!/bin/bash
# ================================================================
# XYZ Football API - Bash Script untuk Menjalankan Server
# ================================================================

echo ""
echo "╔═══════════════════════════════════════════════════════════╗"
echo "║                                                           ║"
echo "║        XYZ FOOTBALL API - Server Launcher                ║"
echo "║                                                           ║"
echo "╚═══════════════════════════════════════════════════════════╝"
echo ""

# Check if .env file exists
if [ ! -f .env ]; then
    echo "[ERROR] File .env tidak ditemukan!"
    echo ""
    echo "Silakan copy .env.example menjadi .env dan sesuaikan konfigurasi:"
    echo "    cp .env.example .env"
    echo ""
    exit 1
fi

# Check Go installation
echo "[INFO] Checking Go installation..."
if ! command -v go &> /dev/null; then
    echo "[ERROR] Go tidak terinstall atau tidak ada di PATH!"
    echo ""
    echo "Silakan install Go terlebih dahulu dari: https://golang.org/dl/"
    echo ""
    exit 1
fi

echo "[OK] Go terinstall: $(go version)"
echo ""

# Download dependencies
echo "[INFO] Downloading dependencies..."
if ! go mod download; then
    echo "[ERROR] Gagal download dependencies!"
    exit 1
fi

echo "[OK] Dependencies ready"
echo ""

# Start server
echo "[INFO] Starting server..."
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

go run cmd/api/main.go

if [ $? -ne 0 ]; then
    echo ""
    echo "[ERROR] Server gagal start!"
    echo ""
    echo "Troubleshooting:"
    echo "1. Pastikan PostgreSQL sudah running"
    echo "2. Cek konfigurasi database di file .env"
    echo "3. Pastikan database 'football_db' sudah dibuat"
    echo ""
    exit 1
fi
