package models

type User struct {
	ID       uint
	Name     string `gorm:"varchar(20)"`
	Password string `gorm:"varchar(20)"`
	Phone    string `gorm:"varchar(11)"`
}
