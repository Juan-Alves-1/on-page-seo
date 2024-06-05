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

func TestCheckURLLength(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{url: "/short", expected: TooShortURLMessage},
		{url: "/it-will-be-sooooooo-long-that-fails", expected: TooLongURLMessage},
		{url: "/a-good-one", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			result := checkURLLength(tt.url)
			if result != tt.expected {
				t.Errorf("checkURLLength(%s) = %v; expected %v", tt.url, result, tt.expected)
			}
		})
	}
}

func TestCheckSpecialCharacters(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{url: "say-@-instead-of-a", expected: SpecialCharMessage},
		{url: "what-does-!-mean", expected: SpecialCharMessage},
		{url: "currency-€-meaning", expected: SpecialCharMessage},
		{url: "coding-what-does-<>-means", expected: SpecialCharMessage},
		{url: "we-are-in-2022", expected: ""},
		{url: "i-have/two-directories", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			result := checkSpecialCharacters(tt.url)
			if result != tt.expected {
				t.Errorf("checkSpecialCharacters(%s) = %v; expected: %v", tt.url, result, tt.expected)
			}
		})
	}
}

func TestCheckNumbers(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"no-numbers", ""},
		{"contains-123", NumberWarning},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			result := checkNumbers(tt.url)
			if result != tt.expected {
				t.Errorf("checkNumbers(%s) = %v; expected %v", tt.url, result, tt.expected)
			}
		})
	}
}

func TestCheckKeywordInURL(t *testing.T) {
	tests := []struct {
		url      string
		keyword  string
		expected string
	}{
		{"keyword-in-url", "keyword", ""},
		{"no-keyword-here", "missing", KeywordMissingMessage},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			result := checkKeywordInURL(tt.url, tt.keyword)
			if result != tt.expected {
				t.Errorf("checkKeywordInURL(%s, %s) = %v; expected %v", tt.url, tt.keyword, result, tt.expected)
			}
		})
	}
}

func TestExtractSlug(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://example.com/slug", "slug"},
		{"https://example.com/another-slug", "another-slug"},
		{"example.com/missing-protocol", "missing-protocol"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ExtractSlug(tt.input)
			if err != nil {
				t.Errorf("ExtractSlug(%s) returned an error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("ExtractSlug(%s) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
