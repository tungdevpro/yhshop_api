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

	bizUser "coffee_api/modules/user/business"
	implUser "coffee_api/modules/user/repository/repo_impl"
	restUser "coffee_api/modules/user/transport/rest"

	bizShop "coffee_api/modules/shop/business"
	implShop "coffee_api/modules/shop/repository/repo_impl"
	restShop "coffee_api/modules/shop/transport/rest"

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
	apiUpload := restUpload.NewApi(bizUpload.NewBusiness(implUpload.NewUploadRepoImpl(*appCtx)))
	apiAuth := restAuth.NewApi(bizAuth.NewBusiness(implAuth.NewAuthRepoImpl(*appCtx)))
	apiUser := restUser.NewApi(bizUser.NewBusiness(implUser.NewUserRepoImpl(*appCtx)))
	apiShop := restShop.NewApi(bizShop.NewBusiness(implShop.NewShopRepoImpl(*appCtx)))

	engine := gin.Default()

	v1 := engine.Group(prefix.V1)
	{
		// Upload handler
		v1.POST(prefix.Upload, apiUpload.UploadFile())

		// Authentication handler
		auth := v1.Group(prefix.Auth)
		{
			auth.POST(prefix.Register, apiAuth.RegisterHandler())
			auth.POST(prefix.Login, apiAuth.LoginHandler())
		}

		// User handler
		user := v1.Group(prefix.User)
		{
			user.GET(prefix.Profile, apiUser.GetProfileHandler())
			user.PUT(prefix.Empty, apiUser.UpdateProfileHandler())
			user.DELETE(prefix.Empty, apiUser.DeleteUserHandler())
		}

		// Shop handler
		shop := v1.Group(prefix.Shop)
		{
			shop.POST(prefix.Empty, apiShop.CreateShopHandler())
			shop.GET(prefix.GetShop, apiShop.GetShopHandler())
			shop.PUT(prefix.GetShop, apiShop.UpdateShopHandler())
			shop.DELETE(prefix.DelShop, apiShop.DeleteShopHandler())
		}
		v1.GET(prefix.ListShop, apiShop.ListShopHandler())

	}

	if err := engine.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		helpers.Fatal(err)
	}
}
