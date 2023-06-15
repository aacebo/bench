package strings

import (
	"errors"
	"fmt"
	"strconv"
)

func min(v any, tag string) (any, error) {
	value, valid := v.(string)

	if valid == false {
		return value, errors.New("must be a string")
	}

	min, err := strconv.Atoi(tag)

	if err != nil {
		return value, err
	}

	if len(value) < min {
		return value, errors.New(fmt.Sprintf("length must be at least %d", min))
	}

	return value, nil
}
