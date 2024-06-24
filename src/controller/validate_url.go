package controller

import (
	"net/url"
	"on-page-seo/src/repositories"
	"path"
	"strings"
	"unicode"
)

const (
	EmptyURLMessage       = "You must fill out the URL field with a slug"
	TooLongURLMessage     = "Too long - consider making it more concise"
	TooShortURLMessage    = "Too short - you might consider targeting a less competitive and more specific keyword"
	SpecialCharMessage    = "The URL contains special characters, consider removing them"
	KeywordMissingMessage = "The keyword is not present in the URL. Consider inserting it"
	NumberWarning         = "It's not part of SEO best practices to have numbers such as dates in your URL"
	OptimizedURLMessage   = "Well done! Your URL looks great"
	MaxURLLength          = 25
	MinURLLength          = 6
)

func ExtractSlug(input string) (string, error) {
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		input = "http://" + input // Prepend default protocol if missing
	}

	inp, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	pathSegments := path.Clean(inp.Path)
	lastSegment := path.Base(pathSegments)
	return lastSegment, nil
}

// ValidateURL aggregates the results of individual validation functions
func ValidateURL(url string, keyword string, slug string) []string {
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

	resultBody := repositories.ResultBody{
		URL:     url,
		Keyword: keyword,
		Slug:    slug,
		Result:  messages,
	}

	err := repositories.SaveResults(resultBody)
	if err != nil {
		return []string{err.Error()}
	}

	if len(messages) == 0 {
		return []string{OptimizedURLMessage}
	}
	return messages
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

// Checks if there are numbers in the url
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
