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

	engine := gin.Default()

	v1 := engine.Group(prefix.V1)
	{
		auth := v1.Group(prefix.Auth)
		auth.POST(prefix.Register, apiAuth.RegisterHandler())
		auth.POST(prefix.Login, apiAuth.LoginHandler())

	}

	if err := engine.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		helpers.Fatal(err)
	}
}
