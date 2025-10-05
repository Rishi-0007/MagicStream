package repository

import (
	"context"
	"time"

	"github.com/rishi-0007/magicstream-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieRepository struct { col *mongo.Collection }
func NewMovieRepository(db *mongo.Database) *MovieRepository { return &MovieRepository{col: db.Collection("movies")} }

func (r *MovieRepository) Create(ctx context.Context, m *models.Movie) error {
	now := time.Now().UTC(); m.CreatedAt, m.UpdatedAt = now, now
	_, err := r.col.InsertOne(ctx, m); return err
}

func projectStage() bson.D {
	return bson.D{{Key: "$project", Value: bson.D{
		{Key: "imdb_id", Value: 1},
		{Key: "title", Value: 1},
		{Key: "poster_path", Value: 1},
		{Key: "release_date", Value: 1},
		{Key: "created_at", Value: 1},
		{Key: "updated_at", Value: 1},
		{Key: "overview", Value: bson.D{{Key: "$ifNull", Value: bson.A{"$overview", "$admin_review"}}}},
		{Key: "genres", Value: bson.D{{Key: "$cond", Value: bson.D{
			{Key: "if", Value: bson.D{{Key: "$ne", Value: bson.A{bson.D{{Key: "$type", Value: "$genres"}}, "missing"}}}},
			{Key: "then", Value: "$genres"},
			{Key: "else", Value: bson.D{{Key: "$map", Value: bson.D{{Key: "input", Value: "$genre"}, {Key: "as", Value: "g"}, {Key: "in", Value: "$$g.genre_name"}}}}},
		}}}},
		{Key: "ranking", Value: bson.D{{Key: "$toDouble", Value: bson.D{{Key: "$ifNull", Value: bson.A{"$ranking.ranking_value", "$ranking.value", "$ranking"}}}}}},
	}}}
}

func (r *MovieRepository) UpsertByIMDBID(ctx context.Context, m *models.Movie) error {
	now := time.Now().UTC(); m.UpdatedAt = now
	up := true
	_, err := r.col.UpdateOne(ctx, bson.M{"imdb_id": m.IMDBID}, bson.M{"$set": m}, &options.UpdateOptions{Upsert: &up})
	return err
}

func (r *MovieRepository) List(ctx context.Context, limit int64) ([]models.Movie, error) {
	cur, err := r.col.Aggregate(ctx, mongo.Pipeline{
		projectStage(),
		bson.D{{Key: "$limit", Value: limit}},
	})
	if err != nil { return nil, err }
	defer cur.Close(ctx)
	var out []models.Movie
	for cur.Next(ctx) {
		var m models.Movie
		if err := cur.Decode(&m); err != nil { return nil, err }
		out = append(out, m)
	}
	return out, cur.Err()
}

func (r *MovieRepository) ByIMDBID(ctx context.Context, id string) (*models.Movie, error) {
	cur, err := r.col.Aggregate(ctx, mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: "imdb_id", Value: id}}}},
		projectStage(),
		bson.D{{Key: "$limit", Value: 1}},
	})
	if err != nil { return nil, err }
	defer cur.Close(ctx)
	if cur.Next(ctx) {
		var m models.Movie
		if err := cur.Decode(&m); err != nil { return nil, err }
		return &m, nil
	}
	return nil, mongo.ErrNoDocuments
}

func (r *MovieRepository) SearchByTitle(ctx context.Context, q string, limit int64) ([]models.Movie, error) {
	cur, err := r.col.Aggregate(ctx, mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: q}, {Key: "$options", Value: "i"}}}}}},
		projectStage(),
		bson.D{{Key: "$limit", Value: limit}},
	})
	if err != nil { return nil, err }
	defer cur.Close(ctx)
	var out []models.Movie
	for cur.Next(ctx) {
		var m models.Movie
		if err := cur.Decode(&m); err != nil { return nil, err }
		out = append(out, m)
	}
	return out, cur.Err()
}
