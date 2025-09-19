package service

import (
	"prompt-share-backend/config"
	"prompt-share-backend/database"
	"prompt-share-backend/model"
	"prompt-share-backend/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RegisterUser(u *model.User, plainPwd string) error {
	hash, err := utils.HashPassword(plainPwd)
	if err != nil {
		return err
	}
	u.PasswordHash = hash
	return database.DB.Create(u).Error
}

func Login(username, password string) (string, *model.User, error) {
	var u model.User
	err := database.DB.Where("username = ?", username).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil, err
		}
		return "", nil, err
	}
	if !utils.CheckPassword(u.PasswordHash, password) {
		return "", nil, gorm.ErrInvalidData
	}

	claims := jwt.MapClaims{
		"uid": u.ID,
		"exp": time.Now().Add(time.Duration(config.Cfg.JWT.ExpireHours) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString([]byte(config.Cfg.JWT.Secret))
	return ts, &u, err
}
