package models

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	Title string `json:"title" gorm:"title"`
}
