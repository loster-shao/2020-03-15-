package model

import "github.com/jinzhu/gorm"

type Voteuser struct{
	gorm.Model
	Id int `gorm:"primary_key"`
	Voteid int
	Xsuserid int
	Xsusername string
	Votetotalcount int
}
