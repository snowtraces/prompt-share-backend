package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// 加载语言文件
	loadTranslationFiles()
}

func loadTranslationFiles() {
	// 假设语言文件放在 locales 目录下
	localesDir := "locales"
	err := filepath.Walk(localesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".json" {
			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			_, err = bundle.ParseMessageFileBytes(buf, path)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %v", path, err)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func GetLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, lang)
}
