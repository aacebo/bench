package tests

import (
	"bench/owl"
	"net/http"

	"github.com/go-chi/render"
)

type Create struct {
	Input  *string `json:"input" owl:"required"`
	Output *string `json:"output" owl:"required"`
}

func NewCreate(r *http.Request) (*Create, owl.Errors) {
	self := Create{}
	err := render.DecodeJSON(r.Body, &self)

	if err != nil {
		return nil, owl.NewErrors().Push(err.Error(), nil)
	}

	errs := owl.Validate(self)
	return &self, errs
}
