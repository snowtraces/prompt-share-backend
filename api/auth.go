package api

import (
	"prompt-share-backend/model"
	"prompt-share-backend/service"
	"prompt-share-backend/utils"

	"github.com/gin-gonic/gin"
)

// Register 注册
// @Summary Register
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "user"
// @Success 200 {object} map[string]interface{}
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var in struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	u := &model.User{
		Username: in.Username,
		Email:    in.Email,
	}
	if err := service.RegisterUser(u, in.Password); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, gin.H{"message": "registered"})
}

// Login 登录
// @Summary Login
// @Tags auth
// @Accept json
// @Produce json
// @Param cred body object true "credentials"
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var in struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	token, user, err := service.Login(in.Username, in.Password)
	if err != nil {
		utils.Error(c, 1, "username or password invalid")
		return
	}
	// hide password
	user.PasswordHash = ""
	utils.Success(c, gin.H{"token": token, "user": user})
}
