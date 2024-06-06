package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestShowHomepage(t *testing.T) {
	// Step 1: Set up the Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.LoadHTMLGlob("../../templates/*")
	// Define the route to test
	router.GET("/", ShowHomepage)
	// Step 2: Create a new HTTP GET request
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	// Step 3: Record the response
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	// Step 4: Run the test and check the response
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Contains(t, res.Body.String(), "Welcome to this beautiful tool")

}
