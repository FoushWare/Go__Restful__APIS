package main

import "gorm.io/gorm"

type Post struct {
	Title string `gorm:"type:varchar(100);"`
	Des   string `gorm:"type:varchar(500);"`
	gorm.Model

}