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
	NumberWarning         = "It's not part of SEO best practices to have numbers in your URL, such as dates"
	OptimizedURLMessage   = "Optimized"
	MaxURLLength          = 25
	MinURLLength          = 6
)

// ValidateURL aggregates the results of individual validation functions
func ValidateURL(url string, keyword string) string {
	url = strings.TrimSpace(url)
	keyword = strings.TrimSpace(keyword)

	messages := []string{}

	lengthMessage := checkURLLength(url)
	if lengthMessage != "" {
		messages = append(messages, lengthMessage)
	}

	specialCharMessage := checkSpecialCharacters(url)
	if specialCharMessage != "" {
		messages = append(messages, specialCharMessage)
	}

	keywordMessage := checkKeywordInURL(url, keyword)
	if keywordMessage != "" {
		messages = append(messages, keywordMessage)
	}

	numberMessage := checkNumbers(url)
	if numberMessage != "" {
		messages = append(messages, NumberWarning)
	}

	if len(messages) == 0 {
		return OptimizedURLMessage
	}
	return strings.Join(messages, " | ")
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
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '-' && char != '/' && char != ' ' {
			return SpecialCharMessage
		}
	}
	return ""
}

func checkNumbers(url string) string {
	for _, char := range url {
		if unicode.IsDigit(char) {
			return NumberWarning
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
