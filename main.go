package main

import (
	"on-page-seo/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handler.ShowHomepage)

	r.GET("/url-checker", handler.ShowChecker)

	r.POST("/url-checker/analyze", handler.UrlCheckerAnalysis)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
