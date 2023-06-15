package strings

import (
	"errors"
	_url "net/url"
)

func url(v any, _ string) (any, error) {
	value, valid := v.(string)

	if valid == false {
		return value, errors.New("must be a string")
	}

	_, err := _url.ParseRequestURI(value)

	if err != nil {
		return value, errors.New("must be a url")
	}

	return value, nil
}
