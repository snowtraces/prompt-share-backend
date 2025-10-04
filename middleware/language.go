// middleware/language.go
package middleware

import (
	"prompt-share-backend/i18n"

	"github.com/gin-gonic/gin"
)

func Language() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Accept-Language
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "en"
		}
		localizer := i18n.GetLocalizer(lang)
		c.Set("localizer", localizer)
		c.Next()
	}
}
