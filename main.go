package main

import (
	"fmt"
	"on-page-seo/config"
	"on-page-seo/database"
	"on-page-seo/internal/handler"
	"on-page-seo/internal/middleware"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Wasn't able to connect with the database: %s", err)
	}
	database.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/health", handler.Readiness)
	r.GET("/", handler.ShowHomepage)
	r.GET("/url-checker", middleware.UUIDCreationMiddleware(), handler.ShowChecker)

	authGroup := r.Group("/")
	authGroup.Use(middleware.UUIDValidationnMiddleware())
	{
		r.POST("/url-checker/analyze", handler.UrlCheckerAnalysis)
		r.POST("/api/save-results", handler.SaveResults)
		r.GET("/api/result/", handler.UserResults)
		r.DELETE("/api/result/:id", handler.DeleteResults)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
