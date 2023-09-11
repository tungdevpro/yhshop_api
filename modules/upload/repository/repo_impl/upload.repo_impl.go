package repoimpl

import (
	"coffee_api/commons"
	"coffee_api/modules/upload"
	"coffee_api/modules/upload/entity"
	"coffee_api/modules/upload/uploadprovider"
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

func (impl *uploadRepoImpl) UploadFile(ctx context.Context, uploadDto *entity.UploadDTO) (*commons.Image, error) {
	cfg := impl.appCtx.Cfg
	s3 := uploadprovider.NewS3Provider(cfg.S3BucketName, cfg.S3Region, cfg.S3ApiKey, cfg.S3SecretKey, cfg.S3Domain)

	img, err := s3.SaveFileUploaded(ctx, uploadDto)
	if err != nil {
		return nil, err
	}
	return img, nil
}
