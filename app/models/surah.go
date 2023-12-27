package models

import "gorm.io/gorm"

type translation struct {
	En string `json:"en" gorm:"not null"`
	Ar string `json:"ar" gorm:"not null"`
	Id string `json:"id" gorm:"not null"`
}

type audioFile struct {
	ID       uint   `json:"id" gorm:"not null;unique"`
	Chapter  uint   `json:"chapter" gorm:"not null;unique"`
	FileSize uint   `json:"file_size" gorm:"not null"`
	Format   string `json:"format" gorm:"not null"`
	AudioURL string `json:"audio_url" gorm:"not null"`
}

type Surah struct {
	gorm.Model
	SurahNumber uint        `json:"surah_number" gorm:"not null;primaryKey"`
	Name        string      `json:"name" gorm:"not null"`
	Translation translation `json:"translation" gorm:"embedded"`
	Ayahs       uint        `json:"ayahs" gorm:"not null"`
	Type        string      `json:"type" gorm:"not null"`
	AudioFile   audioFile   `json:"audio_file" gorm:"embedded"`
}
