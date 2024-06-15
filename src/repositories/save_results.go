package repositories

import (
	"log"
	"on-page-seo/database"
)

type ResultBody struct {
	URL     string `json:"url"`
	Slug    string `json:"slug"`
	Keyword string `json:"keyword"`
	Result  string `json:"result"`
}

func SaveResults(resultBody ResultBody) error {
	query := "INSERT INTO results (url, slug, keyword, result) VALUES (?, ?, ?, ?)"
	_, err := database.DB.Exec(query, resultBody.URL, resultBody.Slug, resultBody.Keyword, resultBody.Result)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
	}
	return err
}
