package entity

type RegisterRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}


type LoginRequest struct {
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}


