package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Get the database connection information from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Form the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Verify the connection is valid
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established")
}

// if err = AutoMigrate(); err != nil {
// 	return err
// }

// func AutoMigrate() error {
// 	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

// 	if err := DB.AutoMigrate(Tables...); err != nil { // ENCONTRAR UMA FUNÇÃO DO SQL.DB COMPATIVEL A ESSA
// 		return err
// 	}
// 	return nil
// }

// var Tables = []interface{}{
// 	&repositories.ResultBody{},
// }
