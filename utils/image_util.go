package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
)

// GenerateThumbnail 生成缩略图并返回 Base64 字符串
// data: 原始图片字节
// maxSize: 最大边长度
// quality: JPEG 压缩质量 (1-100)
func GenerateThumbnail(data []byte, maxSize int, quality int) (string, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	// 计算缩放尺寸（等比缩放，最大边不超过 maxSize）
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width > height {
		if width > maxSize {
			height = height * maxSize / width
			width = maxSize
		}
	} else {
		if height > maxSize {
			width = width * maxSize / height
			height = maxSize
		}
	}

	// 缩放
	thumb := imaging.Resize(img, width, height, imaging.Lanczos)

	// 输出到 buffer
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, thumb, &jpeg.Options{Quality: quality})
	if err != nil {
		return "", err
	}

	// 转 base64
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Str, nil
}

func IsImage(contentType string) bool {
	imageTypes := []string{
		"image/jpeg",
		"image/jpg",
		"image/png",
		"image/gif",
		"image/webp",
		"image/bmp",
		"image/svg+xml",
	}

	for _, imgType := range imageTypes {
		if contentType == imgType {
			return true
		}
	}

	return false
}
