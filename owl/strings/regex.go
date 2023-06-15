package strings

import (
	"errors"
	"fmt"
	"regexp"
)

func regex(v any, tag string) (any, error) {
	value, valid := v.(string)

	if valid == false {
		return value, errors.New("must be a string")
	}

	valid, err := regexp.MatchString(tag, value)

	if err != nil {
		return value, err
	}

	if valid == false {
		return value, errors.New(fmt.Sprintf("must match pattern \"%s\"", tag))
	}

	return value, nil
}
