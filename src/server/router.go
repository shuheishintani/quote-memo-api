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

	public := r.Group("/api/public")
	public.Use(func(c *gin.Context) {
		c.Set("auth", auth)
	})
	public.GET("/users", controller.GetUsers)
	public.GET("/users/:id", controller.GetUserById)
	public.GET("/external_books", controller.GetExternalBooks)
	public.GET("/books", controller.GetBooks)
	public.GET("books/:id", controller.GetBookById)
	public.GET("/tags", controller.GetTags)
	public.GET("/quotes", controller.GetPublicQuotes)
	public.POST("/auth/login", controller.Login)

	private := r.Group("/api")
	private.Use(func(c *gin.Context) {
		c.Set("auth", auth)
	})
	private.Use(middleware.AuthMiddleware())
	private.POST("/users", controller.CreateOrUpdateUser)
	private.GET("/users/me", controller.GetMe)
	private.DELETE("/users", controller.DeleteUser)
	private.POST("/quotes", controller.PostQuote)
	private.GET("/quotes", controller.GetPrivateQuotes)
	private.GET("/quotes/favorite", controller.GetFavoriteQuotes)
	private.PUT("/quotes/:id", controller.UpdateQuote)
	private.DELETE("/quotes/:id", controller.DeleteQuote)
	private.PUT("/quotes/:id/like", controller.AddFavoriteQuote)
	private.PUT("/quotes/:id/unlike", controller.RemoveFavoriteQuote)

	return r
}
