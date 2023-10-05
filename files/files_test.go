package files_test

import (
	"fmt"
	"testing"

	"gitee.com/tasphere/gotools/files"
)

func TestGetWd(t *testing.T) {

	fmt.Printf("files.GetWorkDirectory(): %v\n", files.GetWorkDirectory())
}
