package handler

func UrlCheckerAnalyzer(c *gin.Context) c *gin.Context{
	url := c.PostForm("url")
		keyword := c.PostForm("keyword")
		slug, err := ExtractSlug(url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid URL format",
			})
			return
		}

	urlResult := controller.ValidateURL(slug, keyword)
	return	c.HTML(http.StatusOK, "url_results.html", gin.H{
			"URL":     url,
			"Slug":    slug,
			"Keyword": keyword,
			"Result":  urlResult,
	})
}