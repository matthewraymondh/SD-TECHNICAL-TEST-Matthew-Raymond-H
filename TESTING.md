# 🧪 API Testing Guide

Panduan lengkap untuk testing XYZ Football API.

---

## 📋 Prerequisites

Sebelum testing, pastikan:
- ✅ Server sudah berjalan di `http://localhost:8080`
- ✅ Database PostgreSQL sudah setup
- ✅ File `.env` sudah dikonfigurasi dengan benar

---

## 🔍 Testing Tools

### Option 1: cURL (Command Line)
Cocok untuk quick testing dan automation.

### Option 2: Postman
Cocok untuk comprehensive testing dan debugging.

### Option 3: Browser (untuk GET endpoints)
Cocok untuk quick check.

---

## 🚀 Testing Flow

### Step 1: Health Check ✓

Test bahwa server berjalan dengan baik.

**cURL:**
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

---

### Step 2: Login & Get Token 🔐

**cURL:**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"admin\", \"password\": \"admin123\"}"
```

**Expected Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzI5MTUyMDAwLCJpYXQiOjE3MjkwNjU2MDAsIm5iZiI6MTcyOTA2NTYwMH0.xxx"
}
```

**⚠️ PENTING:** Simpan token ini untuk digunakan di request berikutnya!

**Set Token sebagai Variable (Bash):**
```bash
export TOKEN="your_token_here"
```

**Set Token sebagai Variable (PowerShell):**
```powershell
$TOKEN = "your_token_here"
```

---

## ⚽ Testing Teams Endpoints

### 1. Create Team

**cURL (Bash):**
```bash
curl -X POST http://localhost:8080/teams \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Garuda FC",
    "founded_year": 2020,
    "headquarters_city": "Jakarta",
    "headquarters_address": "Jl. Sudirman No. 1",
    "logo_url": "https://example.com/garuda.png"
  }'
```

**cURL (PowerShell):**
```powershell
curl -X POST http://localhost:8080/teams `
  -H "Authorization: Bearer $TOKEN" `
  -H "Content-Type: application/json" `
  -d '{"name": "Garuda FC", "founded_year": 2020, "headquarters_city": "Jakarta"}'
```

**Expected Response (201):**
```json
{
  "id": 1,
  "name": "Garuda FC",
  "founded_year": 2020,
  "headquarters_city": "Jakarta",
  "headquarters_address": "Jl. Sudirman No. 1",
  "logo_url": "https://example.com/garuda.png",
  "created_at": "2025-10-16T12:00:00+07:00",
  "updated_at": "2025-10-16T12:00:00+07:00"
}
```

### 2. Get All Teams

**cURL:**
```bash
curl -X GET http://localhost:8080/teams \
  -H "Authorization: Bearer $TOKEN"
```

**Expected Response (200):**
```json
[
  {
    "id": 1,
    "name": "Garuda FC",
    "founded_year": 2020,
    "headquarters_city": "Jakarta",
    "created_at": "2025-10-16T12:00:00+07:00",
    "updated_at": "2025-10-16T12:00:00+07:00"
  }
]
```

### 3. Get Team by ID

**cURL:**
```bash
curl -X GET http://localhost:8080/teams/1 \
  -H "Authorization: Bearer $TOKEN"
```

### 4. Update Team

**cURL:**
```bash
curl -X PUT http://localhost:8080/teams/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Garuda Perkasa FC",
    "headquarters_address": "Jl. Thamrin No. 5"
  }'
```

### 5. Delete Team

**cURL:**
```bash
curl -X DELETE http://localhost:8080/teams/1 \
  -H "Authorization: Bearer $TOKEN"
```

**Expected Response (200):**
```json
{
  "message": "Team deleted successfully"
}
```

---

## 👥 Testing Players Endpoints

### 1. Create Team First
Buat minimal 1 team untuk testing players.

### 2. Create Player

**cURL:**
```bash
curl -X POST http://localhost:8080/players \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Budi Santoso",
    "team_id": 1,
    "position": "penyerang",
    "jersey_number": 10,
    "height_cm": 175,
    "weight_kg": 70
  }'
```

**Valid Positions:**
- `penyerang`
- `gelandang`
- `bertahan`
- `penjaga gawang`

**Expected Response (201):**
```json
{
  "id": 1,
  "team_id": 1,
  "name": "Budi Santoso",
  "position": "penyerang",
  "jersey_number": 10,
  "height_cm": 175,
  "weight_kg": 70,
  "created_at": "2025-10-16T12:00:00+07:00",
  "updated_at": "2025-10-16T12:00:00+07:00"
}
```

### 3. Get Players by Team

**cURL:**
```bash
curl -X GET http://localhost:8080/teams/1/players \
  -H "Authorization: Bearer $TOKEN"
```

### 4. Update Player

**cURL:**
```bash
curl -X PUT http://localhost:8080/players/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "height_cm": 180,
    "weight_kg": 75
  }'
```

### 5. Test Jersey Number Validation

Try creating another player with same jersey number:

**cURL:**
```bash
curl -X POST http://localhost:8080/players \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ahmad Dahlan",
    "team_id": 1,
    "position": "gelandang",
    "jersey_number": 10
  }'
```

**Expected Error (400):**
```json
{
  "error": "Nomor punggung sudah digunakan di tim ini"
}
```

---

## 🏆 Testing Matches Endpoints

### 1. Prepare Data
Buat minimal 2 teams dan beberapa players untuk setiap tim.

### 2. Create Match

**cURL:**
```bash
curl -X POST http://localhost:8080/matches \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "home_team_id": 1,
    "away_team_id": 2,
    "match_datetime": "2025-12-20T19:00:00Z"
  }'
```

**Expected Response (201):**
```json
{
  "id": 1,
  "home_team_id": 1,
  "away_team_id": 2,
  "match_datetime": "2025-12-20T19:00:00Z",
  "status": "scheduled",
  "home_score": 0,
  "away_score": 0,
  "created_at": "2025-10-16T12:00:00+07:00",
  "updated_at": "2025-10-16T12:00:00+07:00"
}
```

### 3. Report Match Result

**cURL:**
```bash
curl -X POST http://localhost:8080/matches/1/result \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
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
  }'
```

**⚠️ Important Validations:**
- Total goals MUST equal home_score + away_score
- All player_id must be from teams that are playing
- goal_time must be between 1-120 minutes

**Expected Response (200):**
```json
{
  "message": "Match result reported successfully"
}
```

### 4. Get Match Report

**cURL:**
```bash
curl -X GET http://localhost:8080/matches/1/report \
  -H "Authorization: Bearer $TOKEN"
```

**Expected Response (200):**
```json
{
  "schedule": "2025-12-20T19:00:00+07:00",
  "home_team": "Garuda FC",
  "away_team": "Elang FC",
  "final_score": "2-1",
  "match_result": "Tim Home Menang",
  "top_scorer_in_match": "Budi Santoso (2 gol)",
  "home_team_total_wins": 1,
  "away_team_total_wins": 0
}
```

---

## 🎯 Complete Testing Scenario

### Full Flow Test:

```bash
# 1. Health Check
curl http://localhost:8080/health

# 2. Login
TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}' \
  | jq -r '.token')

# 3. Create Team 1
curl -X POST http://localhost:8080/teams \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Garuda FC","founded_year":2020,"headquarters_city":"Jakarta"}'

# 4. Create Team 2
curl -X POST http://localhost:8080/teams \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Elang FC","founded_year":2019,"headquarters_city":"Bandung"}'

# 5. Create Players for Team 1
curl -X POST http://localhost:8080/players \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Budi Santoso","team_id":1,"position":"penyerang","jersey_number":10}'

# 6. Create Players for Team 2
curl -X POST http://localhost:8080/players \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Eko Prasetyo","team_id":2,"position":"penyerang","jersey_number":9}'

# 7. Create Match
curl -X POST http://localhost:8080/matches \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"home_team_id":1,"away_team_id":2,"match_datetime":"2025-12-20T19:00:00Z"}'

# 8. Report Match Result
curl -X POST http://localhost:8080/matches/1/result \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"home_score":2,"away_score":1,"goals":[{"player_id":1,"goal_time":44},{"player_id":1,"goal_time":78},{"player_id":2,"goal_time":65}]}'

# 9. Get Match Report
curl -X GET http://localhost:8080/matches/1/report \
  -H "Authorization: Bearer $TOKEN"
```

---

## 🔍 Error Testing

### Test Authentication Error

**Missing Token:**
```bash
curl -X GET http://localhost:8080/teams
```

**Expected (401):**
```json
{
  "error": "Token tidak ditemukan"
}
```

**Invalid Token:**
```bash
curl -X GET http://localhost:8080/teams \
  -H "Authorization: Bearer invalid_token"
```

**Expected (401):**
```json
{
  "error": "Token tidak valid atau sudah expired"
}
```

### Test Validation Errors

**Invalid Position:**
```bash
curl -X POST http://localhost:8080/players \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","team_id":1,"position":"striker","jersey_number":11}'
```

**Expected (400):**
```json
{
  "error": "Data player tidak valid: ..."
}
```

---

## 📊 Testing Checklist

### ✅ Authentication
- [ ] Login dengan credentials valid
- [ ] Login dengan credentials invalid
- [ ] Access protected endpoint tanpa token
- [ ] Access protected endpoint dengan token invalid
- [ ] Access protected endpoint dengan token expired

### ✅ Teams
- [ ] Create team
- [ ] Get all teams
- [ ] Get team by ID (exists)
- [ ] Get team by ID (not exists - 404)
- [ ] Update team
- [ ] Delete team
- [ ] Verify soft delete

### ✅ Players
- [ ] Create player
- [ ] Create player dengan jersey number duplikat (should fail)
- [ ] Get players by team
- [ ] Update player
- [ ] Delete player
- [ ] Test posisi validation

### ✅ Matches
- [ ] Create match
- [ ] Create match dengan sama home dan away team (should fail)
- [ ] Report match result
- [ ] Report match result dengan jumlah gol tidak sesuai (should fail)
- [ ] Report match result dengan player bukan dari tim yang main (should fail)
- [ ] Get match report

---

## 🎉 Success Criteria

Testing dianggap berhasil jika:

✅ Semua endpoint return status code yang benar
✅ Response format sesuai spesifikasi
✅ Validasi berjalan dengan baik
✅ Error handling proper
✅ Soft delete berfungsi
✅ JWT authentication bekerja
✅ Database constraints dijaga

---

## 📝 Testing Notes

**Tips:**
- Gunakan Postman untuk save request history
- Test edge cases (empty values, null, etc)
- Verify database state setelah operations
- Test concurrent requests untuk race conditions
- Monitor logs untuk debugging

**Common Issues:**
- Token expired → Login ulang
- Database connection error → Check PostgreSQL service
- Port already in use → Change SERVER_PORT di .env

---

**Happy Testing!** 🧪⚽
