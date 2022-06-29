package model

type User struct {
	baseModel
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(500);not null " json:"password"`
}

func (table *User) TableName() string {
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
