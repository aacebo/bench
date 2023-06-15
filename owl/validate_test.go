package owl

import "testing"

type A struct {
	AString string `owl:"min=5"`
	AInt    int
	Nested  []*B
}

type B struct {
	BString string `owl:"eq=10, gt=15"`
}

func TestValidateStruct(t *testing.T) {
	if len(Validate(A{"tttt", 0, []*B{}})) != 1 {
		t.Error("should have error")
	}
}

func TestValidateSlice(t *testing.T) {
	nested := B{"test"}
	Validate(A{"", 25, []*B{&nested, &nested}})
}
