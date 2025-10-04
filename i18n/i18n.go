// i18n/i18n.go
package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

//go:embed locales/*
var localesFS embed.FS

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// 加载语言文件
	loadTranslationFiles()
}

func loadTranslationFiles() {
	entries, err := localesFS.ReadDir("locales")
	if err != nil {
		panic(fmt.Errorf("failed to read locales directory: %v", err))
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		path := "locales/" + filename

		if filepath.Ext(filename) == ".yaml" || filepath.Ext(filename) == ".json" {
			buf, err := localesFS.ReadFile(path)
			if err != nil {
				panic(fmt.Errorf("failed to read file %s: %v", path, err))
			}

			_, err = bundle.ParseMessageFileBytes(buf, path)
			if err != nil {
				panic(fmt.Errorf("failed to parse %s: %v", path, err))
			}
		}
	}
}

func GetLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, lang)
}
