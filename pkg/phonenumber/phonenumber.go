package phonenumber

import (
	"q-game-app/pkg/errmsg"
	"q-game-app/pkg/richerror"
	"strconv"
)

func IsValid(phoneNumber string) (bool, error) {
	const op = "service.ValidatePhoneNumber"
	if len(phoneNumber) != 11 {
		return false, richerror.New(op).WithMessage(errmsg.ErrorMsgPasswordLetterTwelve)
	}
	if phoneNumber[0:2] != "09" {
		return false, richerror.New(op).WithMessage(errmsg.ErrorMsgStartedWithZeroNine)
	}
	if _, err := strconv.Atoi(phoneNumber[2:]); err != nil {
		return false, richerror.New(op).WithMessage(errmsg.ErrorMsgMustInt)
	}
	return true, nil
}
