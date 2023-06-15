package strings

import (
	"errors"
	"fmt"
	"strconv"
)

func max(v any, tag string) (any, error) {
	value, valid := v.(string)

	if valid == false {
		return value, errors.New("must be a string")
	}

	max, err := strconv.Atoi(tag)

	if err != nil {
		return value, err
	}

	if len(value) > max {
		return value, errors.New(fmt.Sprintf("length must be at most %d", max))
	}

	return value, nil
}
