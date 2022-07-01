package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 0, data, msg)
}

func Fail(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, -1, data, msg)
}
