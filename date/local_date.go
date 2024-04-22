package date

import (
	"errors"
	"fmt"
	"log/slog"
	"time"
)

type LocalDate struct {
	time.Time
}

func NewLocalDate(t time.Time) LocalDate {
	return LocalDate{t}
}

func (t LocalDate) ToTime() time.Time {
	return t.Time
}

/*
*
格式化日期
yyyy-MM-dd HH:mm:ss
*/
func (t LocalDate) MarshalJSON() ([]byte, error) {
	//格式化秒
	return []byte(fmt.Sprintf("\"%v\"", FormatDateOnly(t.Time))), nil
}

/** 解析JSON */
func (t *LocalDate) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		return errors.New("param invalid")
	}
	slog.Info("", "", data)
	tmpValue := string(data[1 : len(data)-1])
	value, err := ParseDateTime(tmpValue)
	if err != nil {
		return err
	}
	t.Time = value
	return nil
}
