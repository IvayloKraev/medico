package dto

import "errors"

type AdminLogin struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

func (a AdminLogin) Validate() error {
	return errors.Join(a.Email.Validate(), a.Password.Validate())
}
