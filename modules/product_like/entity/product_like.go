package entity

import "time"

type ProductLike struct {
	ProductId int        `json:"product_id" gorm:"column:product_id"`
	UserId    int        `json:"user_id" gorm:"column:user_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (ProductLike) TableName() string { return "product_likes" }
