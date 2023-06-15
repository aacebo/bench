package pointers

import (
	"errors"
	"reflect"
)

func required(v any, _ string) (any, error) {
	if reflect.ValueOf(v).IsNil() {
		return v, errors.New("required")
	}

	return v, nil
}
