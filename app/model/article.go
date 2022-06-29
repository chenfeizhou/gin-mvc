package model

type Article struct {
	baseModel
	User     User     `gorm:"ForeignKey:Uid"`
	Uid      uint     `gorm:"not null"`
	Category Category `gorm:"ForeignKey:cid"`
	Cid      uint     `gorm:"not null"`
	Title    string   `gorm:"type:varchar(100);not null " json:"title"`
	Content  string   `gorm:"type:longtext;not null " json:"content"`
}
