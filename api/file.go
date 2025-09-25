package api

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"prompt-share-backend/database"
	"prompt-share-backend/model"
	"prompt-share-backend/service"
	"prompt-share-backend/utils"
	"strconv"
	"time"

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

// PreviewFile 预览文件
// @Summary 预览文件
// @Description 预览文件
// @Tags files
// @Accept  json
// @Route /files/preview/{id} [get]
func PreviewFile(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	var f model.File

	if err := database.DB.First(&f, uint(id)).Error; err != nil {
		utils.Error(c, 1, "file not found")
		return
	}

	c.Header("Cache-Control", "public, max-age=31536000, immutable, s-maxage=31536000")
	c.Header("Expires", time.Now().AddDate(1, 0, 0).UTC().Format(http.TimeFormat))

	rc, err := service.GetFileReader(f.Path)
	if err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	defer rc.Close()
	c.Header("Content-Type", f.Type)
	c.Header("Content-Length", strconv.FormatInt(f.Size, 10))
	if _, err := io.Copy(c.Writer, rc); err != nil {
		// 可以加日志
		fmt.Println("预览文件写出失败:", err)
	}
}

// Thumbnail 预览文件
// @Summary 预览文件
// @Description 预览文件
// @Tags files
// @Accept  json
// @Route /files/thumbnail/{id} [get]
func Thumbnail(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	var f model.File

	if err := database.DB.First(&f, uint(id)).Error; err != nil {
		utils.Error(c, 1, "file not found")
		return
	}

	thumbnail := f.Thumbnail
	if thumbnail == "" {
		genThumbnail, err := service.GenThumbnail(&f)
		if err != nil {
			utils.Error(c, 1, "generate thumbnail failed")
			return
		}
		thumbnail = genThumbnail
	}

	// thumbnail 为base64字符串，转图片输出
	// 解码 base64 字符串
	data, err := base64.StdEncoding.DecodeString(thumbnail)
	if err != nil {
		utils.Error(c, 1, "decode thumbnail failed")
		return
	}

	// 设置响应头
	c.Header("Content-Type", "image/jpeg") // 根据实际图像类型调整 MIME 类型
	c.Header("Cache-Control", "public, max-age=31536000, immutable, s-maxage=31536000")
	c.Header("Expires", time.Now().AddDate(1, 0, 0).UTC().Format(http.TimeFormat))

	// 输出图像数据
	c.Data(http.StatusOK, "image/jpeg", data)
}

// ListFiles 获取文件列表
func ListFiles(c *gin.Context) {
	q := c.Query("q")
	tag := c.Query("tag")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	list, total, err := service.QueryFiles(q, tag, page, size)
	if err != nil {
		utils.Error(c, 1, err.Error())
		return
	}
	utils.Success(c, gin.H{"list": list, "total": total})
}
