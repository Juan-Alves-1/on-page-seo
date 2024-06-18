package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// if err = AutoMigrate(); err != nil {
	// 	return err
	// }

	log.Println("Database connection established")
}

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
