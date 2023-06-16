package response

import (
	"bench/owl"
	"net/http"

	"github.com/go-chi/render"
)

type Response[T any] struct {
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Errors owl.Errors             `json:"errors,omitempty"`
	Data   T                      `json:"data"`

	status *int
}

func New[T any](data T) *Response[T] {
	self := Response[T]{
		Meta:   map[string]interface{}{},
		Errors: owl.Errors{},
		Data:   data,
	}

	return &self
}

func NotFound(path ...string) *Response[any] {
	self := New[interface{}](nil)
	return self.Status(404).AddError("not found", path...)
}

func Unauthorized() *Response[any] {
	self := New[interface{}](nil)
	return self.Status(401).AddError("unauthorized")
}

func BadReqest(errs owl.Errors) *Response[any] {
	self := New[interface{}](nil)
	self = self.Status(400)
	self.Errors = errs
	return self
}

func Conflict(path ...string) *Response[any] {
	self := New[interface{}](nil)
	return self.Status(409).AddError("conflict", path...)
}

func Internal(message string) *Response[any] {
	self := New[interface{}](nil)
	return self.Status(500).AddError(message)
}

func (self *Response[T]) Status(code int) *Response[T] {
	self.status = &code
	return self
}

func (self *Response[T]) SetMeta(key string, value any) *Response[T] {
	self.Meta[key] = value
	return self
}

func (self *Response[T]) AddError(message string, path ...string) *Response[T] {
	self.Errors = self.Errors.Push(message, path)
	return self
}

func (self *Response[T]) Do(w http.ResponseWriter, r *http.Request) {
	status := 200

	if self.status != nil {
		status = *self.status
	}

	render.Status(r, status)
	render.JSON(w, r, self)
}
