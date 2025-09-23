package database

import (
	"log"
	"os"
	"path/filepath"
	"prompt-share-backend/config"
	"prompt-share-backend/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// ensure directory for sqlite and files exists
	dbPath := config.Cfg.Database.Path
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal("create db dir failed:", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	DB = db

	// Auto migrate
	if err := DB.AutoMigrate(
		&model.User{},
		&model.Prompt{},
		&model.Comment{},
		&model.File{},
		&model.PromptImg{},
	); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
}
