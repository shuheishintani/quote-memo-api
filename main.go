package main

import (
	"fmt"
	"os"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shuheishintani/quote-manager-api/src/config"
	"github.com/shuheishintani/quote-manager-api/src/controllers"
	"github.com/shuheishintani/quote-manager-api/src/middleware"
	"github.com/shuheishintani/quote-manager-api/src/services"
	"gorm.io/gorm"
)

func setRouter(db *gorm.DB, auth *auth.Client) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CLIENT_ORIGIN")},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	service := services.NewService(db)
	controller := controllers.NewController(service)

	public := r.Group("/api")
	public.GET("/books", controller.GetBooks)
	public.GET("/tags", controller.GetTags)

	private := r.Group("/api")

	private.Use(func(c *gin.Context) {
		c.Set("auth", auth)
	})
	private.Use(middleware.AuthMiddleware())

	private.GET("/quotes", controller.GetQuotes)
	private.POST("/quotes", controller.PostQuote)

	return r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	db, err := config.GormConnect()
	if err != nil {
		fmt.Println("Failed to connect database: ", err)
	}

	auth, err := config.InitAuth()
	if err != nil {
		fmt.Println("Failed to init firebase auth: ", err)
	}

	r := setRouter(db, auth)
	r.Run(":8080")
}
