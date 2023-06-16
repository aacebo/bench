package problems

import (
	"bench/owl"
	"net/http"

	"github.com/go-chi/render"
)

type Create struct {
	Name        *string `json:"name" owl:"required, max=50, regex=^[0-9a-zA-Z_-]+$"`
	DisplayName *string `json:"display_name" owl:"required, max=50"`
	Description *string `json:"description" owl:"required, max=500"`
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
