package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/model"
	"github.com/gin-mvc/app/repository"
	"github.com/gin-mvc/helpers"
)

type UserController struct {
}

func (userController *UserController) Router(r *gin.Engine) {
	r.GET("/users", userController.Index)
	r.GET("/users/:id", userController.Show)
	r.POST("/users", userController.Create)
	r.DELETE("/users/:id", userController.Destory)
}

func (userController *UserController) Index(c *gin.Context) {

	username := c.DefaultQuery("user_name", "")
	pageSize := c.DefaultQuery("page_size", "20")
	pageNum := c.DefaultQuery("page", "1")

	var data interface{}
	var total int64

	data, total = new(repository.UserRepository).GetUsers(username, pageNum, pageSize)

	helpers.Success(c, gin.H{"list": data, "count": total}, "success")
}

func (userController *UserController) Show(c *gin.Context) {

	type UserRequest struct {
		ID string `uri:"id"`
	}

	var request UserRequest

	if err := c.ShouldBindUri(&request); err != nil {
		return
	}

	userId, _ := strconv.Atoi(request.ID)

	user := new(repository.UserRepository).GetUserDetail(userId)

	helpers.Success(c, gin.H{"data": user}, "success")

}

func (userController *UserController) Create(c *gin.Context) {

	var creatUserRequest model.CreateUserRequest

	if c.ShouldBind(&creatUserRequest) == nil {
		return
	}

	user := new(repository.UserRepository).CreateUser(creatUserRequest)

	helpers.Success(c, gin.H{"data": user}, "success")
}

func (userController *UserController) Destory(c *gin.Context) {

	type UserRequest struct {
		ID string `uri:"id"`
	}

	var request UserRequest

	if err := c.ShouldBindUri(&request); err != nil {
		return
	}

	userId, _ := strconv.Atoi(request.ID)

	new(repository.UserRepository).DeleteUser(userId)

	helpers.Success(c, gin.H{"data": nil}, "success")
}
