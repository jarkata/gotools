package date_test

import (
	"fmt"
	"testing"

	"github.com/jarkata/gotools/date"
)

func TestParseDateTime(t *testing.T) {
	fmt.Println(date.ParseDateTime("2014-06-01T12:00:00.000Z"))
	fmt.Println(date.ParseDateTime("2014-06-01T12:00:00.000"))
	fmt.Println(date.ParseDateTime("2024-04-21T12:00:00.000"))
}
