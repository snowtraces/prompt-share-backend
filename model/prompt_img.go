package model

type PromptImg struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	PromptID uint   `json:"prompt_id"`
	FileId   uint   `json:"file_id"`
	Tags     string `gorm:"size:255" json:"tags"` // comma separated
	FileUrl  string `gorm:"size:255" json:"file_url"`
}
