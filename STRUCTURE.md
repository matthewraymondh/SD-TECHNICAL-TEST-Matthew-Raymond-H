# Struktur Proyek XYZ Football API

## Ringkasan Lengkap Struktur Direktori

```
xyz-football-api/
│
├── cmd/                                # Aplikasi entry points
│   └── api/
│       └── main.go                     # Main entry point server
│
├── config/                             # Konfigurasi aplikasi
│   └── config.go                       # Load environment variables
│
├── internal/                           # Kode internal aplikasi
│   ├── api/                            # HTTP Layer
│   │   ├── handler/                    # HTTP Request Handlers
│   │   │   ├── auth_handler.go         # Login endpoint
│   │   │   ├── team_handler.go         # Teams CRUD endpoints
│   │   │   ├── player_handler.go       # Players CRUD endpoints
│   │   │   └── match_handler.go        # Matches & Reports endpoints
│   │   ├── middleware/                 # HTTP Middlewares
│   │   │   ├── auth.go                 # JWT Authentication middleware
│   │   │   ├── logger.go               # Request logging middleware
│   │   │   └── cors.go                 # CORS middleware
│   │   └── router.go                   # Routes configuration
│   ├── model/                          # Database Models (GORM)
│   │   ├── team.go                     # Team model
│   │   ├── player.go                   # Player model
│   │   ├── match.go                    # Match model
│   │   └── goal.go                     # Goal model
│   └── repository/                     # Data Access Layer
│       ├── team_repository.go          # Team database operations
│       ├── player_repository.go        # Player database operations
│       ├── match_repository.go         # Match database operations
│       └── goal_repository.go          # Goal database operations
│
├── pkg/                                # Public packages
│   ├── database/                       # Database connection
│   │   └── database.go                 # PostgreSQL init & migration
│   └── utils/                          # Utility functions
│       ├── jwt.go                      # JWT generation & validation
│       └── response.go                 # Standard HTTP responses
│
├── .env.example                        # Environment variables template
├── .gitignore                          # Git ignore file
├── database_init.sql                   # Database initialization script
├── go.mod                              # Go module definition
├── go.sum                              # Go dependencies checksums
├── postman_collection.json             # Postman API collection
├── QUICKSTART.md                       # Quick start guide
└── README.md                           # Complete documentation
```

## Penjelasan Setiap Layer

### 📂 cmd/api/
**Entry Point Aplikasi**
- `main.go`: Inisialisasi server, config, database, dan router

### ⚙️ config/
**Manajemen Konfigurasi**
- Load environment variables dari `.env`
- Validasi konfigurasi aplikasi

### 🔧 internal/api/
**HTTP Layer - Request Handling**

#### handler/
- **auth_handler.go**: Login dan JWT token generation
- **team_handler.go**: Create, Read, Update, Delete teams
- **player_handler.go**: CRUD players dengan validasi jersey number
- **match_handler.go**: Schedule matches, report results, generate reports

#### middleware/
- **auth.go**: Validasi JWT token di setiap protected endpoint
- **logger.go**: Logging setiap HTTP request
- **cors.go**: Cross-Origin Resource Sharing setup

#### router.go
- Setup semua routes (public & protected)
- Binding handlers ke endpoints

### 🗂️ internal/model/
**Database Models dengan GORM**
- Definisi struct yang mapping ke tabel database
- Relasi antar tabel (foreign keys)
- Soft delete support

### 💾 internal/repository/
**Data Access Layer**
- Semua interaksi dengan database melalui GORM
- Query builder untuk operasi kompleks
- Transaction management

### 📦 pkg/database/
**Database Management**
- Inisialisasi koneksi PostgreSQL
- Auto migration GORM
- Database connection pool

### 🛠️ pkg/utils/
**Utility Functions**
- **jwt.go**: Generate & validate JWT tokens
- **response.go**: Standardized JSON responses

## File Pendukung

### .env.example
Template file environment variables untuk konfigurasi aplikasi

### database_init.sql
SQL script untuk:
- Membuat tabel database
- Insert sample data
- Create indexes untuk performa

### postman_collection.json
Koleksi lengkap semua API endpoints untuk testing di Postman

### README.md
Dokumentasi lengkap:
- Instalasi
- Konfigurasi
- API Endpoints
- Contoh penggunaan
- Troubleshooting

### QUICKSTART.md
Panduan cepat untuk menjalankan aplikasi dalam 5 menit

## Pattern Architecture

Proyek ini menggunakan **Clean Architecture** dengan pemisahan layer:

1. **Presentation Layer** (handlers)
   - Menerima HTTP request
   - Validasi input
   - Return HTTP response

2. **Business Logic Layer** (handlers + repositories)
   - Logika bisnis aplikasi
   - Validasi data

3. **Data Access Layer** (repositories)
   - Operasi database
   - Query building

4. **Database Layer** (PostgreSQL + GORM)
   - Persistent storage

## Flow Request

```
HTTP Request
    ↓
Router (router.go)
    ↓
Middleware (auth, logger, cors)
    ↓
Handler (validation, business logic)
    ↓
Repository (database operations)
    ↓
Database (PostgreSQL)
    ↓
Response ← ← ← ← ←
```

## Keamanan

✅ JWT Authentication untuk semua protected endpoints
✅ Password di environment variables (tidak di code)
✅ CORS middleware untuk web security
✅ Input validation di semua endpoints
✅ Soft delete untuk data integrity

## Best Practices Implemented

✅ **Separation of Concerns**: Setiap layer punya tanggung jawab jelas
✅ **Repository Pattern**: Abstraksi database operations
✅ **Dependency Injection**: Loose coupling antar komponen
✅ **Environment Configuration**: Tidak hardcode credentials
✅ **Error Handling**: Standardized error responses
✅ **Logging**: Request logging untuk debugging
✅ **Documentation**: Lengkap dan terstruktur

---

**Struktur ini production-ready dan mudah di-maintain!**
