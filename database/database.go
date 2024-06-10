package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	dsn := "username:password@tcp(127.0.0.1:3306)/friendly_url_results"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established")
}

// is it in the right place? and double-check the difference with SaveResults
func SaveResult(url, slug, keyword, result string) error {
	query := "INSERT INTO results (url, slug, keyword, result) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, url, slug, keyword, result)
	return err
}
