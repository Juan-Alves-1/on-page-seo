package handler

import (
	"net/http"
	"on-page-seo/src/repositories"

	"github.com/gin-gonic/gin"
)

func UserResults(c *gin.Context) {
	uuid, err := c.Cookie("session_uuid")
	if err != nil || uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID not found"})
		return
	}

	results, err := repositories.GetResultsByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't retrieve data from database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Results": results})

}
