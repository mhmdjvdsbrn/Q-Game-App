package uservalidator

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"q-game-app/param"
	"q-game-app/pkg/errmsg"
	"q-game-app/pkg/richerror"
	"regexp"
)

func (v Validator) ValidateLoginRequest(req param.LoginRequest) error {
	const op = "service.ValidateLoginRequest"
	err := validation.ValidateStruct(&req,
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
			validation.By(v.doesPhoneNumberExist),
		),
		validation.Field(&req.Password,
			validation.Required,
		),
	)

	if err != nil {
		fieldErrors := make(map[string][]string)
		if errV, ok := err.(validation.Errors); ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = []string{value.Error()}
				}
			}
		}

		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{
				"errors":      fieldErrors,
				"phoneNumber": req.PhoneNumber,
			})
	}

	return nil
}

func (v Validator) doesPhoneNumberExist(value interface{}) error {
	var phoneNumber = value.(string)
	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return fmt.Errorf(errmsg.ErrorMsgNotFound)
	}
	return nil
}
