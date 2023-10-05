package files

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

func GetWorkDirectory() string {
	var basePath = os.Getenv("BASE_PATH")
	if len(basePath) <= 0 {
		basePathTmp, _ := filepath.Abs(".")
		basePath = basePathTmp
	}
	if len(basePath) <= 0 {
		basePathTmp, _ := os.Getwd()
		basePath = basePathTmp
	}
	return basePath
}

func ReadContent(data []byte, splitChar byte) []string {
	var str []string
	buffer := bytes.NewBuffer(data)
	for {
		line, _ := buffer.ReadString(splitChar)
		line = strings.TrimSpace(line)
		if len(line) <= 0 {
			break
		}
		str = append(str, line)
	}
	return str
}
