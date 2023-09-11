package uploadprovider

import (
	"coffee_api/commons"
	"coffee_api/modules/upload/entity"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(context.Context, *entity.UploadDTO) (*commons.Image, error)
}
