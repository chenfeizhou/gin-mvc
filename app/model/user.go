package model

type User struct {
	baseModel
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(500);not null " json:"password"`
}

func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {

	var users []User
	var total int64

	if username != "" {

		db.Select("id,username,created_at").Where(
			"username like ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)

		db.Model(&users).Where(
			"username Like ?", username+"%",
		).Count(&total)

		return users, total
	}

	db.Select("id,username,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	return users, total
}
