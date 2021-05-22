package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shuheishintani/quote-memo-api/src/config"
	"github.com/shuheishintani/quote-memo-api/src/server"
)

func main() {
	if os.Getenv("APP_ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	db, err := config.GormConnect()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	auth, err := config.InitAuth()
	if err != nil {
		log.Fatal("Failed to init firebase auth: ", err)
	}

	r := server.SetRouter(db, auth)
	r.Run(":8080")
}
