package repository

import (
	"strconv"

	"github.com/gin-mvc/app/model"
)

type User struct {
}

func (user *User) GetUsers(username string, pageNum string, pageSize string) ([]model.User, int64) {

	var users []model.User
	var total int64

	page, _ := strconv.Atoi(pageNum)
	page_size, _ := strconv.Atoi(pageSize)

	users, total = model.GetUsers(username, page, page_size)

	return users, total
}
