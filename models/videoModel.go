package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	VideoUrl    string
	Title       string
	Description string
}
