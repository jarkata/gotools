package tid_test

import (
	"fmt"
	"testing"

	"gitee.com/tasphere/gotools/tid"
)

func TestGettime(t *testing.T) {
	fmt.Printf("tid.GetInitTime(): %v\n", tid.GetInitTime())
}
