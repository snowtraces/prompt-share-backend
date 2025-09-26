package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": msg})
}

func ErrorWithHttpCode(c *gin.Context, httpCode int, code int, msg string) {
	c.JSON(httpCode, gin.H{"code": code, "message": msg})
}
