package api

import (
	"io"
	"path/filepath"
	"prompt-share-backend/database"
	"prompt-share-backend/model"
	"prompt-share-backend/service"
	"prompt-share-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UploadFile 上传
// @Summary upload file
// @Tags files
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} model.File
// @Router /files/upload [post]
func UploadFile(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	uidRaw, _ := c.Get("user_id")
	var uid uint
	if uidRaw != nil {
		uid = uidRaw.(uint)
	}
	fi, err := service.SaveUploadedFile(fh, "up", uid)
	if err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, fi)
}

// DownloadFile 上传
// @Summary upload file
// @Tags files
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} model.File
// @Router /files/download/{id} [get]
func DownloadFile(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var f model.File
	if err := database.DB.First(&f, uint(id)).Error; err != nil {
		utils.Error(c, 1, "file not found")
		return
	}
	rc, err := service.GetFileReader(f.Path)
	if err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	defer rc.Close()
	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(f.Name))
	c.Header("Content-Type", f.Type)
	c.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, rc)
		return err == nil
	})
}
