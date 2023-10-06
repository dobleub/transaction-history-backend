package helpers

import (
	"encoding/json"
	"math"
	"strconv"
	"time"
	"unicode"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/leekchan/accounting"
)

func StringToInt32(str string) int32 {
	i, _ := strconv.Atoi(str)
	return int32(i)
}

func StringToInt64(str string) int64 {
	i, _ := strconv.Atoi(str)
	return int64(i)
}

func StringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func StringToBool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}

func StringToDate(str string) time.Time {
	t, _ := time.Parse("2006-01-02", str)
	return t
}

func StringToTime(str string) time.Time {
	t, _ := time.Parse(time.RFC3339, str)
	return t
}

func StringToTimeWithFormat(str string, format string) time.Time {
	t, _ := time.Parse(format, str)
	return t
}

func StringToTimeWithLocation(str string, location *time.Location) time.Time {
	t, _ := time.ParseInLocation(time.RFC3339, str, location)
	return t
}

func TimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}

func TimeToStringWithFormat(t time.Time, format string) string {
	return t.Format(format)
}

func TimeToStringWithLocation(t time.Time, location *time.Location) string {
	return t.In(location).Format(time.RFC3339)
}

func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.UnixNano() - t.Unix()*1e9),
	}
}

func TimestampToTime(ts *timestamp.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}

func TimestampToString(ts *timestamp.Timestamp) string {
	return TimeToString(TimestampToTime(ts))
}

func TimestampToStringWithFormat(ts *timestamp.Timestamp, format string) string {
	return TimeToStringWithFormat(TimestampToTime(ts), format)
}

func TimestampToStringWithLocation(ts *timestamp.Timestamp, location *time.Location) string {
	return TimeToStringWithLocation(TimestampToTime(ts), location)
}

func DecimalAdjust(value float64, exp int) float64 {
	if exp < 0 {
		return value
	}
	return float64(int64(value*float64(exp))) / float64(exp)
}

func Round(value float64, precision int) float64 {
	return DecimalAdjust(value, int(math.Pow10(precision)))
}

func FormatMoney(value float64, precision int) string {
	ac := accounting.Accounting{Symbol: "$", Precision: precision}
	return ac.FormatMoney(value)
}

func ObjectToJsonString(obj interface{}) string {
	json, err := json.Marshal(obj)

	if err != nil || obj == nil || json == nil || string(json) == "null" {
		return ""
	}

	return string(json)
}

func Capitalize(str string) string {
	runes := []rune(str)
  runes[0] = unicode.ToUpper(runes[0])
  return string(runes)
}
