package repoimpl

import (
	"coffee_api/commons"
	jwtexplore "coffee_api/middleware/jwt_explore"
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

func (impl *jwtexploreRepoImpl) FindUser(ctx context.Context, id int) (*commons.SimpleUser, error) {
	impl.appCtx.L.Lock()
	defer impl.appCtx.L.Unlock()

	db := impl.appCtx.GetDB()

	simpler := commons.SimpleUser{}
	simpler.Id = id

	db.Begin()
	if err := db.Where("id = ?", id).Find(&simpler).Error; err != nil {
		db.Rollback()
		return nil, err
	}

	db.Commit()
	return &simpler, nil
}
