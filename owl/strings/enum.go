package strings

import (
	"errors"
	"fmt"
	str "strings"
)

func enum(v any, tag string) (any, error) {
	value := fmt.Sprintf("%v", v)
	options := str.Split(tag[1:len(tag)-1], "|")
	found := false

	for _, option := range options {
		if value == str.TrimSpace(option) {
			found = true
			break
		}
	}

	if !found {
		return value, errors.New(fmt.Sprintf("must be one of enum %s", tag))
	}

	return value, nil
}
