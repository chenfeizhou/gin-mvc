package repository

import (
	"strconv"

	"github.com/gin-mvc/app/model"
)

type UserRepository struct {
}

func (userRepository *UserRepository) GetUsers(username string, pageNum string, pageSize string) ([]model.User, int64) {

	var users []model.User
	var total int64

	page, _ := strconv.Atoi(pageNum)
	page_size, _ := strconv.Atoi(pageSize)

	users, total = model.GetUsers(username, page, page_size)

	return users, total
}

func (userRepository *UserRepository) GetUserDetail(id int) model.User {
	return model.GetUserDetail(id)
}

func (userRepository *UserRepository) CreateUser(createUserRequest model.CreateUserRequest) model.User {
	return model.CreateUser(createUserRequest)
}

func (userRepository *UserRepository) DeleteUser(id int) {
	model.DB.Delete(&model.User{}, id)
}
