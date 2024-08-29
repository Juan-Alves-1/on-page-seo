package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Readiness(c *gin.Context) {
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
