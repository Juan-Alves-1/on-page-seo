package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ShowChecker(c *gin.Context) {
	exisUUID, err := c.Cookie("session_uuid")
	if err != nil || exisUUID == "" {
		newUUID := uuid.New().String()

		c.SetCookie("session_uuid", newUUID, 24*30, "/", "localhost", false, true)
		c.HTML(http.StatusOK, "checker.html", gin.H{"UUID": newUUID})
	} else {
		c.HTML(http.StatusOK, "checker.html", gin.H{"UUID": exisUUID})
	}

}
