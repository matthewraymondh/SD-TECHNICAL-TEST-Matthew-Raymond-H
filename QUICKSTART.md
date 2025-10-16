# 🚀 Quick Start Guide - XYZ Football API

Panduan cepat untuk menjalankan aplikasi dalam 5 menit!

## Langkah 1: Setup Database

### Buat Database PostgreSQL
```sql
CREATE DATABASE football_db;
```

### (Opsional) Jalankan SQL Script untuk Sample Data
```bash
psql -U postgres -d football_db -f database_init.sql
```

## Langkah 2: Setup Environment

### Copy file .env
```bash
# Windows PowerShell
Copy-Item .env.example .env

# Linux/Mac
cp .env.example .env
```

### Edit file .env
Sesuaikan minimal nilai berikut:
```env
DB_PASSWORD=your_actual_password
JWT_SECRET=your_random_secret_key_min_32_chars
```

## Langkah 3: Install Dependencies

```bash
go mod download
go mod tidy
```

## Langkah 4: Jalankan Aplikasi

```bash
go run cmd/api/main.go
```

**Server akan berjalan di:** `http://localhost:8080`

## Langkah 5: Test API

### 1. Health Check (tanpa autentikasi)
```bash
curl http://localhost:8080/health
```

**Expected Response:**
```json
{
  "status": "ok",
  "message": "XYZ Football API is running"
}
```

### 2. Login untuk mendapatkan Token
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"admin\", \"password\": \"admin123\"}"
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**💡 Simpan token ini untuk request selanjutnya!**

### 3. Test Endpoint Protected (dengan Token)

Ganti `YOUR_TOKEN_HERE` dengan token dari response login:

```bash
# Create Team
curl -X POST http://localhost:8080/teams \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d "{\"name\": \"Test FC\", \"founded_year\": 2024, \"headquarters_city\": \"Jakarta\"}"

# Get All Teams
curl -X GET http://localhost:8080/teams \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## 🎯 Alternatif: Gunakan Postman

1. Import file `postman_collection.json` ke Postman
2. Jalankan request "Login" di folder "Authentication"
3. Token akan otomatis tersimpan di environment variables
4. Semua request lainnya sudah siap digunakan!

## ⚠️ Troubleshooting

### Error: "Failed to connect to database"
- ✅ Pastikan PostgreSQL service running
- ✅ Cek username, password, dan DB name di `.env`
- ✅ Pastikan database `football_db` sudah dibuat

### Error: "Port already in use"
- ✅ Ubah `SERVER_PORT` di `.env` ke port lain (misal: 8081)
- ✅ Atau stop aplikasi yang menggunakan port 8080

### Error: "Token tidak valid"
- ✅ Login ulang untuk mendapatkan token baru
- ✅ Pastikan format header: `Authorization: Bearer <token>`

## 📚 Next Steps

Setelah aplikasi berjalan, Anda bisa:

1. 📖 Baca dokumentasi lengkap di `README.md`
2. 🔍 Eksplorasi semua endpoint di Postman Collection
3. 🗄️ Cek struktur database di `database_init.sql`
4. 💻 Mulai develop dengan membaca struktur kode

## 🎉 Selamat!

Aplikasi XYZ Football API sudah berjalan dengan baik!

**API Base URL:** `http://localhost:8080`
**Default Credentials:**
- Username: `admin`
- Password: `admin123`

---

**Happy Coding! ⚽🚀**
