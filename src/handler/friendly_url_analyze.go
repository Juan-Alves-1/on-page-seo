package handler

import (
	"net/http"
	"strings"

	controller "on-page-seo/src/controller"
	repositories "on-page-seo/src/repositories"

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

	urlResult := controller.ValidateURL(slug, keyword)
	joinedResult := strings.Join(urlResult, ",")

	// Save the results
	resultBody := repositories.ResultBody{
		URL:     url,
		Slug:    slug,
		Keyword: keyword,
		Result:  joinedResult,
	}
	if err := repositories.SaveResults(resultBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save results",
		})
		return
	}

	c.HTML(http.StatusOK, "url_results.html", gin.H{
		"URL":     url,
		"Slug":    slug,
		"Keyword": keyword,
		"Result":  urlResult,
	})
}
