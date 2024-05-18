package main

import (
	"strings"
	"unicode"
)

const (
	EmptyURLMessage       = "You must fill out the URL field with a slug"
	TooLongURLMessage     = "Too long - consider making it more concise"
	TooShortURLMessage    = "Too short - you might consider targeting a less competitive and more specific keyword"
	SpecialCharMessage    = "The URL contains special characters, consider removing them"
	KeywordMissingMessage = "The keyword is not present in the URL. Consider inserting it"
	OptimizedURLMessage   = "Optimized"
	MaxURLLength          = 25
	MinURLLength          = 6
)

// ValidateURL aggregates the results of individual validation functions
func ValidateURL(url string, keyword string) string {
	url = strings.TrimSpace(url)
	keyword = strings.TrimSpace(keyword)

	lengthMessage := checkURLLength(url)
	if lengthMessage != "" {
		return lengthMessage
	}

	specialCharMessage := checkSpecialCharacters(url)
	if specialCharMessage != "" {
		return specialCharMessage
	}

	keywordMessage := checkKeywordInURL(url, keyword)
	if keywordMessage != "" {
		return keywordMessage
	}

	return OptimizedURLMessage
}

// checkURLLength checks if the URL length is within acceptable limits
func checkURLLength(url string) string {
	normalizedURL := strings.ReplaceAll(url, "-", "")
	normalizedURL = strings.ReplaceAll(normalizedURL, "/", "")

	switch {
	case len(normalizedURL) == 0:
		return EmptyURLMessage
	case len(normalizedURL) > MaxURLLength:
		return TooLongURLMessage
	case len(normalizedURL) <= MinURLLength:
		return TooShortURLMessage
	default:
		return ""
	}
}

// checkSpecialCharacters checks if the URL contains special characters
func checkSpecialCharacters(url string) string {
	for _, char := range url {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '-' && char != '/' {
			return SpecialCharMessage
		}
	}
	return ""
}

// checkKeywordInURL checks if the keyword is present in the URL
func checkKeywordInURL(url string, keyword string) string {
	normalizedURL := strings.ReplaceAll(url, "-", " ")
	if !strings.Contains(normalizedURL, keyword) {
		return KeywordMissingMessage
	}
	return ""
}
