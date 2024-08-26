package handler

import (
	"net/http"
	"on-page-seo/src/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteResults(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Converts the string into an integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return // Early return
	}

	err = repositories.DeleteResultID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "to delete result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": true})
}
