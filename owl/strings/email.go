package strings

import (
	"errors"
	"fmt"
	"net/mail"
)

func email(v any, tag string) (any, error) {
	value := fmt.Sprintf("%v", v)
	_, err := mail.ParseAddress(value)

	if err != nil {
		return value, errors.New("must be a valid email address")
	}

	return value, nil
}
