package handler

import (
	"net/http"
	"on-page-seo/database"

	"github.com/gin-gonic/gin"
)

// is there a way to retrieve the results from the previous endpoint instead of creating a new form?
func SaveResults(c *gin.Context) {
	url := c.PostForm("url")
	slug := c.PostForm("slug")
	keyword := c.PostForm("keyword")
	result := c.PostForm("result")

	// Insert the result into the database
	// why is this format though?
	err := database.SaveResult(url, slug, keyword, result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Result saved successfully"})
}
