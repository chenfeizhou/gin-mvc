package model

type Category struct {
	baseModel
	Name string `gorm:"type:varchar(100);not null " json:"name"`
}
