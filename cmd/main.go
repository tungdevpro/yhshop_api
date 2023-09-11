package main

import (
	"coffee_api/commons"
	"coffee_api/configs"
	"coffee_api/configs/prefix"
	"coffee_api/db"
	"coffee_api/helpers"
	bizAuth "coffee_api/modules/auth/business"
	implAuth "coffee_api/modules/auth/repository/repo_impl"
	restAuth "coffee_api/modules/auth/transport/rest"

	bizUpload "coffee_api/modules/upload/business"
	implUpload "coffee_api/modules/upload/repository/repo_impl"

	restUpload "coffee_api/modules/upload/transport/rest"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.NewConfiguration()
	db, err := db.CreateMysqlDB(cfg)
	if err != nil {
		helpers.Fatal(err)
	}

	appCtx := commons.NewAppContext(db, cfg)
	apiAuth := restAuth.NewApi(bizAuth.NewBusiness(implAuth.NewAuthRepoImpl(*appCtx)))
	apiUpload := restUpload.NewApi(bizUpload.NewBusiness(implUpload.NewUploadRepoImpl(*appCtx)))

	engine := gin.Default()

	v1 := engine.Group(prefix.V1)
	{
		v1.POST(prefix.Upload, apiUpload.UploadFile())

		auth := v1.Group(prefix.Auth)
		auth.POST(prefix.Register, apiAuth.RegisterHandler())
		auth.POST(prefix.Login, apiAuth.LoginHandler())
	}

	if err := engine.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		helpers.Fatal(err)
	}
}
