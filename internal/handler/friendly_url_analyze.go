package handler

import (
	"encoding/json"
	"net/http"

	controller "on-page-seo/internal/controller"

	"github.com/gin-gonic/gin"
)

func UrlCheckerAnalysis(c *gin.Context) {
	url := c.PostForm("url")
	keyword := c.PostForm("keyword")
	slug, err := controller.ExtractSlug(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid URL format",
		})
		return
	}

	urlResult := controller.ValidateSlug(url, keyword, slug)

	// Convert the result array to a JSON string for JavaScript usage
	resultJSON, err := json.Marshal(urlResult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JSON"})
		return
	}

	c.HTML(http.StatusOK, "url_results.html", gin.H{
		"URL":        url,
		"Slug":       slug,
		"Keyword":    keyword,
		"Result":     urlResult,          // For HTML rendering
		"ResultJSON": string(resultJSON), // For JavaScript
	})
}
