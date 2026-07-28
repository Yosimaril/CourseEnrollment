package helpers

func GetLang(c *gin.Context) string {
	lang := c.GetHeader("Accept-Language")

	if lang == "" {
		return "en"
	}

	return lang
}
