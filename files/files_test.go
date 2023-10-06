package files_test

import (
	"fmt"
	"testing"

	"github.com/jarkata/gotools/files"
)

func TestGetWd(t *testing.T) {

	fmt.Printf("files.GetWorkDirectory(): %v\n", files.GetWorkDirectory())
}
