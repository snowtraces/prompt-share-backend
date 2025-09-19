package api

import (
	"prompt-share-backend/database"
	"prompt-share-backend/model"
	"prompt-share-backend/service"
	"prompt-share-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPrompts 列表
// @Summary list prompts
// @Tags prompts
// @Produce json
// @Param q query string false "query"
// @Param tag query string false "tag"
// @Param page query int false "page"
// @Param size query int false "page size"
// @Success 200 {object} map[string]interface{}
// @Router /prompts [get]
func GetPrompts(c *gin.Context) {
	q := c.Query("q")
	tag := c.Query("tag")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	list, total, err := service.QueryPrompts(q, tag, page, size)
	if err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, gin.H{"list": list, "total": total})
}

// CreatePrompt 创建
// @Summary create prompt
// @Tags prompts
// @Accept json
// @Produce json
// @Param prompt body model.Prompt true "prompt"
// @Success 200 {object} model.Prompt
// @Router /prompts [post]
func CreatePrompt(c *gin.Context) {
	var p model.Prompt
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	uidRaw, _ := c.Get("user_id")
	if uidRaw != nil {
		p.UserID = uidRaw.(uint)
	}
	if p.AuthorName == "" {
		p.AuthorName = "anonymous"
	}
	if err := service.CreatePrompt(&p); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, p)
}

// GetPrompt 详情
// @Summary get prompt
// @Tags prompts
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Prompt
// @Router /prompts/{id} [get]
func GetPrompt(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	p, err := service.GetPromptByID(uint(id))
	if err != nil {
		utils.Error(c, 1, "not found")
		return
	}
	utils.Success(c, p)
}

// UpdatePrompt 更新
// @Summary update prompt
// @Tags prompts
// @Accept json
// @Produce json
// @Param prompt body model.Prompt true "prompt"
// @Success 200 {object} model.Prompt
// @Router /prompts/{id} [put]
func UpdatePrompt(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var in model.Prompt
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	in.ID = uint(id)
	if err := database.DB.Save(&in).Error; err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, in)
}

// DeletePrompt 删除
// @Summary delete prompt
// @Tags prompts
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Router /prompts/{id} [delete]
func DeletePrompt(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if err := database.DB.Delete(&model.Prompt{}, uint(id)).Error; err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, gin.H{"deleted": id})
}

// LikePrompt 点赞
// @Summary like prompt
// @Tags prompts
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Router /prompts/{id}/like [post]
func LikePrompt(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if err := service.LikePrompt(uint(id)); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, gin.H{"ok": true})
}

// FavoritePrompt 收藏
// @Summary favorite prompt
// @Tags prompts
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Router /prompts/{id}/favorite [post]
func FavoritePrompt(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if err := service.FavoritePrompt(uint(id)); err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, gin.H{"ok": true})
}
