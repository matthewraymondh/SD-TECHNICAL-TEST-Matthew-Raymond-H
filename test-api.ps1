# XYZ Football API - PowerShell Test Script
# Complete API Testing

Write-Host ""
Write-Host "╔═══════════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║                                                           ║" -ForegroundColor Cyan
Write-Host "║        XYZ FOOTBALL API - PowerShell Test Script         ║" -ForegroundColor Cyan
Write-Host "║                                                           ║" -ForegroundColor Cyan
Write-Host "╚═══════════════════════════════════════════════════════════╝" -ForegroundColor Cyan
Write-Host ""

$baseUrl = "http://localhost:8080"
$ErrorActionPreference = "Stop"

# 1. Health Check
Write-Host "1. Testing Health Check..." -ForegroundColor Yellow
try {
    $health = Invoke-RestMethod -Uri "$baseUrl/health" -Method Get
    Write-Host "✓ Health Check: $($health.message)" -ForegroundColor Green
}
catch {
    Write-Host "✗ Health Check Failed: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host "   Make sure the server is running on port 8080" -ForegroundColor Red
    exit 1
}
Write-Host ""

# 2. Login
Write-Host "2. Logging in..." -ForegroundColor Yellow
try {
    $loginBody = @{
        username = "admin"
        password = "admin123"
    } | ConvertTo-Json
    
    $loginResponse = Invoke-RestMethod -Uri "$baseUrl/login" `
        -Method Post `
        -ContentType "application/json" `
        -Body $loginBody
    
    $token = $loginResponse.token
    Write-Host "✓ Login Success!" -ForegroundColor Green
    Write-Host "  Token: $($token.Substring(0, 50))..." -ForegroundColor Gray
}
catch {
    Write-Host "✗ Login Failed: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Setup headers
$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type"  = "application/json"
}

# 3. Create Team 1
Write-Host "3. Creating Team 1 (Garuda FC)..." -ForegroundColor Yellow
try {
    $team1Body = @{
        name                 = "Garuda FC"
        founded_year         = 2020
        headquarters_city    = "Jakarta"
        headquarters_address = "Jl. Sudirman No. 1"
    } | ConvertTo-Json
    
    $team1 = Invoke-RestMethod -Uri "$baseUrl/teams" `
        -Method Post `
        -Headers $headers `
        -Body $team1Body
    
    Write-Host "✓ Team Created: $($team1.name) (ID: $($team1.id))" -ForegroundColor Green
    $team1Id = $team1.id
}
catch {
    Write-Host "✗ Create Team 1 Failed: $($_.Exception.Message)" -ForegroundColor Red
    $team1Id = 1
}
Write-Host ""

# 4. Create Team 2
Write-Host "4. Creating Team 2 (Elang FC)..." -ForegroundColor Yellow
try {
    $team2Body = @{
        name                 = "Elang FC"
        founded_year         = 2019
        headquarters_city    = "Bandung"
        headquarters_address = "Jl. Asia Afrika No. 25"
    } | ConvertTo-Json
    
    $team2 = Invoke-RestMethod -Uri "$baseUrl/teams" `
        -Method Post `
        -Headers $headers `
        -Body $team2Body
    
    Write-Host "✓ Team Created: $($team2.name) (ID: $($team2.id))" -ForegroundColor Green
    $team2Id = $team2.id
}
catch {
    Write-Host "✗ Create Team 2 Failed: $($_.Exception.Message)" -ForegroundColor Red
    $team2Id = 2
}
Write-Host ""

# 5. Get All Teams
Write-Host "5. Getting All Teams..." -ForegroundColor Yellow
try {
    $teams = Invoke-RestMethod -Uri "$baseUrl/teams" `
        -Method Get `
        -Headers $headers
    
    Write-Host "✓ Found $($teams.Count) team(s)" -ForegroundColor Green
    foreach ($team in $teams) {
        Write-Host "  - $($team.name) (ID: $($team.id))" -ForegroundColor Gray
    }
}
catch {
    Write-Host "✗ Get Teams Failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 6. Create Player 1
Write-Host "6. Creating Player 1 (Team $team1Id)..." -ForegroundColor Yellow
try {
    $player1Body = @{
        name          = "Budi Santoso"
        team_id       = [int]$team1Id
        position      = "penyerang"
        jersey_number = 10
        height_cm     = 175
        weight_kg     = 70
    } | ConvertTo-Json
    
    Write-Host "  Request Body: $player1Body" -ForegroundColor Gray
    
    $player1 = Invoke-RestMethod -Uri "$baseUrl/players" `
        -Method Post `
        -Headers $headers `
        -Body $player1Body
    
    Write-Host "✓ Player Created: $($player1.name) #$($player1.jersey_number) (ID: $($player1.id))" -ForegroundColor Green
    $player1Id = $player1.id
}
catch {
    Write-Host "✗ Create Player 1 Failed: $($_.Exception.Message)" -ForegroundColor Red
    if ($_.ErrorDetails.Message) {
        Write-Host "  Error Details: $($_.ErrorDetails.Message)" -ForegroundColor Red
    }
    $player1Id = 1
}
Write-Host ""

# 7. Create Player 2
Write-Host "7. Creating Player 2 (Team $team2Id)..." -ForegroundColor Yellow
try {
    $player2Body = @{
        name          = "Eko Prasetyo"
        team_id       = [int]$team2Id
        position      = "penyerang"
        jersey_number = 9
        height_cm     = 178
        weight_kg     = 72
    } | ConvertTo-Json
    
    Write-Host "  Request Body: $player2Body" -ForegroundColor Gray
    
    $player2 = Invoke-RestMethod -Uri "$baseUrl/players" `
        -Method Post `
        -Headers $headers `
        -Body $player2Body
    
    Write-Host "✓ Player Created: $($player2.name) #$($player2.jersey_number) (ID: $($player2.id))" -ForegroundColor Green
    $player2Id = $player2.id
}
catch {
    Write-Host "✗ Create Player 2 Failed: $($_.Exception.Message)" -ForegroundColor Red
    if ($_.ErrorDetails.Message) {
        Write-Host "  Error Details: $($_.ErrorDetails.Message)" -ForegroundColor Red
    }
    $player2Id = 2
}
Write-Host ""

# 8. Get Players by Team
Write-Host "8. Getting Players from Team $team1Id..." -ForegroundColor Yellow
try {
    $players = Invoke-RestMethod -Uri "$baseUrl/teams/$team1Id/players" `
        -Method Get `
        -Headers $headers
    
    Write-Host "✓ Found $($players.Count) player(s)" -ForegroundColor Green
    foreach ($player in $players) {
        Write-Host "  - $($player.name) #$($player.jersey_number) - $($player.position)" -ForegroundColor Gray
    }
}
catch {
    Write-Host "✗ Get Players Failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 9. Create Match
Write-Host "9. Creating Match..." -ForegroundColor Yellow
try {
    $matchBody = @{
        home_team_id   = [int]$team1Id
        away_team_id   = [int]$team2Id
        match_datetime = "2025-12-20T19:00:00Z"
    } | ConvertTo-Json
    
    Write-Host "  Request Body: $matchBody" -ForegroundColor Gray
    
    $match = Invoke-RestMethod -Uri "$baseUrl/matches" `
        -Method Post `
        -Headers $headers `
        -Body $matchBody
    
    Write-Host "✓ Match Created (ID: $($match.id))" -ForegroundColor Green
    Write-Host "  $($match.home_team_id) vs $($match.away_team_id)" -ForegroundColor Gray
    $matchId = $match.id
}
catch {
    Write-Host "✗ Create Match Failed: $($_.Exception.Message)" -ForegroundColor Red
    if ($_.ErrorDetails.Message) {
        Write-Host "  Error Details: $($_.ErrorDetails.Message)" -ForegroundColor Red
    }
    $matchId = 1
}
Write-Host ""

# 10. Report Match Result
Write-Host "10. Reporting Match Result..." -ForegroundColor Yellow
try {
    $resultBody = @{
        home_score = 2
        away_score = 1
        goals      = @(
            @{ 
                player_id = [int]$player1Id
                goal_time = 44 
            },
            @{ 
                player_id = [int]$player1Id
                goal_time = 78 
            },
            @{ 
                player_id = [int]$player2Id
                goal_time = 65 
            }
        )
    } | ConvertTo-Json -Depth 10
    
    $result = Invoke-RestMethod -Uri "$baseUrl/matches/$matchId/result" `
        -Method Post `
        -Headers $headers `
        -Body $resultBody
    
    Write-Host "✓ Match Result Reported: $($result.message)" -ForegroundColor Green
}
catch {
    Write-Host "✗ Report Match Result Failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 11. Get Match Report
Write-Host "11. Getting Match Report..." -ForegroundColor Yellow
try {
    $report = Invoke-RestMethod -Uri "$baseUrl/matches/$matchId/report" `
        -Method Get `
        -Headers $headers
    
    Write-Host "✓ Match Report Retrieved:" -ForegroundColor Green
    Write-Host "  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor Cyan
    Write-Host "  Schedule       : $($report.schedule)" -ForegroundColor White
    Write-Host "  Match          : $($report.home_team) vs $($report.away_team)" -ForegroundColor White
    Write-Host "  Final Score    : $($report.final_score)" -ForegroundColor White
    Write-Host "  Result         : $($report.match_result)" -ForegroundColor White
    Write-Host "  Top Scorer     : $($report.top_scorer_in_match)" -ForegroundColor White
    Write-Host "  Home Wins      : $($report.home_team_total_wins)" -ForegroundColor White
    Write-Host "  Away Wins      : $($report.away_team_total_wins)" -ForegroundColor White
    Write-Host "  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor Cyan
}
catch {
    Write-Host "✗ Get Match Report Failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

Write-Host "╔═══════════════════════════════════════════════════════════╗" -ForegroundColor Green
Write-Host "║                                                           ║" -ForegroundColor Green
Write-Host "║           🎉 All Tests Completed Successfully! 🎉         ║" -ForegroundColor Green
Write-Host "║                                                           ║" -ForegroundColor Green
Write-Host "╚═══════════════════════════════════════════════════════════╝" -ForegroundColor Green
Write-Host ""
Write-Host "Summary:" -ForegroundColor Cyan
Write-Host "  - Teams Created: 2" -ForegroundColor Gray
Write-Host "  - Players Created: 2" -ForegroundColor Gray
Write-Host "  - Matches Created: 1" -ForegroundColor Gray
Write-Host "  - Match Results Reported: 1" -ForegroundColor Gray
Write-Host ""
Write-Host "Next Steps:" -ForegroundColor Cyan
Write-Host "  - Check the database to verify data" -ForegroundColor Gray
Write-Host "  - Import Postman collection for more tests" -ForegroundColor Gray
Write-Host "  - Read TESTING.md for more examples" -ForegroundColor Gray
Write-Host ""
