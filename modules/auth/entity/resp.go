package entity

type LoginResponse struct {
	Id          int    `json:"-"`
	Uid         string `json:"id"`
	FullName    string `json:"fullname"`
	AccessToken string `json:"access_token"`
}

type RegisterReponse struct {
	Id          int    `json:"-"`
	Uid         string `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"fullname"`
	AccessToken string `json:"access_token"`
}
