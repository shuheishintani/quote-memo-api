package config

import (
	"fmt"
	"os"

	"github.com/shuheishintani/quote-memo-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormConnect() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	if os.Getenv("APP_ENV") == "development" {
		dbUsername := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbTable := os.Getenv("DB_TABLE")
		dbPort := os.Getenv("DB_PORT")

		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
		db, err := gorm.Open(postgres.Open(dsn))
		if err != nil {
			return db, err
		}
		// db.Migrator().DropTable("quotes")
		// db.Migrator().DropTable("books")
		// db.Migrator().DropTable("tags")
		// db.Migrator().DropTable("users")
		// db.Migrator().DropTable("quotes_tags")
		// db.Migrator().DropTable("users_quotes")
		db.AutoMigrate(&models.User{}, &models.Quote{}, &models.Book{}, &models.Tag{})
		return db, nil
	} else {
		dbUser := os.Getenv("DB_USER")
		dbPwd := os.Getenv("DB_PASS")
		instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
		dbName := os.Getenv("DB_NAME")

		socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
		if !isSet {
			socketDir = "/cloudsql"
		}

		dsn := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)
		db, err := gorm.Open(postgres.Open(dsn))
		if err != nil {
			return db, err
		}

		db.AutoMigrate(&models.User{}, &models.Quote{}, &models.Book{}, &models.Tag{})
		return db, nil
	}

}
