package routes

import (
	"github.com/aminramezanipoor/url-shortener/internal/handler"
	"github.com/aminramezanipoor/url-shortener/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authHandler *handler.AuthHandler,
	urlHandler *handler.URLHandler,
	jwtSecret string,
) {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	urls := api.Group("/urls")
	urls.Use(middleware.AuthMiddleware(jwtSecret))
	{
		urls.POST("", urlHandler.CreateShortURL)
		urls.GET("/me", urlHandler.GetMyURLs)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(jwtSecret))
	admin.Use(middleware.RequireRole("admin"))
	{
		admin.GET("/urls", urlHandler.GetAllURLs)
		admin.DELETE("/urls/:id", urlHandler.DeleteURL)
	}

	r.GET("/:shortCode", urlHandler.Redirect)
}