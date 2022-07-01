package model

import "strconv"

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(20);not null " json:"username" `
	Password string `gorm:"type:varchar(500);not null " json:"password"`
}

//创建用户的请求
type CreateUserRequest struct {
	ID       string `form:"id" json:"id"`
	Username string `form:"username" json:"username" binding:"required,min=2,max=100"`
	Password string `form:"password" json:"password" binding:"required,max=1000"`
}

func (*User) TableName() string {
	return "users"
}

func GetUsers(username string, pageNum int, pageSize int) ([]User, int64) {

	var users []User
	var total int64

	if username != "" {
		DB.Select("id,username,created_at").Where(
			"username like ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)

		DB.Model(&users).Where(
			"username Like ?", username+"%",
		).Count(&total)

		return users, total
	}

	DB.Select("id,username,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	DB.Model(&users).Count(&total)

	return users, total
}

func GetUserDetail(id int) User {

	var user User
	DB.First(&user, "id = ?", id)
	return user
}

func CreateUser(creatUserRequest CreateUserRequest) User {

	DB.Create(&creatUserRequest)

	var user User

	id, _ := strconv.Atoi(creatUserRequest.ID)

	user = GetUserDetail(id)
	return user
}
