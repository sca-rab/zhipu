package zhipu

import (
	"bytes"
	"encoding/json"
)

type StringOr[T any] struct {
	String *string
	Value  *T
}

var (
	_ json.Marshaler   = StringOr[float64]{}
	_ json.Unmarshaler = &StringOr[float64]{}
)

func (f *StringOr[T]) SetString(v string) {
	f.String = &v
	f.Value = nil
}

func (f *StringOr[T]) SetValue(v T) {
	f.String = nil
	f.Value = &v
}

func (f StringOr[T]) MarshalJSON() ([]byte, error) {
	if f.Value != nil {
		return json.Marshal(f.Value)
	}
	return json.Marshal(f.String)
}

func (f *StringOr[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	if data[0] == '"' {
		f.String = new(string)
		f.Value = nil
		return json.Unmarshal(data, f.String)
	} else {
		f.Value = new(T)
		f.String = nil
		return json.Unmarshal(data, f.Value)
	}
}
