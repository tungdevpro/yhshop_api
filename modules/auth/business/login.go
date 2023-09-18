package business

import (
	"coffee_api/modules/auth/entity"
	"context"

	"github.com/indrasaputra/hashids"
)

func (biz *business) Login(ctx context.Context, req *entity.LoginDTO) (*entity.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	result, err := biz.repository.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	xId, _ := hashids.EncodeID(hashids.ID(result.Id))

	result.Uid = string(xId)

	return result, nil
}
