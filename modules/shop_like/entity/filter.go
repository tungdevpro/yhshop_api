package entity

type Filter struct {
	ShopId int `json:"shop_id" form:"shop_id"`
	UserId int `json:"user_id" form:"user_id"`
}
