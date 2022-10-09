package controller

import "github.com/gin-gonic/gin"

func GetAcceptLanguage(c *gin.Context) string {
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "en" //default language
	}
	return lang
}
