package entity

type OTPUser struct {
	Email string `json:"email"`
	Otp   string `json:"json"`
}

func (OTPUser) TableName() string { return User{}.TableName() }
