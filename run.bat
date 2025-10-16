@echo off
REM ================================================================
REM XYZ Football API - Windows Batch Script untuk Menjalankan Server
REM ================================================================

echo.
echo ╔═══════════════════════════════════════════════════════════╗
echo ║                                                           ║
echo ║        XYZ FOOTBALL API - Server Launcher                ║
echo ║                                                           ║
echo ╚═══════════════════════════════════════════════════════════╝
echo.

REM Check if .env file exists
if not exist .env (
    echo [ERROR] File .env tidak ditemukan!
    echo.
    echo Silakan copy .env.example menjadi .env dan sesuaikan konfigurasi:
    echo     copy .env.example .env
    echo.
    pause
    exit /b 1
)

echo [INFO] Checking Go installation...
go version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Go tidak terinstall atau tidak ada di PATH!
    echo.
    echo Silakan install Go terlebih dahulu dari: https://golang.org/dl/
    echo.
    pause
    exit /b 1
)

echo [OK] Go terinstall
echo.

echo [INFO] Downloading dependencies...
go mod download
if errorlevel 1 (
    echo [ERROR] Gagal download dependencies!
    pause
    exit /b 1
)

echo [OK] Dependencies ready
echo.

echo [INFO] Starting server...
echo ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
echo.

go run cmd/api/main.go

if errorlevel 1 (
    echo.
    echo [ERROR] Server gagal start!
    echo.
    echo Troubleshooting:
    echo 1. Pastikan PostgreSQL sudah running
    echo 2. Cek konfigurasi database di file .env
    echo 3. Pastikan database 'football_db' sudah dibuat
    echo.
    pause
)
