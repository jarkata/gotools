package number

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
)

type Decimal struct {
	Value uint64
}

func NewDecimal(value uint64) Decimal {
	return Decimal{Value: value}
}

func (t Decimal) ToInt64() uint64 {
	return t.Value
}

/*
*
 */
func (t Decimal) MarshalJSON() ([]byte, error) {
	//格式化秒
	return []byte(fmt.Sprintf("\"%v\"", FormatDecimal(t.Value))), nil
}

/** 解析JSON */
func (t *Decimal) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		return errors.New("param invalid")
	}
	slog.Info("数据", "{}", data)
	tmpValue := string(data[1 : len(data)-1])
	value, err := ParseDecimal(tmpValue)
	if err != nil {
		return err
	}
	t.Value = value
	return nil
}

func ParseDecimal(value string) (uint64, error) {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return uint64(result * 100), nil
}

func FormatDecimal(value uint64) string {
	return strconv.FormatFloat(float64(value)/100, 'f', 2, 64)
}
