package main

import (
	"fmt"

	"coffee_api/commons"
	"coffee_api/configs"
	__prefix "coffee_api/configs/prefix"
	"coffee_api/consumers"
	"coffee_api/db"
	"coffee_api/helpers"
	"coffee_api/middleware"
	bizJwt "coffee_api/middleware/jwt_explore/business"
	implJwt "coffee_api/middleware/jwt_explore/repository/repo_impl"

	bizAuth "coffee_api/modules/auth/business"
	implAuth "coffee_api/modules/auth/repository/repo_impl"
	restAuth "coffee_api/modules/auth/transport/rest"

	bizUser "coffee_api/modules/user/business"
	implUser "coffee_api/modules/user/repository/repo_impl"
	restUser "coffee_api/modules/user/transport/rest"

	bizShop "coffee_api/modules/shop/business"
	implShop "coffee_api/modules/shop/repository/repo_impl"
	restShop "coffee_api/modules/shop/transport/rest"

	bizShopLike "coffee_api/modules/shop_like/business"
	implShopLike "coffee_api/modules/shop_like/repository/repo_impl"
	restShopLike "coffee_api/modules/shop_like/transport/rest"

	bizUpload "coffee_api/modules/upload/business"
	implUpload "coffee_api/modules/upload/repository/repo_impl"
	restUpload "coffee_api/modules/upload/transport/rest"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func main() {
	cfg := configs.NewConfiguration()
	db, err := db.CreateMysqlDB(cfg)
	if err != nil {
		helpers.Fatal(err)
	}

	appCtx := commons.NewAppContext(db, cfg)
	consumers.NewCusumnerEngine(*appCtx).Start()

	jwtHandler := bizJwt.NewBusiness(implJwt.NewJwtExploreRepoImpl(*appCtx))
	apiUpload := restUpload.NewApi(bizUpload.NewBusiness(implUpload.NewUploadRepoImpl(*appCtx)))

	apiAuth := restAuth.NewApi(bizAuth.NewBusiness(implAuth.NewAuthRepoImpl(*appCtx)))
	apiUser := restUser.NewApi(bizUser.NewBusiness(implUser.NewUserRepoImpl(*appCtx)))

	bizShop := bizShop.NewBusiness(implShop.NewShopRepoImpl(*appCtx))
	__implShopLike := implShopLike.NewShopLikeRepoImpl(*appCtx)
	apiShopLike := restShopLike.NewApi(bizShopLike.NewBusiness(__implShopLike, bizShop))
	apiShop := restShop.NewApi(bizShop)

	engine := gin.Default()
	engine.Use(middleware.AuthRequired(*appCtx, jwtHandler))

	v1 := engine.Group(__prefix.V1)
	{
		// Upload handler
		v1.POST(__prefix.Upload, apiUpload.UploadFile())

		// Authentication handler
		auth := v1.Group(__prefix.Auth)
		{
			auth.POST(__prefix.Register, apiAuth.RegisterHandler())
			auth.POST(__prefix.Login, apiAuth.LoginHandler())
			auth.POST(__prefix.VerifyOTP, apiAuth.VerifyOTPHandler())
		}

		// User handler
		user := v1.Group(__prefix.User)
		{
			user.GET(__prefix.Profile, apiUser.GetProfileHandler())
			user.PUT(__prefix.Empty, apiUser.UpdateProfileHandler())
			user.DELETE(__prefix.Empty, apiUser.DeleteUserHandler())
		}

		// Shop handler
		shop := v1.Group(__prefix.Shop)
		{
			shop.POST(__prefix.Empty, apiShop.CreateShopHandler())
			shop.GET(__prefix.GetShop, apiShop.GetShopHandler())
			shop.PUT(__prefix.GetShop, apiShop.UpdateShopHandler())
			shop.DELETE(__prefix.DelShop, apiShop.DeleteShopHandler())
			shop.GET(__prefix.LikedUsers, apiShopLike.GetLikedUsersHandler())
			shop.POST(__prefix.CreateUserLike, apiShopLike.CreateUserLikeHandler())
			shop.DELETE(__prefix.DeleteUserLike, apiShopLike.DeleteUserLikeHandler())

		}
		v1.GET(__prefix.ListShop, apiShop.ListShopHandler())
	}

	// startSocketIOServer(engine)
	if err := engine.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		helpers.Fatal(err)
	}
}

func startSocketIOServer(engine *gin.Engine) {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	server.OnConnect("/", func(c socketio.Conn) error {
		c.SetContext("")
		fmt.Println("connected: ", c.ID(), "IP: ", c.RemoteAddr())
		return nil
	})

	go server.Serve()
}
