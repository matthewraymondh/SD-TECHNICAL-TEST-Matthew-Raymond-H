# Docker Deployment Guide

Panduan untuk menjalankan XYZ Football API menggunakan Docker.

## 🐳 Prerequisites

- Docker Engine 20.10+
- Docker Compose 2.0+

## 🚀 Quick Start

### 1. Build dan Run dengan Docker Compose

```bash
# Build dan start semua services
docker-compose up -d

# Cek logs
docker-compose logs -f

# Stop services
docker-compose down

# Stop dan hapus volumes
docker-compose down -v
```

### 2. Akses Aplikasi

- **API**: http://localhost:8080
- **PostgreSQL**: localhost:5432

### 3. Test Health Check

```bash
curl http://localhost:8080/health
```