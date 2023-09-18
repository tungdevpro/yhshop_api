package business

import (
	"coffee_api/modules/auth"
)

type business struct {
	repository auth.Repository
}

func NewBusiness(r auth.Repository) auth.Business {
	return &business{
		repository: r,
	}
}
