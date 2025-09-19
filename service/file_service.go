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
	"time"
)

var Store storage.Storage

func InitStorage() {
	base := config.Cfg.Storage.Local.BasePath
	Store = storage.NewLocalStorage(base)
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

	fi := &model.File{
		UploaderID: uploaderID,
		Path:       path,
		Name:       fh.Filename,
		Size:       fh.Size,
		Type:       fh.Header.Get("Content-Type"),
	}
	if err := database.DB.Create(fi).Error; err != nil {
		return nil, err
	}
	return fi, nil
}

func GetFileReader(path string) (io.ReadCloser, error) {
	return Store.Open(path)
}
