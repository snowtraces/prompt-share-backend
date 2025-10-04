// utils/response.go
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}

func Error(c *gin.Context, code int, msg string) {
	localizer, _ := c.Get("localizer")
	if localizer == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": msg})
		return
	}
	localizedMsg := localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{
		MessageID: msg,
	})
	c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": localizedMsg})
}

func ErrorWithHttpCode(c *gin.Context, httpCode int, code int, msg string) {
	localizer, _ := c.Get("localizer")
	if localizer == nil {
		c.JSON(httpCode, gin.H{"code": code, "message": msg})
		return
	}
	localizedMsg := localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{
		MessageID: msg,
	})
	c.JSON(httpCode, gin.H{"code": code, "message": localizedMsg})
}
