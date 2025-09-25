package model

import "time"

type File struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UploaderID uint      `json:"uploader_id"`
	Path       string    `gorm:"size:512" json:"path"`
	Name       string    `gorm:"size:255" json:"name"`
	Size       int64     `json:"size"`
	Type       string    `gorm:"size:100" json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	Thumbnail  string    `gorm:"blob" json:"thumbnail"`
}
