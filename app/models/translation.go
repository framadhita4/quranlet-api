package models

import "gorm.io/gorm"

type Translation struct {
	gorm.Model
	VerseID      uint   `json:"-"`
	Text         string `json:"text"`
	ResourceName string `json:"resource_name"`
}
