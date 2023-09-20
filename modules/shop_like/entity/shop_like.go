package entity

import (
	"coffee_api/commons"
	"time"
)

type ShopLike struct {
	ProductId int                 `json:"shop_id" gorm:"column:shop_id;"`
	UserId    int                 `json:"-" gorm:"column:user_id;"`
	CreatedAt *time.Time          `json:"created_at" gorm:"column:created_at;"`
	User      *commons.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (ShopLike) TableName() string { return "shop_likes" }
