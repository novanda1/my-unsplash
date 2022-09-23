package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Label string
	Url   string
}
