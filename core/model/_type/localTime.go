package _type

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	if tTime.IsZero() {
		return []byte(fmt.Sprintf("\"%v\"", "")), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(DefaultTimeFormat))), nil
}
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
