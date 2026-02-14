package uservalidator

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"q-game-app/param"
	"q-game-app/pkg/errmsg"
	"q-game-app/pkg/richerror"
	"regexp"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) error {
	const op = "service.ValidateRegisterRequest"
	err := validation.ValidateStruct(&req,
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(3, 50),
		),
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
			validation.By(v.checkPhoneNumberUnique),
		),
		validation.Field(&req.Password,
			validation.Required,
			validation.Length(8, 0),
			validation.By(func(value interface{}) error {
				s, _ := value.(string)
				if !regexp.MustCompile(`[a-z]`).MatchString(s) {
					return fmt.Errorf("password must contain a lowercase letter")
				}
				if !regexp.MustCompile(`[A-Z]`).MatchString(s) {
					return fmt.Errorf("password must contain an uppercase letter")
				}
				if !regexp.MustCompile(`\d`).MatchString(s) {
					return fmt.Errorf("password must contain a digit")
				}
				return nil
			}),
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

func (v Validator) checkPhoneNumberUnique(value interface{}) error {
	var phoneNumber = value.(string)
	if isUnique, err := v.repo.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}
		return fmt.Errorf(errmsg.ErrorMsgPhoneNumberIsNotUnique)
	}
	return nil
}
