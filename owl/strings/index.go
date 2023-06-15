package strings

func New() map[string]func(any, string) (any, error) {
	return map[string]func(any, string) (any, error){
		"min":      min,
		"max":      max,
		"regex":    regex,
		"url":      url,
		"enum":     enum,
		"email":    email,
		"required": required,
	}
}
