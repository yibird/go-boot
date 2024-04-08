package _type

import (
	"database/sql/driver"
	"errors"
)

type BitBool bool

func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	}
	return []byte{0}, nil
}

func (b *BitBool) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte _type assertion")
	}
	*b = v[0] == 1
	return nil
}
