package jwtexplore

import (
	"coffee_api/commons"
	"context"
)

type Repository interface {
	FindUser(context.Context, int) (*commons.SimpleUser, error)
}

type Business interface {
	FindUser(context.Context, int) (*commons.SimpleUser, error)
}
