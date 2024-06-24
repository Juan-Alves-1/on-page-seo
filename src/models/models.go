package models

type ResultBody struct {
	URL     string   `json:"url"`
	Slug    string   `json:"slug"`
	Keyword string   `json:"keyword"`
	Result  []string `json:"result"`
}
