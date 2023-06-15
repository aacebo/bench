package owl

import (
	"fmt"
	"reflect"
	"strings"
)

func Validate(v any) Errors {
	return validate([]string{}, v)
}

func validate(path []string, v any) Errors {
	errors := NewErrors()
	value := getValue(v)

	if !isComplex(v) {
		return errors.Push("must be type Struct or Array", path)
	}

	if value.Kind() == reflect.Struct {
		errors = append(errors, validateStruct(path, value)...)
	} else if value.Kind() == reflect.Map {
		errors = append(errors, validateMap(path, value)...)
	} else {
		errors = append(errors, validateSlice(path, value)...)
	}

	return errors
}

func isComplex(v any) bool {
	value := getValue(v)
	kind := value.Kind()

	return kind == reflect.Struct ||
		kind == reflect.Array ||
		kind == reflect.Slice ||
		kind == reflect.Map
}

func validateStruct(path []string, v reflect.Value) Errors {
	t := v.Type()
	errors := NewErrors()

	for i := 0; i < v.NumField(); i++ {
		errs := NewErrors()

		if v.Field(i).CanInterface() {
			if isComplex(v.Field(i).Interface()) {
				errs = validate(append(path, t.Field(i).Tag.Get("json")), v.Field(i).Interface())
			} else {
				errs = validateField(path, t.Field(i), v.Field(i))
			}
		}

		errors = append(errors, errs...)
	}

	return errors
}

func validateMap(path []string, v reflect.Value) Errors {
	errors := NewErrors()

	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		errs := NewErrors()

		if value.CanInterface() {
			if isComplex(value.Interface()) {
				errs = validate(append(path, key.String()), value.Interface())
			}
		}

		errors = append(errors, errs...)
	}

	return errors
}

func validateSlice(path []string, v reflect.Value) Errors {
	errors := NewErrors()

	for i := 0; i < v.Len(); i++ {
		errs := NewErrors()

		if isComplex(v.Index(i).Interface()) {
			errs = validate(append(path, fmt.Sprint(i)), v.Index(i).Interface())
		}

		errors = append(errors, errs...)
	}

	return errors
}

func validateField(path []string, f reflect.StructField, v reflect.Value) Errors {
	kind := v.Kind()

	if kind == reflect.Pointer && !v.IsNil() {
		v = getValue(v.Interface())
		kind = v.Kind()
	}

	errors := NewErrors()
	tag := f.Tag.Get("owl")

	if tag != "" {
		rules := rules[kind]

		if rules != nil {
			for _, token := range strings.Split(tag, ",") {
				rule := strings.Split(token, "=")
				key := strings.TrimSpace(rule[0])
				value := ""

				if len(rule) == 2 {
					value = strings.TrimSpace(rule[1])
				}

				fn := rules[key]

				if fn != nil {
					_, err := fn(v.Interface(), value)

					if err != nil {
						errors = errors.Push(
							err.Error(),
							append(path, f.Tag.Get("json")),
						)
					}
				}
			}
		}
	}

	return errors
}
