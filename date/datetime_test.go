package date_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/jarkata/gotools/date"
)

func TestParseDateTime(t *testing.T) {
	fmt.Println(date.ParseDateTime("2014-06-01T12:00:00.000Z"))
	fmt.Println(date.ParseDateTime("2014-06-01T12:00:00.000"))
	fmt.Println(date.ParseDateTime("2024-04-21T12:00:00.000"))
}

type Demo struct {
	Date date.LocalDate `json:"date"`
}

func TestParseJsonLocalDate(t *testing.T) {
	tmp := make(map[string]date.LocalDate, 1)
	tmp["date"] = date.NewLocalDate(time.Now())
	jsonBytes, _ := json.Marshal(tmp)
	fmt.Println(string(jsonBytes))
	demo := Demo{}
	err := json.Unmarshal(jsonBytes, &demo)
	slog.Error("", "", err)

	slog.Info("", "本地时间", demo.Date)

}
