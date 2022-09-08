package model

import "gorm.io/gorm"

type Todo struct {
	ID     uint   `json:"id" gorm:"primaryKey; not null"`
	Text   string `json:"text"`
	UserID *uint  `json:"userId"`
	User   *User  `json:"user"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Todos    []Todo `json:"todos"`
}
