package business

import "context"

type Business interface {
	Register(context.Context) error
	Login(context.Context) error
}
