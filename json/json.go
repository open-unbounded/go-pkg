package json

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

// Unmarshal parses the JSON-encoded data.
func Unmarshal[T any](data []byte) (*T, error) {
	out := new(T)
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

// UnmarshalString parses the JSON-encoded data.
func UnmarshalString[T any](data string) (*T, error) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&data))
	var b []byte
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Len = sh.Len
	bh.Data = sh.Data
	bh.Cap = sh.Len

	return Unmarshal[T](b)
}
