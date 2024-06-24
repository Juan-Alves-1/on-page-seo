package repositories

import (
	"log"
	"on-page-seo/database"
	"on-page-seo/src/models"
)

func SaveResults(resultBody models.ResultBody) error {
	if err := database.DB.Create(&resultBody).Error; err != nil {
		log.Printf("Failed to execute query: %v", err)
		return err
	}
	return nil
}
