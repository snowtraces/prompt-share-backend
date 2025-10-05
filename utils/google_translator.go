package utils

import (
	"context"
	"fmt"
	"io"
	"sync"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

var (
	translateClient *translate.Client
	translateOnce   sync.Once
)

// getTranslateClient 负责复用全局 client
func getTranslateClient(ctx context.Context) (*translate.Client, error) {
	var err error
	translateOnce.Do(func() {
		UseSystemProxy()
		translateClient, err = translate.NewClient(ctx)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init translate client: %w", err)
	}
	return translateClient, nil
}

// TranslateText 复用 client
func TranslateText(w io.Writer, targetLanguage, sourceLanguage, text string) error {
	ctx := context.Background()

	client, err := getTranslateClient(ctx)
	if err != nil {
		return err
	}

	// 解析目标语言
	targetLang, err := language.Parse(targetLanguage)
	if err != nil {
		return fmt.Errorf("language.Parse target: %w", err)
	}

	var opts *translate.Options
	if sourceLanguage != "" {
		sourceLang, err := language.Parse(sourceLanguage)
		if err != nil {
			return fmt.Errorf("language.Parse source: %w", err)
		}
		opts = &translate.Options{Source: sourceLang}
	}

	resp, err := client.Translate(ctx, []string{text}, targetLang, opts)
	if err != nil {
		return fmt.Errorf("client.Translate error: %w", err)
	}
	if len(resp) == 0 {
		return fmt.Errorf("client.Translate returned empty response")
	}

	fmt.Fprint(w, resp[0].Text)
	return nil
}
