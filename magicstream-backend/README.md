# MagicStream Backend (v3)

Refined Go backend compatible with your Atlas dataset (legacy `genre`/`ranking` fields). Includes JWT auth, rate-limiting, Docker, robust seeder, and tools to migrate+index.

## Run (Atlas)
1. `cp .env.example .env` and set your `MONGO_URI`, `MONGO_DB=magicstream`, `JWT_SECRET`.
2. `go mod tidy`
3. `go run ./cmd/api`
4. (optional) `go run ./internal/seed` to upsert demo data.

## Troubleshooting
- If `go run ./internal/seed` complains directory missing, run `go run internal/seed/seed.go`.
- Create indexes: `go run ./internal/tools/indexes.go`.
- Migrate legacy docs to canonical: `go run ./internal/tools/migrate.go`.
