package models

import "time"

type Movie struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	IMDBID      string    `bson:"imdb_id" json:"imdb_id" validate:"required"`
	Title       string    `bson:"title" json:"title" validate:"required"`
	Overview    string    `bson:"overview" json:"overview"`
	PosterPath  string    `bson:"poster_path" json:"poster_path"`
	ReleaseDate string    `bson:"release_date" json:"release_date"`
	Genres      []string  `bson:"genres" json:"genres"`
	Ranking     float64   `bson:"ranking" json:"ranking"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
