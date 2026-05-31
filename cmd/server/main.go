package main

import (
	"log"
	"net/http"

	"github.com/aminramezanipoor/url-shortener/internal/config"
	"github.com/aminramezanipoor/url-shortener/internal/database"
	"github.com/aminramezanipoor/url-shortener/internal/handler"
	"github.com/aminramezanipoor/url-shortener/internal/middleware"
	"github.com/aminramezanipoor/url-shortener/internal/repository"
	"github.com/aminramezanipoor/url-shortener/internal/routes"
	"github.com/aminramezanipoor/url-shortener/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal("cannot migrate database: ", err)
	}

	userRepo := repository.NewUserRepository(db)
	urlRepo := repository.NewURLRepository(db)

	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	urlService := service.NewURLService(urlRepo)

	authHandler := handler.NewAuthHandler(authService)
	urlHandler := handler.NewURLHandler(urlService)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "URL Shortener API is running",
		})
	})

	routes.SetupRoutes(r, authHandler, urlHandler, cfg.JWTSecret)

	authorized := r.Group("/protected")
	authorized.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		authorized.GET("/", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			role, _ := c.Get("role")

			c.JSON(http.StatusOK, gin.H{
				"user_id": userID,
				"role":    role,
			})
		})
	}

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}