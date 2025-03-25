package auth

import (
	"encoding/json"
	"go.uber.org/multierr"
	"regexp"
)

type Validator interface {
	Validate() error
}

type RequestAdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	_ json.Unmarshaler = (*RequestAdminLogin)(nil)
	_ Validator        = (*RequestAdminLogin)(nil)
)

func (a RequestAdminLogin) UnmarshalJSON(bytes []byte) (err error) {
	type AliasRequestAdminLogin RequestAdminLogin
	temp := new(AliasRequestAdminLogin)
	err = json.Unmarshal(bytes, &temp)

	if err != nil {
		return ErrWrongData
	}

	a.Email = temp.Email
	a.Password = temp.Password

	err = a.Validate()
	if err != nil {
		a.Email = ""
		a.Password = ""
		return err
	}

	return nil
}

func (a RequestAdminLogin) Validate() (err error) {
	multierr.AppendFunc(&err, func() error {
		if !regexp.MustCompile(emailPattern).MatchString(a.Email) {
			return ErrEmailIncorrect
		}
		return nil
	})

	return err
}
