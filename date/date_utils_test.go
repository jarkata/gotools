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
	Date      date.LocalDate `json:"date"`
	LocalTime date.LocalTime `json:"localTime"`
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

func TestParseTime(t *testing.T) {
	fmt.Println(date.ParseTime("12:01:40"))
	fmt.Println(date.ParseTime("120230"))
	fmt.Println(date.ParseTime("1202"))
}
func TestJsonTime(t *testing.T) {
	demo := Demo{LocalTime: date.NewLocalTime(time.Now())}
	jsonBytes, _ := json.Marshal(demo)
	fmt.Println(string(jsonBytes))
}
func TestParseJsonTime(t *testing.T) {
	demo := Demo{LocalTime: date.NewLocalTime(time.Now())}
	jsonBytes, _ := json.Marshal(demo)
	fmt.Println(string(jsonBytes))
	demo2 := Demo{}
	json.Unmarshal(jsonBytes, &demo2)
	slog.Info("", "解析的数据", demo2)

}
