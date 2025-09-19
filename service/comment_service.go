package service

import (
	"prompt-share-backend/database"
	"prompt-share-backend/model"
)

func CreateComment(c *model.Comment) error {
	return database.DB.Create(c).Error
}

func ListComments(promptID uint) ([]model.Comment, error) {
	var list []model.Comment
	if err := database.DB.Where("prompt_id = ?", promptID).Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
