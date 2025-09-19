package service

import (
	"prompt-share-backend/database"
	"prompt-share-backend/model"
	"strings"

	"gorm.io/gorm"
)

func CreatePrompt(p *model.Prompt) error {
	return database.DB.Create(p).Error
}

func GetPromptByID(id uint) (*model.Prompt, error) {
	var p model.Prompt
	if err := database.DB.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func QueryPrompts(q string, tag string, page, pageSize int) ([]model.Prompt, int64, error) {
	var list []model.Prompt
	var total int64
	db := database.DB.Model(&model.Prompt{})

	if q != "" {
		db = db.Where("title LIKE ? OR content LIKE ?", "%"+q+"%", "%"+q+"%")
	}
	if tag != "" {
		db = db.Where("tags LIKE ?", "%"+tag+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	if err := db.Order("created_at desc").Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func LikePrompt(id uint) error {
	return database.DB.Model(&model.Prompt{}).Where("id = ?", id).UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
}

func FavoritePrompt(id uint) error {
	return database.DB.Model(&model.Prompt{}).Where("id = ?", id).UpdateColumn("fav_count", gorm.Expr("fav_count + ?", 1)).Error
}

func ParseTags(tags []string) string {
	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
	}
	return strings.Join(tags, ",")
}
