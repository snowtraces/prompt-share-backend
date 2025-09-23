package model

import "time"

type Prompt struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"size:255;not null" json:"title"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	Tags       string    `gorm:"size:255" json:"tags"` // comma separated
	UserID     uint      `json:"user_id"`
	AuthorName string    `gorm:"size:100" json:"author_name"`
	LikeCount  int64     `gorm:"default:0" json:"like_count"`
	FavCount   int64     `gorm:"default:0" json:"fav_count"`
	SourceBy   string    `gorm:"size:100" json:"source_by"`
	SourceURL  string    `gorm:"size:255" json:"source_url"`
	SourceTags string    `gorm:"size:100" json:"source_tags"` // comma separated
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
