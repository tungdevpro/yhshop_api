package repoimpl

import (
	"coffee_api/commons"
	jwtexplore "coffee_api/middleware/jwt_explore"
	"coffee_api/modules/user/entity"
	"context"
)

type jwtexploreRepoImpl struct {
	appCtx commons.AppContext
}

func NewJwtExploreRepoImpl(appCtx commons.AppContext) jwtexplore.Repository {
	return &jwtexploreRepoImpl{
		appCtx: appCtx,
	}
}

func (impl *jwtexploreRepoImpl) FindUser(ctx context.Context, id int) (*entity.User, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	db := impl.appCtx.GetDB()
	db.Begin()

	simpler := entity.User{
		SQLModel: &commons.SQLModel{
			Id: id,
		},
	}

	if err := db.Where("id = ?", id).Find(&simpler).Error; err != nil {
		db.Rollback()
		return nil, err
	}

	db.Commit()
	return &simpler, nil
}
