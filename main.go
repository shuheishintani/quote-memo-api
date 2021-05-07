package main

import (
	"fmt"

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

	r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		c.Set("auth", auth)
	})
	r.Use(middleware.AuthMiddleware())

	service := services.NewService(db)
	controller := controllers.NewController(service)

	api := r.Group("/api")

	api.GET("/books", controller.GetBooks)
	api.GET("/tags", controller.GetTags)
	api.GET("/quotes", controller.GetQuotes)
	api.POST("/quotes", controller.PostQuote)

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
