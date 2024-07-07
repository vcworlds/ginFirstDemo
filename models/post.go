package models

import "time"

type Post struct {
	Id         uint   `json:"id" gorm:"unique;primary_key"`
	Title      string `json:"title" gorm:"type:varchar(40);not null"`
	Content    string `json:"content" gorm:"not null;type:text"`
	UserId     uint   `json:"userId" gorm:"not null"`
	User       *User
	CategoryId uint `json:"categoryId" gorm:"not null"`
	Category   *Category
	PostImg    string    `json:"postImg" gorm:"not null"`
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
}
