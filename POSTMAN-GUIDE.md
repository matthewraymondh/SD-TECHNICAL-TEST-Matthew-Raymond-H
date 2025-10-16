# 📬 Postman Testing Guide - XYZ Football API

## 📥 Step 1: Import Collection ke Postman

### Cara 1: Import File (Recommended)

1. **Buka Postman Application**
   - Download Postman dari: https://www.postman.com/downloads/
   - Atau gunakan Postman Web: https://web.postman.com

2. **Import Collection**
   - Klik tombol **"Import"** di kiri atas
   - Pilih **"Upload Files"** atau drag & drop file
   - Browse ke file: `postman_collection.json`
   - Klik **"Import"**

3. **Verifikasi Import**
   - Collection **"XYZ Football API"** akan muncul di sidebar kiri
   - Expand collection untuk melihat 4 folders:
     - Authentication
     - Teams (5 requests)
     - Players (4 requests)
     - Matches (3 requests)
     - Health Check

### Cara 2: Copy-Paste JSON

1. Buka file `postman_collection.json`
2. Copy seluruh isi file (Ctrl+A, Ctrl+C)
3. Di Postman, klik **"Import"** → **"Raw text"**
4. Paste JSON → Klik **"Import"**

---

## ⚙️ Step 2: Setup Environment (Optional tapi Recommended)

### Mengapa Perlu Environment?
- Memudahkan switch antara Development/Production
- Token tersimpan otomatis setelah login
- Base URL bisa diubah dengan mudah

### Cara Setup Environment:

1. **Buka Environment Settings**
   - Klik icon ⚙️ (Environments) di kiri atas
   - Atau klik dropdown "No Environment" di kanan atas

2. **Create New Environment**
   - Klik **"Create Environment"** atau tombol **"+"**
   - Nama: `XYZ Football API - Local`

3. **Add Variables**
   
   | Variable Name | Type    | Initial Value          | Current Value          |
   |--------------|---------|------------------------|------------------------|
   | `base_url`   | default | `http://localhost:8080`| `http://localhost:8080`|
   | `token`      | secret  | (kosongkan)            | (kosongkan)            |

4. **Save Environment**
   - Klik **"Save"** atau Ctrl+S
   - Pilih environment dari dropdown di kanan atas

> **Note**: Token akan otomatis terisi setelah login berhasil!

---

## 🚀 Step 3: Testing API (Flow Lengkap)

### Prerequisites
✅ API server sudah running: `go run cmd/api/main.go`  
✅ Database PostgreSQL sudah running  
✅ Collection sudah di-import  
✅ Environment sudah aktif (pilih dari dropdown)

---

### 🔐 Test 1: Login & Get Token

1. **Expand folder "Authentication"**
2. **Click "Login"**
3. **Lihat Request**:
   - Method: `POST`
   - URL: `{{base_url}}/login`
   - Body (JSON):
     ```json
     {
         "username": "admin",
         "password": "admin123"
     }
     ```

4. **Click "Send"**

5. **Expected Response** (Status: 200 OK):
   ```json
   {
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
   }
   ```

6. **Verify Token Saved**:
   - Klik icon 👁️ (Environment quick look) di kanan atas
   - Lihat variable `token` sudah terisi
   - Token ini akan digunakan untuk semua request berikutnya

> **✨ Magic!** Token otomatis disimpan ke environment variable via script:
> ```javascript
> pm.collectionVariables.set("token", jsonData.token);
> ```

---

### 🏆 Test 2: Create Teams

1. **Expand folder "Teams"**
2. **Click "Create Team"**
3. **Lihat Headers**:
   - `Authorization: Bearer {{token}}` (otomatis menggunakan token dari login)
   - `Content-Type: application/json`

4. **Edit Body** (sesuaikan data):
   ```json
   {
       "name": "Garuda FC",
       "founded_year": 2020,
       "headquarters_city": "Jakarta",
       "headquarters_address": "Jl. Sudirman No. 1",
       "logo_url": "https://example.com/garuda-logo.png"
   }
   ```

5. **Click "Send"**

6. **Expected Response** (Status: 201 Created):
   ```json
   {
       "id": 1,
       "name": "Garuda FC",
       "founded_year": 2020,
       "headquarters_city": "Jakarta",
       "headquarters_address": "Jl. Sudirman No. 1",
       "logo_url": "https://example.com/garuda-logo.png",
       "created_at": "2025-10-16T13:30:00Z",
       "updated_at": "2025-10-16T13:30:00Z"
   }
   ```

7. **Save Team ID**: Catat `id` yang dikembalikan (misal: 1)

8. **Create Team Kedua**: Ulangi dengan data berbeda (misal: "Elang FC")

---

### ⚽ Test 3: Create Players

1. **Click "Create Player"**
2. **Edit Body** (ganti `team_id` dengan ID team yang sudah dibuat):
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

3. **Valid Positions**:
   - `penyerang` (forward)
   - `gelandang` (midfielder)
   - `bertahan` (defender)
   - `penjaga gawang` (goalkeeper)

4. **Click "Send"**

5. **Expected Response** (Status: 201 Created):
   ```json
   {
       "id": 1,
       "name": "Budi Santoso",
       "team_id": 1,
       "position": "penyerang",
       "jersey_number": 10,
       "height_cm": 175,
       "weight_kg": 70,
       "created_at": "2025-10-16T13:35:00Z",
       "updated_at": "2025-10-16T13:35:00Z"
   }
   ```

6. **Create More Players**: Minimal 2 players untuk test match nanti

---

### 📋 Test 4: Get All Teams

1. **Click "Get All Teams"**
2. **Click "Send"**
3. **Expected Response** (Status: 200 OK):
   ```json
   [
       {
           "id": 1,
           "name": "Garuda FC",
           ...
       },
       {
           "id": 2,
           "name": "Elang FC",
           ...
       }
   ]
   ```

---

### 🏟️ Test 5: Create Match

1. **Click "Create Match"**
2. **Edit Body** (gunakan team_id yang valid):
   ```json
   {
       "home_team_id": 1,
       "away_team_id": 2,
       "match_datetime": "2025-12-20T19:00:00Z"
   }
   ```

3. **Format Datetime**: ISO 8601 (`YYYY-MM-DDTHH:MM:SSZ`)

4. **Click "Send"**

5. **Expected Response** (Status: 201 Created):
   ```json
   {
       "id": 1,
       "home_team_id": 1,
       "away_team_id": 2,
       "match_datetime": "2025-12-20T19:00:00Z",
       "status": "scheduled",
       ...
   }
   ```

6. **Save Match ID**: Catat `id` untuk report result

---

### 📊 Test 6: Report Match Result

1. **Click "Report Match Result"**
2. **Edit URL**: Ganti `:id` dengan match ID (misal: `/matches/1/result`)
3. **Edit Body**:
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
               "player_id": 2,
               "goal_time": 65
           }
       ]
   }
   ```

4. **Validation**:
   - Total goals dalam array HARUS = `home_score + away_score`
   - `goal_time` antara 1-120 menit
   - `player_id` harus valid

5. **Click "Send"**

6. **Expected Response** (Status: 200 OK):
   ```json
   {
       "message": "Match result reported successfully"
   }
   ```

---

### 📈 Test 7: Get Match Report

1. **Click "Get Match Report"**
2. **Edit URL**: `/matches/1/report`
3. **Click "Send"**

4. **Expected Response** (Status: 200 OK):
   ```json
   {
       "schedule": "12/21/2025 02:00:00",
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

## 🔄 Test Other Endpoints

### Update Team
- Method: `PUT`
- URL: `/teams/1`
- Body: Hanya field yang mau diupdate
  ```json
  {
      "name": "Garuda Perkasa FC",
      "headquarters_address": "Jl. Thamrin No. 5"
  }
  ```

### Delete Team (Soft Delete)
- Method: `DELETE`
- URL: `/teams/1`
- Response: `204 No Content`

### Get Players by Team
- Method: `GET`
- URL: `/teams/1/players`
- Response: Array of players

### Update Player
- Method: `PUT`
- URL: `/players/1`
- Body:
  ```json
  {
      "height_cm": 180,
      "weight_kg": 75
  }
  ```

---

## 💡 Pro Tips

### 1. Use Collection Variables
Collection sudah punya 2 variables:
- `{{base_url}}` - Base URL API
- `{{token}}` - JWT Token (auto-filled after login)

### 2. Check Pre-request Scripts
Login request punya script untuk auto-save token:
```javascript
if (pm.response.code === 200) {
    var jsonData = pm.response.json();
    pm.collectionVariables.set("token", jsonData.token);
}
```

### 3. Bearer Token Authentication
Semua protected endpoints otomatis menggunakan:
```
Authorization: Bearer {{token}}
```

### 4. Save Responses
Klik **"Save Response"** untuk menyimpan contoh response yang bagus.

### 5. Use Tests Tab
Tambahkan assertions untuk automated testing:
```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Response has token", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('token');
});
```

### 6. Organize Requests
Requests sudah dikelompokkan dalam folders:
- **Authentication** - Login
- **Teams** - CRUD Teams
- **Players** - CRUD Players
- **Matches** - Match operations
- **Health Check** - API status

### 7. Use Request Examples
Setiap request sudah punya contoh body yang valid, tinggal edit sesuai kebutuhan.

---

## 🐛 Troubleshooting

### Error: "Could not send request"
❌ **Problem**: API server tidak running  
✅ **Solution**: Jalankan `go run cmd/api/main.go`

### Error: 401 Unauthorized
❌ **Problem**: Token expired atau tidak ada  
✅ **Solution**: Login ulang untuk mendapatkan token baru

### Error: 400 Bad Request
❌ **Problem**: Request body tidak valid  
✅ **Solution**: 
- Cek format JSON
- Pastikan semua required fields ada
- Cek tipe data (int vs string)

### Error: 404 Not Found
❌ **Problem**: Resource ID tidak ditemukan  
✅ **Solution**: Pastikan ID yang digunakan valid dan resource belum di-delete

### Token tidak auto-save
❌ **Problem**: Environment tidak aktif  
✅ **Solution**: 
- Pilih environment dari dropdown kanan atas
- Atau gunakan Collection Variables

---

## 📚 Complete Testing Flow

```
1. Health Check          → Verify API is running
2. Login                 → Get JWT token (auto-saved)
3. Create Team 1         → Garuda FC
4. Create Team 2         → Elang FC
5. Get All Teams         → Verify 2 teams created
6. Create Player 1       → Budi (Team 1)
7. Create Player 2       → Eko (Team 2)
8. Get Players by Team   → Verify players
9. Create Match          → Team 1 vs Team 2
10. Report Match Result  → Score 2-1, with goals
11. Get Match Report     → Complete statistics
12. Update Team          → Change team info
13. Update Player        → Change player stats
14. Delete Player        → Soft delete
15. Delete Team          → Soft delete
```

---

## 🎯 Testing Checklist

- [ ] Import collection berhasil
- [ ] Environment setup (optional)
- [ ] Health check - API running
- [ ] Login - Token tersimpan
- [ ] Create 2 teams
- [ ] Get all teams
- [ ] Create 2+ players
- [ ] Get players by team
- [ ] Create match
- [ ] Report match result
- [ ] Get match report
- [ ] Update operations
- [ ] Delete operations (soft delete)

---

## 🚀 Next Steps

Setelah familiar dengan Postman:
1. ✅ Explore semua endpoints
2. ✅ Test edge cases (invalid data, missing fields, dll)
3. ✅ Create automated test scripts
4. ✅ Export environment untuk sharing dengan team
5. ✅ Setup Collection Runner untuk regression testing

---

**Happy Testing! 🎉**

Need help? Check:
- `README.md` - Complete project documentation
- `TESTING.md` - General testing guide
- `POWERSHELL-TESTING.md` - PowerShell specific commands
- `QUICKSTART.md` - Quick setup guide
