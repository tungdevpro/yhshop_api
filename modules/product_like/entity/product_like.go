package entity

type ProductLike struct{}

func (ProductLike) TableName() string {
	return "product_likes"
}
