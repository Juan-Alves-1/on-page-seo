package handler

import (
	"net/http"
	"on-page-seo/src/repositories"
	"strings"

	"github.com/gin-gonic/gin"
)

func SaveResultsHandler(c *gin.Context) {
	var reqBody struct {
		URL      string   `json:"url"`
		Keyword  string   `json:"keyword"`
		Slug     string   `json:"slug"`
		Messages []string `json:"messages"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	joinedMessages := strings.Join(reqBody.Messages, " ")

	err := repositories.SaveResults(repositories.ResultBody{
		URL:     reqBody.URL,
		Keyword: reqBody.Keyword,
		Slug:    reqBody.Slug,
		Result:  joinedMessages,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
