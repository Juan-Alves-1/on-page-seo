package database

import (
	"log"
	"on-page-seo/src/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = AutoMigrate(); err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	log.Println("Database connection established")
}

func AutoMigrate() error {
	if err := DB.AutoMigrate(Tables...); err != nil {
		return err
	}
	return nil
}

var Tables = []interface{}{
	&models.ResultBody{},
}
