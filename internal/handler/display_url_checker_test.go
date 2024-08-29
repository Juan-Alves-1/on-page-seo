package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestShowChecker(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.LoadHTMLGlob("../../templates/*")
	router.GET("/url-checker", ShowChecker)

	req, err := http.NewRequest(http.MethodGet, "/url-checker", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Contains(t, res.Body.String(), "SEO friendly URL Checker")

}
