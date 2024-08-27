package repositories

import (
	"log"
	"on-page-seo/database"
	"strings"
)

type ResultBody struct {
	URL     string   `json:"url"`
	Slug    string   `json:"slug"`
	Keyword string   `json:"keyword"`
	Result  []string `json:"result"`
	UUID    string   `json:"uuid"`
}

func SaveResults(resultBody ResultBody) error {
	query := "INSERT INTO results (url, slug, keyword, result, uuid) VALUES (?, ?, ?, ?, ?)"
	joinedMessages := strings.Join(resultBody.Result, " ")
	_, err := database.DB.Exec(query, resultBody.URL, resultBody.Slug, resultBody.Keyword, joinedMessages, resultBody.UUID)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
	}
	return err
}
