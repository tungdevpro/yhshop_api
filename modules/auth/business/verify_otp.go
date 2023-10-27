package business

import (
	"coffee_api/modules/auth/entity"
	"context"
)

func (biz *business) VerifyOTP(ctx context.Context, param *entity.OTPRequest) (bool, error) {
	if err := param.Validate(); err != nil {
		return false, err
	}

	isVerify, err := biz.repository.VerifyOTP(ctx, param)
	if err != nil {
		return false, err
	}

	return isVerify, nil
}
