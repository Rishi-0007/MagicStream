package services

import (
	"context"
	"errors"
	"time"
	"github.com/rishi-0007/magicstream-backend/internal/models"
	"github.com/rishi-0007/magicstream-backend/internal/repository"
	"github.com/rishi-0007/magicstream-backend/internal/utils"
)

type AuthService struct { users *repository.UserRepository; jwtSecret string; accessTTL, refreshTTL time.Duration }
func NewAuthService(users *repository.UserRepository, jwtSecret string, accessTTL, refreshTTL time.Duration)*AuthService{ return &AuthService{users, jwtSecret, accessTTL, refreshTTL} }

func (s *AuthService) Register(ctx context.Context, email, name, password string) (*models.User, error) {
	if _, err := s.users.ByEmail(ctx, email); err == nil { return nil, errors.New("email already in use") }
	u := &models.User{ Email: email, Name: name, Role: models.RoleUser }
	salt := utils.RandomSalt(); u.Password = utils.HashPassword(password, salt) + ":" + salt
	if err := s.users.Create(ctx, u); err != nil { return nil, err }
	return u, nil
}
func (s *AuthService) Login(ctx context.Context, email, password string) (string,string,*models.User,error){
	u, err := s.users.ByEmail(ctx, email); if err != nil { return "", "", nil, errors.New("invalid credentials") }
	parts := utils.Split2(u.Password, ':')
	if len(parts)!=2 || !utils.CheckPasswordHash(password, parts[1], parts[0]) { return "", "", nil, errors.New("invalid credentials") }
	access, err := utils.NewAccessToken(s.jwtSecret, u.ID, string(u.Role), s.accessTTL); if err != nil { return "", "", nil, err }
	refresh, err := utils.NewRefreshToken(s.jwtSecret, u.ID, s.refreshTTL); if err != nil { return "", "", nil, err }
	return access, refresh, u, nil
}
func (s *AuthService) Refresh(token string) (string, error){
	c, err := utils.ParseToken(s.jwtSecret, token); if err != nil { return "", err }
	if c.Role != "refresh" { return "", errors.New("not a refresh token") }
	return utils.NewAccessToken(s.jwtSecret, c.UserID, "user", s.accessTTL)
}
func (s *AuthService) UpsertOAuth(ctx context.Context, email, name, avatarURL, provider, providerID string) (*models.User, error) {
	return s.users.UpsertOAuthUser(ctx, email, name, avatarURL, provider, providerID)
}
