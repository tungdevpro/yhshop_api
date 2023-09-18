package entity

import "time"

type ShopLike struct {
	ProductId int        `json:"shop_id" gorm:"column:shop_id"`
	UserId    int        `json:"user_id" gorm:"column:user_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (ShopLike) TableName() string { return "shop_likes" }
