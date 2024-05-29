package controller

import (
	"strings"
	"testing"
)

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		keyword  string
		expected string
	}{
		{
			name:     "Valid URL with keyword",
			url:      "/my-optimized-url",
			keyword:  "optimized",
			expected: "Well done! Your URL looks great",
		},
		{
			name:     "Too long URL",
			url:      "/this-is-a-very-long-url-that-should-be-shorter",
			keyword:  "shorter",
			expected: "Too long - consider making it more concise",
		},
		{
			name:     "URL with special characters",
			url:      "/my@bad$url",
			keyword:  "bad",
			expected: "The URL contains special characters, consider removing them",
		},
		{
			name:     "Keyword missing",
			url:      "/my-simple-url",
			keyword:  "missing",
			expected: "The keyword is not present in the URL. Consider inserting it",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateURL(tt.url, tt.keyword)
			joinedResult := strings.Join(result, "\n")
			if joinedResult != tt.expected {
				t.Errorf("ValidateURL(%s, %s) = %v; expected %v", tt.url, tt.keyword, joinedResult, tt.expected)
			}
		})
	}
}
