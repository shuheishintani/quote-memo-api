package server

import (
	"os"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-memo-api/src/controllers"
	"github.com/shuheishintani/quote-memo-api/src/middleware"
	"github.com/shuheishintani/quote-memo-api/src/services"
	"gorm.io/gorm"
)

func SetRouter(db *gorm.DB, auth *auth.Client) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CLIENT_ORIGIN")},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	service := services.NewService(db)
	controller := controllers.NewController(service)

	api := r.Group("/api")
	api.Use(func(c *gin.Context) {
		c.Set("auth", auth)
	})

	api.POST("/users", middleware.AuthMiddleware(), controller.CreateOrUpdateUser)
	api.GET("/users", controller.GetUsers)
	api.GET("/users/:id", controller.GetUserById)
	api.GET("/users/me", middleware.AuthMiddleware(), controller.GetMe)
	api.DELETE("/users", middleware.AuthMiddleware(), controller.DeleteUser)

	api.POST("/quotes", middleware.AuthMiddleware(), controller.PostQuote)
	api.GET("/quotes", controller.GetPublicQuotes)
	api.GET("/quotes/me", middleware.AuthMiddleware(), controller.GetPrivateQuotes)
	api.GET("/quotes/my_favorite", middleware.AuthMiddleware(), controller.GetFavoriteQuotes)
	api.GET("/quotes/for_export", middleware.AuthMiddleware(), controller.GetPrivateQuotesForExport)
	api.PUT("/quotes/:id", middleware.AuthMiddleware(), controller.UpdateQuote)
	api.DELETE("/quotes/:id", middleware.AuthMiddleware(), controller.DeleteQuote)
	api.PUT("/quotes/:id/like", middleware.AuthMiddleware(), controller.AddFavoriteQuote)
	api.PUT("/quotes/:id/unlike", middleware.AuthMiddleware(), controller.RemoveFavoriteQuote)

	api.GET("/books", controller.GetBooks)
	api.GET("books/:id", controller.GetBookById)
	api.GET("/external_books", controller.GetExternalBooks)

	api.GET("/tags", controller.GetTags)

	api.POST("/auth/login", controller.Login)
	api.POST("/auth/logout", controller.Logout)

	return r
}
