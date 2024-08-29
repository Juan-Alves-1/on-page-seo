package handler

import (
	"net/http"
	"on-page-seo/internal/repositories"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserResults(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	uuidToken := strings.TrimPrefix(authHeader, "Bearer ")
	if uuidToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID not found"})
		return
	}

	results, err := repositories.GetResultsByUUID(uuidToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't retrieve data from database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Results": results})

}
