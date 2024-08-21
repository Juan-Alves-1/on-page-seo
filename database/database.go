package database

import (
	"database/sql"
	"fmt"
	"log"
	"on-page-seo/config"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func InitDB() {
	var err error
	url := fmt.Sprintf("%s?authToken=%s", config.AppConfig.TursoURL, config.AppConfig.TursoToken)

	DB, err = sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	err = DB.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to db %s: %s", url, err)
		os.Exit(1)
	}

	log.Println("Connected to the Turso database successfully!")
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
