package main

const (
	EmptyURLMessage     = "You must fill out the URL field with a slug"
	TooLongURLMessage   = "Too long - consider making it more concise"
	TooShortURLMessage  = "Too short - you might consider targeting a less competitive and more specific keyword"
	OptimizedURLMessage = "Optimised"
	MaxURLLength        = 25
	MinURLLength        = 6
)

func ValidateURL(url string) string {
	switch {
	case len(url) == 0:
		return EmptyURLMessage
	case len(url) > MaxURLLength:
		return TooLongURLMessage
	case len(url) <= MinURLLength:
		return TooShortURLMessage
	default:
		return OptimizedURLMessage
	}
}
