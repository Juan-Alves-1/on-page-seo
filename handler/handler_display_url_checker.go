package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowChecker(c *gin.Context) {
	c.HTML(http.StatusOK, "checker.html", nil)
}
