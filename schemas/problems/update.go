package problems

import (
	"bench/owl"
	"net/http"

	"github.com/go-chi/render"
)

type Update struct {
	Name        *string `json:"name" owl:"max=50, regex=^[0-9a-zA-Z_-]+$"`
	DisplayName *string `json:"display_name" owl:"max=50"`
	Description *string `json:"description" owl:"max=500"`
}

func NewUpdate(r *http.Request) (*Update, owl.Errors) {
	self := Update{}
	err := render.DecodeJSON(r.Body, &self)

	if err != nil {
		return nil, owl.NewErrors().Push(err.Error(), nil)
	}

	errs := owl.Validate(self)
	return &self, errs
}
