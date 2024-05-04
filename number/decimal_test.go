package number_test

import (
	"encoding/json"
	"fmt"

	"strconv"
	"testing"

	"github.com/jarkata/gotools/number"
)

func TestFormat(t *testing.T) {
	decmail := number.Decimal{Value: 123}
	json.Marshal(decmail)
}

func TestFloat(t *testing.T) {
	fmt.Printf("strconv.FormatFloat(float64(1238/100), 'f', 2, 64): %v\n", fmt.Sprintf("%.2f", float64(1238)/100))
	fmt.Printf("strconv.FormatFloat(float64(1238/100), 'f', 2, 64): %v\n", strconv.FormatFloat(float64(1238)/100, 'f', 2, 64))
}
