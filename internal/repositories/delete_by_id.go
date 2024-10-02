package repositories

import (
	"log"
	"on-page-seo/database"
)

func DeleteResultID(id int) error {
	query := "DELETE FROM results WHERE id = ?"

	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete result with ID %d: %v", id, err)
		return err
	}

	log.Printf("Successfully deleted result with ID %d", id)

	return nil
}
