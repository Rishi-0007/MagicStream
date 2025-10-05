package services

import (
	"context"
	"github.com/rishi-0007/magicstream-backend/internal/models"
	"github.com/rishi-0007/magicstream-backend/internal/repository"
)

type MovieService struct { repo *repository.MovieRepository }
func NewMovieService(r *repository.MovieRepository) *MovieService { return &MovieService{repo: r} }
func (s *MovieService) List(ctx context.Context, limit int64)([]models.Movie,error){ return s.repo.List(ctx, limit) }
func (s *MovieService) ByIMDBID(ctx context.Context, id string)(*models.Movie,error){ return s.repo.ByIMDBID(ctx, id) }
func (s *MovieService) Search(ctx context.Context, q string, limit int64)([]models.Movie,error){ return s.repo.SearchByTitle(ctx, q, limit) }
