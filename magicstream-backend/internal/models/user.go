package models

import "time"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	Email      string    `bson:"email" json:"email" validate:"required,email"`
	Password   string    `bson:"password" json:"-"`
	Name       string    `bson:"name" json:"name" validate:"required,min=2,max=50"`
	AvatarURL  string    `bson:"avatar_url" json:"avatar_url"`
	Provider   string    `bson:"provider" json:"provider"`
	ProviderID string    `bson:"provider_id" json:"provider_id"`
	Role       Role      `bson:"role" json:"role"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}
