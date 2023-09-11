package business

import (
	"coffee_api/modules/upload"
	"coffee_api/modules/upload/entity"
	"context"
)

type business struct {
	repository upload.Repository
}

func NewBusiness(repository upload.Repository) upload.Business {
	return &business{repository: repository}
}

func (biz *business) UploadFile(ctx context.Context, uploadDto *entity.UploadDTO) error {
	return nil
}
