package owl

import (
	"reflect"

	"bench/owl/pointers"
	"bench/owl/strings"
)

var rules = map[reflect.Kind]map[string]func(any, string) (any, error){
	reflect.String:  strings.New(),
	reflect.Pointer: pointers.New(),
}
