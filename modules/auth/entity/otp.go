package entity

type OTPRequest struct {
	Otp int `json:"otp" form:"otp"`
}
