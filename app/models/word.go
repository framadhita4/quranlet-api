package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	VerseID         uint   `json:"-" gorm:"not null"`
	LineNumber      uint   `json:"line_number" gorm:"not null"`
	AudioURL        string `json:"audio_url" gorm:"not null"`
	Text            string `json:"text" gorm:"not null"`
	Translation     string `json:"translation" gorm:"not null"`
	Transliteration string `json:"transliteration" gorm:"not null"`
}
