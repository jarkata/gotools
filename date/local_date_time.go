package date

import (
	"errors"
	"fmt"
	"log/slog"
	"time"
)

type LocalTime struct {
	time.Time
}

func New(t time.Time) LocalTime {
	return LocalTime{t}
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	return []byte(fmt.Sprintf("\"%v\"", t.Time.Format("2006-01-02 15:04:05"))), nil
}
func (t *LocalTime) UnmarshalJSON(data []byte) error {

	if len(data) < 1 {
		return errors.New("param invalid")
	}
	tmp := string(data[1 : len(data)-1])

	value, err := time.ParseInLocation("2006-01-02T15:04:05.999Z", tmp, time.Local)
	slog.Info("", "Value:", value, "ERR:", err)
	slog.Info("", "格式化", value.Local().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	t.Time = value
	return nil
}
