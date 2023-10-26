package entity

import "time"

type VerifyMail struct {
	Id         int        `json:"-" gorm:"column:id;"`
	FullName   string     `json:"fullname" gorm:"column:fullname;"`
	Email      string     `json:"email" gorm:"column:email;type:varchar(100);"`
	SecretCode string     `json:"secret_code" gorm:"column:secret_code;"`
	IsUsed     bool       `json:"is_used" gorm:"column:is_used;"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;"`
	ExpiredAt  *time.Time `json:"expired_at" gorm:"column:expired_at;"`
}

func (VerifyMail) TableName() string { return "verify_mails" }
