package sessions

import (
	"bench/owl"
	"bench/utils"
	"net/http"

	"github.com/go-chi/render"
)

type Create struct {
	Email    *string `json:"email" owl:"required, email, max=255"`
	Password *string `json:"password" owl:"required"`
}

func NewCreate(r *http.Request) (*Create, owl.Errors) {
	self := Create{}
	err := render.DecodeJSON(r.Body, &self)

	if err != nil {
		return nil, owl.NewErrors().Push(err.Error(), nil)
	}

	errs := owl.Validate(self)

	if errs.Len() == 0 {
		password := utils.Sha256(*self.Password)
		self.Password = &password
	}

	return &self, errs
}
