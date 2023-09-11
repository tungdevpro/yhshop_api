package rest

import (
	"coffee_api/commons"
	"coffee_api/modules/upload"
	"coffee_api/modules/upload/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz upload.Business
}

func NewApi(biz upload.Business) *api {
	return &api{biz: biz}
}

func (api *api) UploadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			panic(err)
		}

		folder := ctx.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(err)
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(err)
		}

		img, err := api.biz.UploadFile(ctx.Request.Context(), &entity.UploadDTO{
			Data:     dataBytes,
			Folder:   folder,
			FileName: fileHeader.Filename,
		})

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, commons.SimpleSuccessResp(img))

		// ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))

	}
}
