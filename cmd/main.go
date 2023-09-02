package main

import (
	"coffee_api/commons"
	"coffee_api/configs"
	"coffee_api/db"
	"coffee_api/helpers"
	bizAuth "coffee_api/modules/auth/business"
	implAuth "coffee_api/modules/auth/repository/repo_impl"
	restAuth "coffee_api/modules/auth/transport/rest"
	"coffee_api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.NewConfiguration()
	db, err := db.CreateMysqlDB(cfg)
	if err != nil {
		helpers.Fatal(err)
	}

	appCtx := commons.NewAppContext(db)
	fmt.Println("db create: ", appCtx.GetDB())

	apiAuth := restAuth.NewApi(bizAuth.NewBusiness(implAuth.NewAuthRepoImpl(*appCtx)))
	fmt.Println("apiAuth:: ", apiAuth)

	engine := gin.Default()

	v1 := engine.Group("/v1")
	{
		routes.AuthRoutes()
		routes.UserRoutes(v1)
	}

	if err := engine.Run(":3000"); err != nil {
		helpers.Fatal(err)
	}
}
