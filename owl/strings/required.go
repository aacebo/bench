package strings

import (
	"errors"
	"fmt"
	str "strings"
)

func required(v any, _ string) (any, error) {
	value := fmt.Sprintf("%v", v)

	if str.TrimSpace(value) == "" {
		return value, errors.New("required")
	}

	return value, nil
}
