package models

type Category struct {
	ID       uint   `json:"id" gorm:"unique;"`
	Name     string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreateAt Time   `json:"createAt"`
	UpdateAt Time   `json:"updateAt"`
}
