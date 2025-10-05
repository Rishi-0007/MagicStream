package config

import (
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Port            string
	MongoURI        string
	MongoDB         string
	JWTSecret       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	CORSOrigins     []string
	Env             string
}

func mustEnv(key, def string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		if def != "" { return def }
		log.Fatalf("missing required env var: %s", key)
	}
	return v
}

func Load() *Config {
	if _, err := os.Stat(".env"); err == nil { loadDotEnv(".env") }
	return &Config{
		Port:            mustEnv("PORT", "8080"),
		MongoURI:        mustEnv("MONGO_URI", ""),
		MongoDB:         mustEnv("MONGO_DB", "magicstream"),
		JWTSecret:       mustEnv("JWT_SECRET", ""),
		AccessTokenTTL:  parseDuration(mustEnv("ACCESS_TOKEN_TTL", "15m")),
		RefreshTokenTTL: parseDuration(mustEnv("REFRESH_TOKEN_TTL", "168h")),
		CORSOrigins:     parseCSV(mustEnv("CORS_ORIGINS", "*")),
		Env:             mustEnv("ENV", "development"),
	}
}

func parseCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" { out = append(out, p) }
	}
	return out
}

func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil { log.Fatalf("invalid duration %q: %v", s, err) }
	return d
}

func loadDotEnv(filename string) {
	b, err := os.ReadFile(filename)
	if err != nil { return }
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") { continue }
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 { continue }
		_ = os.Setenv(strings.TrimSpace(kv[0]), strings.Trim(strings.TrimSpace(kv[1]), `"'`))
	}
}
