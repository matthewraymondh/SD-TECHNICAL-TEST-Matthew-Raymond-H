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

## 📦 Manual Docker Build

### Build Image

```bash
docker build -t football-api:latest .
```

### Run Container

```bash
# Run PostgreSQL
docker run -d \
  --name football-db \
  -e POSTGRES_PASSWORD=postgres123 \
  -e POSTGRES_DB=football_db \
  -p 5432:5432 \
  postgres:15-alpine

# Run API
docker run -d \
  --name football-api \
  --link football-db:postgres \
  -e DB_HOST=postgres \
  -e DB_PASSWORD=postgres123 \
  -e JWT_SECRET=your_secret_key \
  -p 8080:8080 \
  football-api:latest
```

## ⚙️ Environment Variables

Edit `docker-compose.yml` untuk mengubah konfigurasi:

```yaml
environment:
  SERVER_PORT: 8080
  DB_HOST: postgres
  DB_PASSWORD: your_secure_password
  JWT_SECRET: your_random_secret_key
```

## 🔍 Useful Commands

### View Logs
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f api
docker-compose logs -f postgres
```

### Access Container Shell
```bash
# API container
docker-compose exec api sh

# PostgreSQL container
docker-compose exec postgres psql -U postgres -d football_db
```

### Restart Services
```bash
# Restart all
docker-compose restart

# Restart specific service
docker-compose restart api
```

### Check Status
```bash
docker-compose ps
```

## 🗄️ Database Management

### Access PostgreSQL
```bash
docker-compose exec postgres psql -U postgres -d football_db
```

### Run SQL Queries
```sql
-- Inside PostgreSQL container
SELECT * FROM teams;
SELECT * FROM players;
SELECT * FROM matches;
```

### Backup Database
```bash
docker-compose exec postgres pg_dump -U postgres football_db > backup.sql
```

### Restore Database
```bash
docker-compose exec -T postgres psql -U postgres football_db < backup.sql
```

## 🔧 Development Mode

### Hot Reload dengan Volume Mount

Edit `docker-compose.yml`:

```yaml
api:
  volumes:
    - .:/app
  command: go run cmd/api/main.go
```

## 🚀 Production Deployment

### Best Practices:

1. **Use Secrets for Sensitive Data**
```yaml
secrets:
  db_password:
    external: true
  jwt_secret:
    external: true
```

2. **Resource Limits**
```yaml
api:
  deploy:
    resources:
      limits:
        cpus: '1'
        memory: 512M
      reservations:
        cpus: '0.5'
        memory: 256M
```

3. **Health Checks**
```yaml
api:
  healthcheck:
    test: ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost:8080/health"]
    interval: 30s
    timeout: 10s
    retries: 3
```

## 🐛 Troubleshooting

### Issue: Container Won't Start

**Check logs:**
```bash
docker-compose logs api
```

**Common causes:**
- Database not ready → Wait for postgres healthcheck
- Port already in use → Change port mapping
- Environment variables missing → Check .env or docker-compose.yml

### Issue: Cannot Connect to Database

**Verify network:**
```bash
docker-compose exec api ping postgres
```

**Check PostgreSQL logs:**
```bash
docker-compose logs postgres
```

### Issue: Permission Denied

**Fix permissions:**
```bash
chmod +x run.sh
docker-compose down
docker-compose up -d --build
```

## 📊 Monitoring

### View Resource Usage
```bash
docker stats
```

### Container Inspect
```bash
docker inspect football-api
docker inspect football-db
```

## 🔄 Update & Rebuild

### After Code Changes:
```bash
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

### Update Only API:
```bash
docker-compose up -d --build api
```

## 🌐 Network Configuration

Default network: `football-network`

Services can communicate using service names:
- API → Database: `postgres:5432`

## 📝 Notes

- **Data Persistence**: Database data stored in `postgres_data` volume
- **Auto Migration**: GORM auto migration runs on startup
- **Sample Data**: Loaded from `database_init.sql` on first run
- **Time Zone**: Set to `Asia/Jakarta`

---

**Happy Deploying! 🐳⚽**
