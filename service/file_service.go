package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"prompt-share-backend/config"
	"prompt-share-backend/database"
	"prompt-share-backend/model"
	"prompt-share-backend/storage"
	"prompt-share-backend/utils"
	"time"
)

var Store storage.Storage

func InitStorage() {
	base := config.Cfg.Storage.Local.BasePath
	//Store = storage.NewLocalStorage(base)
	Store = storage.NewSnowStorage(base)
	// ensure data dir
	_ = os.MkdirAll(base, 0755)
}

func SaveUploadedFile(fh *multipart.FileHeader, prefix string, uploaderID uint) (*model.File, error) {
	src, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	stored := filepath.Join(time.Now().Format("20060102"), fh.Filename)
	path, err := Store.Save(stored, src)
	if err != nil {
		return nil, err
	}

	// 2. 生成缩略图
	// 判断是否为图片
	contentType := fh.Header.Get("Content-Type")
	var thumbnail string
	if utils.IsImage(contentType) {
		reader, err := GetFileReader(path)
		if err != nil {
			return nil, err
		}
		rawImageData, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		thumbnail, err = utils.GenerateThumbnail(rawImageData, 360, 90)
		if err != nil {
			return nil, err
		}
	}

	fi := &model.File{
		UploaderID: uploaderID,
		Path:       path,
		Name:       fh.Filename,
		Size:       fh.Size,
		Type:       contentType,
		Thumbnail:  thumbnail,
	}
	if err := database.DB.Create(fi).Error; err != nil {
		return nil, err
	}
	return fi, nil
}

// 生成缩略图
func GenThumbnail(f *model.File) (string, error) {
	reader, err := GetFileReader(f.Path)
	if err != nil {
		return "", err
	}
	rawImageData, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	thumbnail, err := utils.GenerateThumbnail(rawImageData, 360, 90)
	if err != nil {
		return "", err
	}
	f.Thumbnail = thumbnail
	if err := database.DB.Updates(f).Error; err != nil {
		return "", err
	}
	return thumbnail, nil
}

func GetFileReader(path string) (io.ReadCloser, error) {
	return Store.Open(path)
}

func QueryFiles(q string, tag string, page int, pageSize int) ([]model.File, int64, error) {
	var list []model.File
	var total int64
	db := database.DB.Model(&model.File{})

	if q != "" {
		db = db.Where("name LIKE ?", "%"+q+"%")
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

// DeleteFile 删除文件
func DeleteFile(id uint) error {
	var f model.File
	if err := database.DB.First(&f, id).Error; err != nil {
		return err
	}

	// 删除文件存储中的实际文件
	if err := Store.Delete(f.Path); err != nil {
		return err
	}

	// 从数据库中删除记录
	return database.DB.Delete(&f).Error
}

func IsFileUsed(id uint) bool {
	var count int64
	database.DB.Model(&model.PromptImg{}).Where("file_id = ?", id).Count(&count)
	return count > 0
}
