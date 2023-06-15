package owl

import "reflect"

func getValue(v any) reflect.Value {
	var value reflect.Value

	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		value = reflect.ValueOf(v)
	} else {
		value = reflect.Indirect(reflect.ValueOf(v))
	}

	return value
}
