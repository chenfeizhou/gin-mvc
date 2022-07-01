package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/model"
	"github.com/gin-mvc/helpers"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 authorization header token
		tokenString := c.GetHeader("Authorization")

		// 验证tokenstring
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			helpers.Response(c, http.StatusUnprocessableEntity, 401, nil, "权限不足")
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := helpers.ParseToken(tokenString)

		if err != nil || !token.Valid {
			helpers.Response(c, http.StatusUnprocessableEntity, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 获取token userId
		userId := claims.UserId
		var user model.User
		model.DB.First(&user, userId)

		if user.ID == 0 {
			helpers.Response(c, http.StatusUnprocessableEntity, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 用户存在 将user信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}
