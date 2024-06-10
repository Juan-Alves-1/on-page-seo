package repositories

type ResultBody struct {
	PostForm   string `json:"post_form"`
	Slug       string `json:"slug"`
	Keyword    string `json:"keyword"`
	ResultForm string `json:"result_form"`
}

func SaveResults(resultBody ResultBody) error {
	query := "INSERT INTO results (url, slug, keyword, result) VALUES (?, ?, ?, ?)"
	_, err := database.db.Exec(query, resultBody.PostForm, resultBody.Slug, resultBody.Keyword, resultBody.ResultForm)

	return err
}
