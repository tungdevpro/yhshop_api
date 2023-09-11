package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/upload"
	"coffee_api/modules/upload/entity"
	"context"
)

type uploadRepoImpl struct {
	appCtx commons.AppContext
}

func NewUploadRepoImpl(appCtx commons.AppContext) upload.Repository {
	return &uploadRepoImpl{
		appCtx: appCtx,
	}
}

func (u *uploadRepoImpl) UploadFile(ctx context.Context, uploadDto *entity.UploadDTO) error {
	return nil
}
