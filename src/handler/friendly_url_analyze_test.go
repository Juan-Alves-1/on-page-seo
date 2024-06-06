package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUrlCheckerAnalysis(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.LoadHTMLGlob("../../templates/*")
	router.POST("/url-checker/analyze", UrlCheckerAnalysis)

	// Create form data and add values to the form data
	form := url.Values{}
	form.Add("url", "https://example.com/test-my-keyword")
	form.Add("keyword", "my keyword")

	// The form data is encoded and passed as the body of the request.
	req, err := http.NewRequest(http.MethodPost, "/url-checker/analyze", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	//  Sets the Content-Type header to indicate that the request body contains URL-encoded form data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Contains(t, res.Body.String(), "example.com/test-my-keyword") // Check for the URL
	assert.Contains(t, res.Body.String(), "my keyword")                  // Check for the keyword
}
