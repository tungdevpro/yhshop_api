package entity

import (
	"strconv"
	"strings"
)

const MAX_OTP = 5

type OTPRequest struct {
	Email string `json:"email" form:"email"`
	Otp   int    `json:"otp" form:"otp"`
}

func (req *OTPRequest) Validate() error {
	req.Email = strings.TrimSpace(req.Email)
	if err := isEmailAddress(req.Email); err != nil {
		return err
	}

	otp := len(strconv.Itoa(req.Otp))
	if otp != MAX_OTP {
		return ErrOTPLength
	}

	return nil
}
