package repositories

import (
	"on-page-seo/database"
	"strings"
)

func GetResultsByUUID(uuid string) ([]ResultBody, error) {
	query := "SELECT url, slug, keyword, result FROM results WHERE uuid = ?"
	rows, err := database.DB.Query(query, uuid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var finalResults []ResultBody
	for rows.Next() {
		var alldata ResultBody
		var joinedMessages string
		if err := rows.Scan(&alldata.URL, &alldata.Slug, &alldata.Keyword, &joinedMessages); err != nil {
			return nil, err
		}
		alldata.Result = strings.Split(joinedMessages, " ")
		alldata.UUID = uuid
		finalResults = append(finalResults, alldata)
	}

	return finalResults, nil

}
