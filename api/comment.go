package api

import (
	"prompt-share-backend/model"
	"prompt-share-backend/service"
	"prompt-share-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateComment 创建评论
// @Summary create comment
// @Tags comment
// @Produce json
// @Param id path int true "prompt id"
// @Param data body map[string]interface{} true "data"
// @Success 200 {object} model.Comment
// @Router /prompts/{id}/comments [post]
func CreateComment(c *gin.Context) {
	var in struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	pidStr := c.Param("id")
	pid, _ := strconv.ParseUint(pidStr, 10, 64)
	uidRaw, _ := c.Get("user_id")
	var uid uint
	if uidRaw != nil {
		uid = uidRaw.(uint)
	}
	com := &model.Comment{
		UserID:   uid,
		PromptID: uint(pid),
		Content:  in.Content,
	}
	if err := service.CreateComment(com); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, com)
}

// ListComments 获取评论列表
// @Summary list comments
// @Tags comment
// @Produce json
// @Param id path int true "prompt id"
// @Success 200 {object} []model.Comment
// @Router /prompts/{id}/comments [get]
func ListComments(c *gin.Context) {
	pidStr := c.Param("id")
	pid, _ := strconv.ParseUint(pidStr, 10, 64)
	list, err := service.ListComments(uint(pid))
	if err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, list)
}
