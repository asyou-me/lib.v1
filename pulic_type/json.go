package pulic_type

import (
	"database/sql/driver"
)

type StringNull struct {
	Data string
}

func (t *StringNull) MarshalJSON() ([]byte, error) {
	if t.Data == "" {
		return []byte("null"), nil
	} else {
		return []byte(`"` + t.Data + `"`), nil
	}
}

func (t *StringNull) Value() (driver.Value, error) {
	return driver.Value(t.Data), nil
}
