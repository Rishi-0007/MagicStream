package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rishi-0007/magicstream-backend/internal/config"
	"github.com/rishi-0007/magicstream-backend/internal/database"
	"github.com/rishi-0007/magicstream-backend/internal/handlers"
	"github.com/rishi-0007/magicstream-backend/internal/logger"
	"github.com/rishi-0007/magicstream-backend/internal/middleware"
	"github.com/rishi-0007/magicstream-backend/internal/repository"
	"github.com/rishi-0007/magicstream-backend/internal/services"
	"golang.org/x/time/rate"
)

func main() {
	cfg := config.Load()
	logg := logger.New()

	ctx := context.Background()
	mongo, err := database.Connect(ctx, cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Close(ctx)

	userRepo := repository.NewUserRepository(mongo.DB)
	movieRepo := repository.NewMovieRepository(mongo.DB)

	authSvc := services.NewAuthService(userRepo, cfg.JWTSecret, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)
	movieSvc := services.NewMovieService(movieRepo)

	authH := handlers.NewAuthHandler(authSvc)
	movieH := handlers.NewMovieHandler(movieSvc)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORS(cfg.CORSOrigins))
	r.Use(middleware.RateLimit(rate.Limit(10), 50))
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logg.Printf("%s %s %d %v", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	})

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

	r.POST("/api/auth/register", authH.Register)
	r.POST("/api/auth/login", authH.Login)
	r.POST("/api/auth/refresh", authH.Refresh)
	r.POST("/api/auth/oauth/google", authH.OAuthGoogle)

	r.GET("/api/movies", movieH.List)
	r.GET("/api/movies/:imdb_id", movieH.Get)
	r.GET("/api/search", movieH.Search)

	srv := &http.Server{Addr: ":" + cfg.Port, Handler: r}
	logg.Printf("listening on :%s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
