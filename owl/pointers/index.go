package pointers

func New() map[string]func(any, string) (any, error) {
	return map[string]func(any, string) (any, error){
		"required": required,
	}
}
