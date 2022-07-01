package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/model"
	"github.com/gin-mvc/helpers"
	"github.com/gin-mvc/middleware"
)

type AuthController struct {
}

func (auth *AuthController) Router(r *gin.Engine) {
	r.POST("/auth/register", auth.Register)
	r.POST("/auth/login", auth.Login)
	r.GET("/auth/user/info", middleware.JwtAuth(), auth.Info)
}

func (auth *AuthController) Register(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if new(model.User).IsUserExist(username) {
		helpers.Fail(c, "用户名已存在", nil)
	}

	newUser := model.User{
		Username: username,
		Password: helpers.Md5(password),
	}

	model.DB.Create(&newUser)

	// 发放token
	token, _ := helpers.ReleaseToken(newUser)

	model.DB.Model(&newUser).Update("token", token)

	helpers.Success(c, gin.H{"token": token}, "注册成功")
}

func (auth *AuthController) Login(c *gin.Context) {
	var requestUser = model.User{}
	c.Bind(&requestUser)

	passwd := helpers.Md5(requestUser.Password)

	var user model.User
	model.DB.Where(&model.User{Username: requestUser.Username, Password: passwd}).First(&user)

	if user.ID == 0 {
		helpers.Response(c, http.StatusUnprocessableEntity, 400, nil, "登录失败，用户名不存在")
	}

	// 发放token
	token, _ := helpers.ReleaseToken(user)

	helpers.Success(c, gin.H{"token": token}, "登录成功")
}

// 用户信息
func (auth *AuthController) Info(c *gin.Context) {

	user, _ := c.Get("user")

	helpers.Success(c, gin.H{"user": user}, "用户信息")
}
