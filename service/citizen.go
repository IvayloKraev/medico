package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"medico/dto"
	"medico/models"
	"medico/repo"
	"medico/session"
	"time"
)

type CitizenService interface {
	AuthenticateByEmailAndPassword(email string, password string, citizenAuth *models.CitizenAuth) error
	CreateAuthenticationSession(citizenId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error

	GetMedicalInfo(citizenId uuid.UUID, medicalInfo *dto.ResponseCitizenMedicalInfo) error
	GetPersonalDoctor(citizenId uuid.UUID, doctor *dto.ResponseCitizenPersonalDoctor) error
	FindAllAvailablePharmacies(prescriptionId *dto.QueryCitizenAvailablePharmacyGet, availablePharmacies *[]dto.ResponseCitizenAvailablePharmacy) error
	ListPrescriptions(citizenId uuid.UUID, prescriptionsDto *[]dto.ResponseCitizenPrescription) error
}

type citizenService struct {
	authSession session.AuthSession
	citizenRepo repo.CitizenRepo
}

func NewCitizenService() CitizenService {
	return &citizenService{
		authSession: session.NewAuthSession("citizen"),
		citizenRepo: repo.NewCitizenRepo(),
	}
}

func (c *citizenService) AuthenticateByEmailAndPassword(email string, password string, citizenAuth *models.CitizenAuth) error {

	if err := c.citizenRepo.FindAuthByEmail(email, citizenAuth); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(citizenAuth.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (c *citizenService) CreateAuthenticationSession(citizenId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return c.authSession.CreateAuthSession(citizenId)
}

func (c *citizenService) GetAuthenticationSession(sessionId uuid.UUID) (uuid.UUID, error) {
	return c.authSession.GetAuthSession(sessionId)
}

func (c *citizenService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return c.authSession.DeleteAuthSession(sessionID)
}

func (c *citizenService) GetMedicalInfo(citizenId uuid.UUID, medicalInfo *dto.ResponseCitizenMedicalInfo) error {
	citizen := models.Citizen{}

	if err := c.citizenRepo.FindMedicalInfo(citizenId, &citizen); err != nil {
		return err
	}

	medicalInfo = &dto.ResponseCitizenMedicalInfo{
		FirstName:  citizen.FirstName,
		SecondName: citizen.SecondName,
		LastName:   citizen.LastName,
		BirthDate:  citizen.Birthday,
		Sex:        string(citizen.Sex),
		UCN:        citizen.UCN,
	}

	return nil
}

func (c *citizenService) GetPersonalDoctor(citizenId uuid.UUID, doctorDto *dto.ResponseCitizenPersonalDoctor) error {
	doctor := models.Doctor{}

	if err := c.citizenRepo.FindPersonalDoctor(citizenId, &doctor); err != nil {
		return err
	}

	doctorDto = &dto.ResponseCitizenPersonalDoctor{
		FirstName:  doctor.FirstName,
		SecondName: doctor.SecondName,
		LastName:   doctor.LastName,
		UIN:        doctor.UIN,
		Email:      doctor.Email,
	}

	return nil
}

func (c *citizenService) FindAllAvailablePharmacies(prescriptionId *dto.QueryCitizenAvailablePharmacyGet, availablePharmacies *[]dto.ResponseCitizenAvailablePharmacy) error {
	branches := new([]models.PharmacyBranch)

	if err := c.citizenRepo.FindAvailablePharmacies(prescriptionId.PrescriptionId, branches); err != nil {
		return err
	}

	*availablePharmacies = make([]dto.ResponseCitizenAvailablePharmacy, len(*branches))

	for i, branch := range *branches {
		(*availablePharmacies)[i] = dto.ResponseCitizenAvailablePharmacy{
			Name:      branch.Name,
			Latitude:  branch.Latitude,
			Longitude: branch.Longitude,
		}
	}

	return nil
}

func (c *citizenService) ListPrescriptions(citizenId uuid.UUID, prescriptionsDto *[]dto.ResponseCitizenPrescription) error {
	prescriptions := new([]models.Prescription)

	if err := c.citizenRepo.FindAllPrescriptions(citizenId, prescriptions); err != nil {
		return err
	}

	*prescriptionsDto = make([]dto.ResponseCitizenPrescription, len(*prescriptions))

	for i, prescription := range *prescriptions {
		(*prescriptionsDto)[i] = dto.ResponseCitizenPrescription{
			Doctor: struct {
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				UIN       string `json:"uin"`
			}{
				FirstName: prescription.Doctor.FirstName,
				LastName:  prescription.Doctor.LastName,
				UIN:       prescription.Doctor.UIN,
			},
			Medicaments: make([]struct {
				Name string `json:"name"`
				Unit uint   `json:"unit"`
			}, 0),
		}

		for i2, medicament := range prescription.Medicaments {
			(*prescriptionsDto)[i].Medicaments[i2] = struct {
				Name string `json:"name"`
				Unit uint   `json:"unit"`
			}{
				Name: medicament.Medicament.OfficialName,
				Unit: medicament.Quantity,
			}
		}
	}

	return nil
}
