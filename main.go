package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/url-checker", func(c *gin.Context) {
		c.HTML(http.StatusOK, "checker.html", nil)
	})

	// Define a route to handle the form submission
	r.POST("/url-checker/analyze", func(c *gin.Context) {
		url := c.PostForm("url")
		keyword := c.PostForm("keyword")
		slug, err := ExtractSlug(url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid URL format",
			})
			return
		}

		urlResult := ValidateURL(slug, keyword)
		c.HTML(http.StatusOK, "url_results.html", gin.H{
			"URL":     url,
			"Slug":    slug,
			"Keyword": keyword,
			"Result":  urlResult,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
