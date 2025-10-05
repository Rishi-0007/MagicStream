package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rishi-0007/magicstream-backend/internal/config"
	"github.com/rishi-0007/magicstream-backend/internal/database"
	"github.com/rishi-0007/magicstream-backend/internal/models"
	"github.com/rishi-0007/magicstream-backend/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()
	mongoConn, err := database.Connect(ctx, cfg.MongoURI, cfg.MongoDB)
	if err != nil { log.Fatal(err) }
	defer mongoConn.Close(ctx)
	log.Println("connected to MongoDB")
	db := mongoConn.DB
	if err := seedGenres(ctx, db); err != nil { log.Fatal(err) }
	if err := seedMovies(ctx, db); err != nil { log.Fatal(err) }
	if err := seedUsers(ctx, db); err != nil { log.Fatal(err) }
	fmt.Println("seeding complete")
}

type rawGenre struct{ GenreID int `json:"genre_id"`; GenreName string `json:"genre_name"` }
type rawMovie struct {
	IMDBID any `json:"imdb_id"`; Title any `json:"title"`; PosterPath any `json:"poster_path"`
	ReleaseDate any `json:"release_date"`; AdminReview any `json:"admin_review"`
	Genre []rawGenre `json:"genre"`; Ranking any `json:"ranking"`
}
type rawUserDate struct{ Date string `json:"$date"` }
type rawUser struct {
	UserID string `json:"user_id"`; FirstName string `json:"first_name"`; LastName string `json:"last_name"`
	Email string `json:"email"`; Password string `json:"password"`; Role string `json:"role"`
	CreatedAtWire rawUserDate `json:"created_at"`; UpdatedAtWire rawUserDate `json:"updated_at"`
	FavouriteGenres []rawGenre `json:"favourite_genres"`
}

func seedGenres(ctx context.Context, db *mongo.Database) error {
	path := "api/genres.json"; f, err := os.Open(path)
	if err != nil { log.Printf("genres: %s not found, skipping\n", path); return nil }
	defer f.Close()
	var genres []rawGenre
	if err := json.NewDecoder(f).Decode(&genres); err != nil { return err }
	col := db.Collection("genres")
	for _, g := range genres {
		_, err := col.UpdateOne(ctx, bson.M{"genre_id": g.GenreID}, bson.M{"$set": g}, upsertTrue())
		if err != nil { return err }
	}
	log.Printf("seeded %d genres", len(genres)); return nil
}

func seedMovies(ctx context.Context, db *mongo.Database) error {
	f, err := os.Open("api/movies.json"); if err != nil { return err }
	defer f.Close()
	var raw []rawMovie
	if err := json.NewDecoder(f).Decode(&raw); err != nil { return err }
	col := db.Collection("movies"); now := time.Now().UTC(); count := 0
	for _, r := range raw {
		imdb, title := toString(r.IMDBID), toString(r.Title); if imdb=="" || title=="" { continue }
		m := models.Movie{ IMDBID: imdb, Title: title, PosterPath: toString(r.PosterPath),
			ReleaseDate: toString(r.ReleaseDate), Overview: toString(r.AdminReview),
			Genres: genreNames(r.Genre), Ranking: extractRanking(r.Ranking),
			CreatedAt: now, UpdatedAt: now }
		_, err := col.UpdateOne(ctx, bson.M{"imdb_id": m.IMDBID}, bson.M{"$set": m}, upsertTrue())
		if err != nil { return err }; count++
	}
	log.Printf("upserted %d movies", count); return nil
}

func seedUsers(ctx context.Context, db *mongo.Database) error {
	f, err := os.Open("api/users.json"); if err != nil { return err }
	defer f.Close(); var raw []rawUser
	if err := json.NewDecoder(f).Decode(&raw); err != nil { return err }
	col := db.Collection("users"); count := 0
	for _, u := range raw {
		name := strings.TrimSpace(strings.Join([]string{u.FirstName, u.LastName}, " "))
		role := strings.ToLower(u.Role); if role != string(models.RoleAdmin) { role = string(models.RoleUser) }
		salt := utils.RandomSalt(); hashed := utils.HashPassword("secret123", salt) + ":" + salt
		createdAt := parseMaybeISO(u.CreatedAtWire.Date); updatedAt := parseMaybeISO(u.UpdatedAtWire.Date)
		if createdAt.IsZero() { createdAt = time.Now().UTC() }; if updatedAt.IsZero() { updatedAt = createdAt }
		doc := models.User{ Email: u.Email, Password: hashed, Name: name, Role: models.Role(role), CreatedAt: createdAt, UpdatedAt: updatedAt }
		_, err := col.UpdateOne(ctx, bson.M{"email": doc.Email}, bson.M{"$setOnInsert": doc}, upsertTrue()); if err != nil { return err }
		count++
	}
	log.Printf("upserted %d users (password set to "secret123")", count); return nil
}

func genreNames(gs []rawGenre) []string { out := make([]string,0,len(gs)); for _, g := range gs { if s:=strings.TrimSpace(g.GenreName); s!="" { out = append(out, s) } }; return out }
func extractRanking(v any) float64 {
	switch t := v.(type) {
	case float64: return t
	case string: if f, err := strconv.ParseFloat(t, 64); err == nil { return f }
	case map[string]any:
		if val, ok := t["ranking_value"]; ok { return toFloat(val) }
		if val, ok := t["value"]; ok { return toFloat(val) }
	}
	return 0
}
func toFloat(v any) float64 { switch x := v.(type) { case float64: return x; case string: if f,err:=strconv.ParseFloat(x,64); err==nil { return f } }; return 0 }
func toString(v any) string { switch x := v.(type) { case string: return x; case float64: return strconv.FormatFloat(x,'f',-1,64); case json.Number: return x.String(); default: return "" } }
func parseMaybeISO(s string) time.Time { if s=="" { return time.Time{} }; t,err:=time.Parse(time.RFC3339,s); if err!=nil { return time.Time{} }; return t.UTC() }
func upsertTrue() *mongo.UpdateOptions { up := true; return &mongo.UpdateOptions{Upsert: &up} }
