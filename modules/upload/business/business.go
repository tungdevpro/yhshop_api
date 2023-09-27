package business

import (
	"bytes"
	"coffee_api/commons"
	"coffee_api/helpers"
	"coffee_api/modules/upload"
	"coffee_api/modules/upload/entity"
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"path/filepath"
	"strings"
	"time"
)

type business struct {
	repository upload.Repository
}

func NewBusiness(repository upload.Repository) upload.Business {
	return &business{repository: repository}
}

func (biz *business) UploadFile(ctx context.Context, uploadDto *entity.UploadDTO) (*commons.Image, error) {
	doc := uploadDto
	doc.Dst = fmt.Sprintf("%s/%s", uploadDto.Folder, uploadDto.FileName)
	fileBytes := bytes.NewBuffer(doc.Data)

	width, height, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(doc.Folder) == "" {
		doc.Folder = "img"
	}

	fileExt := filepath.Ext(doc.FileName)
	doc.Dst = fmt.Sprintf("%s/%d-%s%s", doc.Folder, time.Now().UnixNano(), helpers.SnakeCase(doc.FileName), fileExt)

	img, err := biz.repository.UploadFile(ctx, doc)
	img.Width = width
	img.Height = height
	img.CloudName = "s3"
	img.Extension = fileExt

	if err != nil {
		return nil, err
	}
	return img, nil
}

// Handle get dimension of image
func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
