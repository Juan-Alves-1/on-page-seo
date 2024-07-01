package handler

import (
	"net/http"

	controller "on-page-seo/src/controller"

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

	c.HTML(http.StatusOK, "url_results.html", gin.H{
		"URL":     url,
		"Slug":    slug,
		"Keyword": keyword,
		"Result":  urlResult,
	})
}
