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
		handler.UrlCheckerAnalyzer(c)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
