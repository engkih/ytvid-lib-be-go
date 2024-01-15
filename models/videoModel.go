package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	UserId      int
	VideoUrl    string
	Title       string
	Description string
}
