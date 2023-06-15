package owl

type Error struct {
	Message string   `json:"message"`
	Path    []string `json:"path,omitempty"`
}

type Errors []Error

func NewErrors() Errors {
	self := Errors{}
	return self
}

func (self Errors) Push(message string, path []string) Errors {
	self = append(self, Error{
		message,
		path,
	})

	return self
}

func (self Errors) Len() int {
	return len(self)
}
