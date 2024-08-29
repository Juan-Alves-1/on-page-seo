package handler

import (
	"net/http"
	"on-page-seo/internal/repositories"
	"strings"

	"github.com/gin-gonic/gin"
)

func SaveResults(c *gin.Context) {
	var reqBody struct {
		URL      string   `json:"url"`
		Keyword  string   `json:"keyword"`
		Slug     string   `json:"slug"`
		Messages []string `json:"messages"`
		UUID     string   `json:"uuid"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID no found"})
	}

	uuidToken := strings.TrimPrefix(authHeader, "Bearer ")

	err := repositories.SaveResults(repositories.ResultBody{
		URL:     reqBody.URL,
		Keyword: reqBody.Keyword,
		Slug:    reqBody.Slug,
		Result:  reqBody.Messages,
		UUID:    uuidToken,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
