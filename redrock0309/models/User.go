package models

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	Id int `gorm:"primary_key"`
	Username string
	Password string
}