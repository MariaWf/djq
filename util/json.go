package util

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	//stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-01"))
	return []byte(stamp), nil
}

func StringTime4DB(source time.Time) string {
	return source.Format("2006-01-02 15:04:05")
}

func StringDefaultTime4DB() string {
	//return "1970-01-01 00:00:00"
	return "1970-01-01 08:00:01"
	//return "0000-00-00 00:00:00"
}

func ParseTimeFromDB(source string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", source)
}
