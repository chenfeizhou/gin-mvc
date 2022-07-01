package model

type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(100);not null"`
}

func (*Category) TableName() string {
	return "category"
}
