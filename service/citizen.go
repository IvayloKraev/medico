package service

import "errors"

type CitizenService interface {
	AuthenticateByEmailAndPassword() error
	FindAllAvailablePharmacies() error
	ListPrescriptions() error
}

type citizenService struct {
}

func NewCitizenService() CitizenService {
	return &citizenService{}
}

func (c *citizenService) AuthenticateByEmailAndPassword() error {
	return errors.New("not implemented")
}
func (c *citizenService) FindAllAvailablePharmacies() error {
	return errors.New("not implemented")
}

func (c *citizenService) ListPrescriptions() error {
	return errors.New("not implemented")
}
