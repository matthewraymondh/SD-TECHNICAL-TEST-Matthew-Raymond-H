# XYZ Football API - Backend Management System

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-production--ready-brightgreen)

Backend API lengkap untuk manajemen tim sepak bola amatir, pemain, dan jadwal pertandingan. Dibangun dengan **Go (Golang)**, **Gin Framework**, dan **GORM** dengan database **PostgreSQL**.

> 📚 **Dokumentasi Lengkap**: Lihat [INDEX.md](INDEX.md) untuk navigasi semua dokumentasi proyek

---

## 📋 Daftar Isi

- [Fitur Utama](#-fitur-utama)
- [Teknologi yang Digunakan](#-teknologi-yang-digunakan)
- [Struktur Proyek](#-struktur-proyek)
- [Instalasi](#-instalasi)
- [Konfigurasi](#-konfigurasi)
- [Menjalankan Aplikasi](#-menjalankan-aplikasi)
- [API Endpoints](#-api-endpoints)
- [Contoh Penggunaan](#-contoh-penggunaan)
- [Database Schema](#-database-schema)
- [Testing](#-testing)

---

## ✨ Fitur Utama

### 🔐 Authentication & Security
- JWT (JSON Web Token) untuk autentikasi yang aman
- Middleware protection untuk semua endpoint (kecuali login)
- Token expiration otomatis

### ⚽ Team Management (CRUD)
- Membuat, melihat, memperbarui, dan menghapus tim
- Soft delete untuk data integrity
- Informasi lengkap tim (nama, logo, tahun berdiri, alamat markas)

### 👥 Player Management (CRUD)
- Manajemen pemain per tim
- Validasi nomor punggung unik per tim
- Data fisik pemain (tinggi, berat)
- Posisi pemain (penyerang, gelandang, bertahan, penjaga gawang)
- Soft delete

### 🏆 Match Management
- Penjadwalan pertandingan
- Pelaporan hasil pertandingan dengan detail gol
- Laporan lengkap pertandingan:
  - Skor akhir
  - Top scorer dalam pertandingan
  - Total kemenangan masing-masing tim
  - Status pertandingan (scheduled, completed, cancelled)

### 📊 Match Reporting
- Otomatis menghitung pencetak gol terbanyak
- Statistik kemenangan per tim
- Detail waktu gol (menit ke berapa)

---

## 🛠 Teknologi yang Digunakan

| Teknologi | Versi | Kegunaan |
|-----------|-------|----------|
| Go (Golang) | 1.21+ | Bahasa pemrograman utama |
| Gin-Gonic | Latest | Web framework untuk REST API |
| GORM | Latest | ORM untuk database operations |
| PostgreSQL | 12+ | Database relasional |
| JWT | v5 | Autentikasi dan authorization |
| godotenv | Latest | Environment variable management |

---

## 📁 Struktur Proyek

```
xyz-football-api/
├── cmd/
│   └── api/
│       └── main.go              # Entry point aplikasi
├── config/
│   └── config.go                # Konfigurasi environment
├── internal/
│   ├── api/
│   │   ├── handler/             # HTTP handlers (controllers)
│   │   │   ├── auth_handler.go
│   │   │   ├── team_handler.go
│   │   │   ├── player_handler.go
│   │   │   └── match_handler.go
│   │   ├── middleware/          # Middleware functions
│   │   │   ├── auth.go
│   │   │   ├── logger.go
│   │   │   └── cors.go
│   │   └── router.go            # Route definitions
│   ├── model/                   # GORM models
│   │   ├── team.go
│   │   ├── player.go
│   │   ├── match.go
│   │   └── goal.go
│   └── repository/              # Data access layer
│       ├── team_repository.go
│       ├── player_repository.go
│       ├── match_repository.go
│       └── goal_repository.go
├── pkg/
│   ├── database/                # Database connection
│   │   └── database.go
│   └── utils/                   # Utility functions
│       ├── jwt.go
│       └── response.go
├── .env.example                 # Template environment variables
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

## 🚀 Instalasi

### Prasyarat

Pastikan Anda telah menginstall:

1. **Go (Golang)** versi 1.21 atau lebih baru
   ```bash
   go version
   ```

2. **PostgreSQL** versi 12 atau lebih baru
   ```bash
   psql --version
   ```

3. **Git** untuk clone repository
   ```bash
   git version
   ```

### Langkah Instalasi

1. **Clone repository**
   ```bash
   git clone <repository-url>
   cd xyz-football-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   go mod tidy
   ```

3. **Setup Database PostgreSQL**
   
   Buka PostgreSQL client dan buat database baru:
   ```sql
   CREATE DATABASE football_db;
   ```

4. **Konfigurasi Environment Variables**
   
   Copy file `.env.example` menjadi `.env`:
   ```bash
   # Windows PowerShell
   Copy-Item .env.example .env
   
   # Linux/Mac
   cp .env.example .env
   ```
   
   Edit file `.env` sesuai dengan konfigurasi Anda:
   ```env
   SERVER_PORT=8080

   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password_here
   DB_NAME=football_db
   DB_SSLMODE=disable
   DB_TIMEZONE=Asia/Jakarta

   JWT_SECRET=your_super_secret_key_change_in_production
   JWT_EXPIRATION_HOURS=24

   ADMIN_USERNAME=admin
   ADMIN_PASSWORD=admin123
   ```

5. **Jalankan Aplikasi**
   ```bash
   go run cmd/api/main.go
   ```

   Atau build terlebih dahulu:
   ```bash
   # Build
   go build -o football-api.exe cmd/api/main.go
   
   # Run
   ./football-api.exe
   ```

6. **Verifikasi**
   
   Aplikasi akan berjalan di `http://localhost:8080`
   
   Test health check:
   ```bash
   curl http://localhost:8080/health
   ```

---

## ⚙️ Konfigurasi

### Environment Variables

| Variable | Default | Deskripsi |
|----------|---------|-----------|
| `SERVER_PORT` | 8080 | Port server aplikasi |
| `DB_HOST` | localhost | Host database PostgreSQL |
| `DB_PORT` | 5432 | Port database PostgreSQL |
| `DB_USER` | postgres | Username database |
| `DB_PASSWORD` | - | Password database (REQUIRED) |
| `DB_NAME` | football_db | Nama database |
| `DB_SSLMODE` | disable | SSL mode untuk koneksi DB |
| `DB_TIMEZONE` | Asia/Jakarta | Timezone database |
| `JWT_SECRET` | - | Secret key untuk JWT (REQUIRED) |
| `JWT_EXPIRATION_HOURS` | 24 | Durasi token dalam jam |
| `ADMIN_USERNAME` | admin | Username untuk login |
| `ADMIN_PASSWORD` | admin123 | Password untuk login |

### Database Migration

GORM akan otomatis membuat tabel saat aplikasi pertama kali dijalankan (Auto Migration). Tabel yang dibuat:
- `teams`
- `players`
- `matches`
- `goals`

---

## 🏃 Menjalankan Aplikasi

### Development Mode

```bash
go run cmd/api/main.go
```

### Production Mode

```bash
# Build untuk Windows
go build -o football-api.exe cmd/api/main.go

# Build untuk Linux
GOOS=linux GOARCH=amd64 go build -o football-api cmd/api/main.go

# Build untuk Mac
GOOS=darwin GOARCH=amd64 go build -o football-api cmd/api/main.go

# Run
./football-api
```

### Menggunakan Air (Hot Reload untuk Development)

Install Air:
```bash
go install github.com/cosmtrek/air@latest
```

Jalankan dengan Air:
```bash
air
```

---

## 📡 API Endpoints

### Base URL
```
http://localhost:8080
```

### Authentication

#### 🔓 Login
Mendapatkan JWT token untuk mengakses endpoint yang protected.

**Endpoint:** `POST /login`

**Request Body:**
```json
{
  "username": "admin",
  "password": "admin123"
}
```

**Response Success (200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response Error (401):**
```json
{
  "error": "Username atau password salah"
}
```

---

### Teams Endpoints

> 🔒 **Semua endpoint teams memerlukan Authorization header dengan JWT token**

#### 1. Create Team
**Endpoint:** `POST /teams`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Garuda FC",
  "founded_year": 2020,
  "headquarters_city": "Jakarta",
  "headquarters_address": "Jl. Sudirman No. 1",
  "logo_url": "https://example.com/logo.png"
}
```

**Response Success (201):**
```json
{
  "id": 1,
  "name": "Garuda FC",
  "founded_year": 2020,
  "headquarters_city": "Jakarta",
  "headquarters_address": "Jl. Sudirman No. 1",
  "logo_url": "https://example.com/logo.png",
  "created_at": "2025-10-16T10:30:00Z",
  "updated_at": "2025-10-16T10:30:00Z"
}
```

#### 2. Get All Teams
**Endpoint:** `GET /teams`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response Success (200):**
```json
[
  {
    "id": 1,
    "name": "Garuda FC",
    "founded_year": 2020,
    "headquarters_city": "Jakarta",
    "created_at": "2025-10-16T10:30:00Z",
    "updated_at": "2025-10-16T10:30:00Z"
  },
  {
    "id": 2,
    "name": "Elang FC",
    "founded_year": 2019,
    "headquarters_city": "Bandung",
    "created_at": "2025-10-16T11:00:00Z",
    "updated_at": "2025-10-16T11:00:00Z"
  }
]
```

#### 3. Get Team by ID
**Endpoint:** `GET /teams/:id`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response Success (200):**
```json
{
  "id": 1,
  "name": "Garuda FC",
  "founded_year": 2020,
  "headquarters_city": "Jakarta",
  "headquarters_address": "Jl. Sudirman No. 1",
  "logo_url": "https://example.com/logo.png",
  "created_at": "2025-10-16T10:30:00Z",
  "updated_at": "2025-10-16T10:30:00Z"
}
```

#### 4. Update Team
**Endpoint:** `PUT /teams/:id`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Garuda Perkasa FC",
  "headquarters_address": "Jl. Thamrin No. 5"
}
```

**Response Success (200):**
```json
{
  "id": 1,
  "name": "Garuda Perkasa FC",
  "founded_year": 2020,
  "headquarters_city": "Jakarta",
  "headquarters_address": "Jl. Thamrin No. 5",
  "created_at": "2025-10-16T10:30:00Z",
  "updated_at": "2025-10-16T12:00:00Z"
}
```

#### 5. Delete Team
**Endpoint:** `DELETE /teams/:id`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response Success (200):**
```json
{
  "message": "Team deleted successfully"
}
```

---

### Players Endpoints

> 🔒 **Semua endpoint players memerlukan Authorization header dengan JWT token**

#### 1. Create Player
**Endpoint:** `POST /players`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Budi Santoso",
  "team_id": 1,
  "position": "penyerang",
  "jersey_number": 10,
  "height_cm": 175,
  "weight_kg": 70
}
```

**Posisi yang valid:** `penyerang`, `gelandang`, `bertahan`, `penjaga gawang`

**Response Success (201):**
```json
{
  "id": 1,
  "team_id": 1,
  "name": "Budi Santoso",
  "position": "penyerang",
  "jersey_number": 10,
  "height_cm": 175,
  "weight_kg": 70,
  "created_at": "2025-10-16T10:30:00Z",
  "updated_at": "2025-10-16T10:30:00Z"
}
```

**Response Error (400) - Nomor Punggung Duplikat:**
```json
{
  "error": "Nomor punggung sudah digunakan di tim ini"
}
```

#### 2. Get Players by Team
**Endpoint:** `GET /teams/:id/players`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response Success (200):**
```json
[
  {
    "id": 1,
    "team_id": 1,
    "name": "Budi Santoso",
    "position": "penyerang",
    "jersey_number": 10,
    "height_cm": 175,
    "weight_kg": 70,
    "created_at": "2025-10-16T10:30:00Z",
    "updated_at": "2025-10-16T10:30:00Z"
  },
  {
    "id": 2,
    "team_id": 1,
    "name": "Ahmad Dahlan",
    "position": "gelandang",
    "jersey_number": 8,
    "height_cm": 170,
    "weight_kg": 68,
    "created_at": "2025-10-16T10:35:00Z",
    "updated_at": "2025-10-16T10:35:00Z"
  }
]
```

#### 3. Update Player
**Endpoint:** `PUT /players/:id`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "height_cm": 180,
  "weight_kg": 75
}
```

**Response Success (200):**
```json
{
  "id": 1,
  "team_id": 1,
  "name": "Budi Santoso",
  "position": "penyerang",
  "jersey_number": 10,
  "height_cm": 180,
  "weight_kg": 75,
  "created_at": "2025-10-16T10:30:00Z",
  "updated_at": "2025-10-16T12:00:00Z"
}
```

#### 4. Delete Player
**Endpoint:** `DELETE /players/:id`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response Success (200):**
```json
{
  "message": "Player deleted successfully"
}
```

---

### Matches Endpoints

> 🔒 **Semua endpoint matches memerlukan Authorization header dengan JWT token**

#### 1. Create Match
**Endpoint:** `POST /matches`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "home_team_id": 1,
  "away_team_id": 2,
  "match_datetime": "2025-12-20T19:00:00Z"
}
```

**Response Success (201):**
```json
{
  "id": 1,
  "home_team_id": 1,
  "away_team_id": 2,
  "match_datetime": "2025-12-20T19:00:00Z",
  "status": "scheduled",
  "home_score": 0,
  "away_score": 0,
  "created_at": "2025-10-16T10:30:00Z",
  "updated_at": "2025-10-16T10:30:00Z"
}
```

#### 2. Report Match Result
**Endpoint:** `POST /matches/:id/result`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "home_score": 2,
  "away_score": 1,
  "goals": [
    {
      "player_id": 1,
      "goal_time": 44
    },
    {
      "player_id": 1,
      "goal_time": 78
    },
    {
      "player_id": 5,
      "goal_time": 65
    }
  ]
}
```

**Catatan:** 
- Jumlah elemen dalam array `goals` harus sama dengan `home_score + away_score`
- `goal_time` dalam menit (1-120)
- `player_id` harus dari salah satu tim yang bertanding

**Response Success (200):**
```json
{
  "message": "Match result reported successfully"
}
```

**Response Error (400) - Jumlah Gol Tidak Sesuai:**
```json
{
  "error": "Jumlah gol tidak sesuai. Total skor: 3, jumlah detail gol: 2"
}
```

#### 3. Get Match Report
**Endpoint:** `GET /matches/:id/report`

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response Success (200):**
```json
{
  "schedule": "2025-12-20T19:00:00+07:00",
  "home_team": "Garuda FC",
  "away_team": "Elang FC",
  "final_score": "2-1",
  "match_result": "Tim Home Menang",
  "top_scorer_in_match": "Budi Santoso (2 gol)",
  "home_team_total_wins": 5,
  "away_team_total_wins": 3
}
```

**Kemungkinan nilai `match_result`:**
- `Tim Home Menang` - Skor home > away
- `Tim Away Menang` - Skor away > home
- `Seri` - Skor sama
- `Belum Selesai` - Status masih scheduled

---

## 💡 Contoh Penggunaan

### Menggunakan cURL

#### 1. Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

#### 2. Create Team
```bash
curl -X POST http://localhost:8080/teams \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Garuda FC",
    "founded_year": 2020,
    "headquarters_city": "Jakarta"
  }'
```

#### 3. Get All Teams
```bash
curl -X GET http://localhost:8080/teams \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 4. Create Player
```bash
curl -X POST http://localhost:8080/players \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Budi Santoso",
    "team_id": 1,
    "position": "penyerang",
    "jersey_number": 10
  }'
```

### Menggunakan Postman

1. **Import Collection** (opsional, lihat file `postman_collection.json`)

2. **Setup Environment Variables**
   - Buat environment baru
   - Tambahkan variable `base_url` = `http://localhost:8080`
   - Tambahkan variable `token` = kosongkan dulu

3. **Login Flow**
   - Request ke `/login`
   - Copy token dari response
   - Paste ke environment variable `token`

4. **Gunakan Token**
   - Di Authorization tab, pilih "Bearer Token"
   - Value: `{{token}}`

---

## 🗄 Database Schema

### ERD (Entity Relationship Diagram)

```
┌─────────────┐          ┌──────────────┐          ┌─────────────┐
│   TEAMS     │          │   PLAYERS    │          │    GOALS    │
├─────────────┤          ├──────────────┤          ├─────────────┤
│ id (PK)     │─────────<│ team_id (FK) │          │ id (PK)     │
│ name        │          │ id (PK)      │─────────<│ player_id   │
│ logo_url    │          │ name         │          │ match_id    │
│ founded_year│          │ height_cm    │          │ goal_time   │
│ hq_address  │          │ weight_kg    │          │ created_at  │
│ hq_city     │          │ position     │          └─────────────┘
│ created_at  │          │ jersey_number│                  ʌ
│ updated_at  │          │ created_at   │                  │
│ deleted_at  │          │ updated_at   │                  │
└─────────────┘          │ deleted_at   │                  │
      ʌ                  └──────────────┘                  │
      │                                                     │
      │                  ┌──────────────┐                  │
      │                  │   MATCHES    │                  │
      │                  ├──────────────┤                  │
      └──────────────────│ home_team_id │──────────────────┘
                         │ away_team_id │
                         │ id (PK)      │
                         │ match_datetime
                         │ status       │
                         │ home_score   │
                         │ away_score   │
                         │ created_at   │
                         │ updated_at   │
                         │ deleted_at   │
                         └──────────────┘
```

### SQL Schema

Untuk SQL DDL lengkap, lihat spesifikasi di dokumen requirement.

---

## 🧪 Testing

### Manual Testing dengan cURL

Jalankan file test script:
```bash
# Windows PowerShell
./test-api.ps1

# Linux/Mac
./test-api.sh
```

### Unit Testing (Coming Soon)

```bash
go test ./...
```

### Load Testing dengan Apache Bench

```bash
ab -n 1000 -c 10 http://localhost:8080/health
```

---

## 🔧 Troubleshooting

### Error: "could not connect to database"

**Solusi:**
- Pastikan PostgreSQL service berjalan
- Cek kredensial di file `.env`
- Pastikan database sudah dibuat

### Error: "port already in use"

**Solusi:**
- Ubah `SERVER_PORT` di `.env`
- Atau stop aplikasi yang menggunakan port tersebut

### Error: "token tidak valid atau sudah expired"

**Solusi:**
- Login ulang untuk mendapatkan token baru
- Token expired setelah 24 jam (default)

---

## 📝 Notes

- **Soft Delete**: Semua operasi DELETE tidak menghapus data secara permanen dari database, hanya menandai dengan `deleted_at`
- **Jersey Number**: Nomor punggung harus unik per tim (tidak boleh duplikat dalam satu tim)
- **Match Status**: Status pertandingan otomatis berubah menjadi "completed" setelah hasil dilaporkan
- **JWT Token**: Token valid selama 24 jam (bisa diubah di config)

---

## 🤝 Contributing

Untuk berkontribusi:
1. Fork repository ini
2. Buat branch baru (`git checkout -b feature/AmazingFeature`)
3. Commit perubahan (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

---

## 📄 License

Project ini dilisensikan di bawah MIT License.

---

## 👨‍💻 Author

**Senior Backend Developer**

Jika ada pertanyaan atau masalah, silakan buat issue di repository ini.

---

## 🎯 Roadmap

- [ ] Implement comprehensive unit tests
- [ ] Add pagination untuk list endpoints
- [ ] Implement filtering dan sorting
- [x] Add Docker support ✅ (See [DOCKER.md](DOCKER.md))
- [ ] Add CI/CD pipeline
- [ ] Add Swagger/OpenAPI documentation
- [ ] Implement Redis caching
- [ ] Add WebSocket untuk live match updates

---

## 📚 Additional Documentation

Proyek ini memiliki dokumentasi lengkap untuk berbagai kebutuhan:

- 📑 **[INDEX.md](INDEX.md)** - Navigation hub untuk semua dokumentasi
- 🚀 **[QUICKSTART.md](QUICKSTART.md)** - Setup dalam 5 menit
- 🏗️ **[STRUCTURE.md](STRUCTURE.md)** - Detailed project structure & architecture
- 🧪 **[TESTING.md](TESTING.md)** - Comprehensive testing guide
- 🐳 **[DOCKER.md](DOCKER.md)** - Docker deployment guide
- 📊 **[PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)** - Complete project summary
- 📝 **[CHANGELOG.md](CHANGELOG.md)** - Version history & roadmap
- 🤝 **[CONTRIBUTING.md](CONTRIBUTING.md)** - Contribution guidelines

---

**Happy Coding! ⚽🚀**
