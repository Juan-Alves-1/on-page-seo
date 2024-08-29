package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowHomepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
