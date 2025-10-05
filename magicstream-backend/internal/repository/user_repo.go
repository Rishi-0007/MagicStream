package repository

import (
	"context"
	"time"

	"github.com/rishi-0007/magicstream-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct { col *mongo.Collection }
func NewUserRepository(db *mongo.Database) *UserRepository { return &UserRepository{col: db.Collection("users")} }

func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	u.CreatedAt = time.Now().UTC(); u.UpdatedAt = u.CreatedAt
	_, err := r.col.InsertOne(ctx, u); return err
}
func (r *UserRepository) ByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	if err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&u); err != nil { return nil, err }
	return &u, nil
}
func (r *UserRepository) ByID(ctx context.Context, id string) (*models.User, error) {
	var u models.User
	obj, err := primitive.ObjectIDFromHex(id); if err != nil { return nil, err }
	if err := r.col.FindOne(ctx, bson.M{"_id": obj}).Decode(&u); err != nil { return nil, err }
	return &u, nil
}
