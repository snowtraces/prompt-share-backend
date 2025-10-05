package main

import (
	"bytes"
	"fmt"
	"prompt-share-backend/utils"
)

func main() {
	var buf bytes.Buffer
	err := utils.TranslateText(&buf, "zh", "", "Hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(buf.String())
	}

}
