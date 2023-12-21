package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	VideoId  int
	Username string
	Comment  string
}
