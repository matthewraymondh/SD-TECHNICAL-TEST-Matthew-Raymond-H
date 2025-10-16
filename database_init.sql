-- Inisialisasi Database untuk XYZ Football API
-- Jalankan script ini di PostgreSQL untuk membuat database dan tabel

-- 1. Buat database (jalankan di postgres database default)
-- CREATE DATABASE football_db;

-- 2. Connect ke database football_db
-- \c football_db

-- 3. Buat tabel teams
CREATE TABLE IF NOT EXISTS teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    logo_url VARCHAR(255),
    founded_year INT,
    headquarters_address TEXT,
    headquarters_city VARCHAR(100),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

-- 4. Buat tabel players
CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    team_id INT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    height_cm INT,
    weight_kg INT,
    position VARCHAR(50) NOT NULL CHECK (position IN ('penyerang', 'gelandang', 'bertahan', 'penjaga gawang')),
    jersey_number INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,
    UNIQUE (team_id, jersey_number, deleted_at)
);

-- 5. Buat tabel matches
CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    home_team_id INT NOT NULL REFERENCES teams(id),
    away_team_id INT NOT NULL REFERENCES teams(id),
    match_datetime TIMESTAMPTZ NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'scheduled' CHECK (status IN ('scheduled', 'completed', 'cancelled')),
    home_score INT DEFAULT 0,
    away_score INT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

-- 6. Buat tabel goals
CREATE TABLE IF NOT EXISTS goals (
    id SERIAL PRIMARY KEY,
    match_id INT NOT NULL REFERENCES matches(id) ON DELETE CASCADE,
    player_id INT NOT NULL REFERENCES players(id),
    goal_time INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- 7. Buat indexes untuk performa
CREATE INDEX IF NOT EXISTS idx_teams_deleted_at ON teams(deleted_at);
CREATE INDEX IF NOT EXISTS idx_players_deleted_at ON players(deleted_at);
CREATE INDEX IF NOT EXISTS idx_players_team_id ON players(team_id);
CREATE INDEX IF NOT EXISTS idx_matches_deleted_at ON matches(deleted_at);
CREATE INDEX IF NOT EXISTS idx_matches_status ON matches(status);
CREATE INDEX IF NOT EXISTS idx_goals_match_id ON goals(match_id);
CREATE INDEX IF NOT EXISTS idx_goals_player_id ON goals(player_id);

-- 8. Insert sample data (optional)
-- Teams
INSERT INTO teams (name, founded_year, headquarters_city, headquarters_address) VALUES
('Garuda FC', 2020, 'Jakarta', 'Jl. Sudirman No. 1'),
('Elang FC', 2019, 'Bandung', 'Jl. Asia Afrika No. 25'),
('Naga FC', 2021, 'Surabaya', 'Jl. Pemuda No. 10')
ON CONFLICT DO NOTHING;

-- Players for Garuda FC (team_id = 1)
INSERT INTO players (team_id, name, position, jersey_number, height_cm, weight_kg) VALUES
(1, 'Budi Santoso', 'penyerang', 10, 175, 70),
(1, 'Ahmad Dahlan', 'gelandang', 8, 170, 68),
(1, 'Rudi Hartono', 'bertahan', 5, 180, 75),
(1, 'Agus Wijaya', 'penjaga gawang', 1, 185, 80)
ON CONFLICT DO NOTHING;

-- Players for Elang FC (team_id = 2)
INSERT INTO players (team_id, name, position, jersey_number, height_cm, weight_kg) VALUES
(2, 'Eko Prasetyo', 'penyerang', 9, 178, 72),
(2, 'Doni Kusuma', 'gelandang', 7, 172, 69),
(2, 'Cahya Ramadan', 'bertahan', 4, 182, 77),
(2, 'Bambang Surya', 'penjaga gawang', 1, 188, 82)
ON CONFLICT DO NOTHING;

-- Sample Match
INSERT INTO matches (home_team_id, away_team_id, match_datetime, status) VALUES
(1, 2, NOW() + INTERVAL '7 days', 'scheduled')
ON CONFLICT DO NOTHING;

COMMIT;

-- Verifikasi data
SELECT 'Teams Count:' as info, COUNT(*) as total FROM teams;
SELECT 'Players Count:' as info, COUNT(*) as total FROM players;
SELECT 'Matches Count:' as info, COUNT(*) as total FROM matches;
