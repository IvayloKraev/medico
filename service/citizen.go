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

	GetMedicalInfo(citizenId uuid.UUID, medicalInfo *dto.CitizenMedicalInfo) error
	FindAllAvailablePharmacies(prescriptionId *dto.CitizenAvailablePharmacyGet, availablePharmacies *[]dto.CitizenAvailablePharmacy) error
	ListPrescriptions(citizenId uuid.UUID, prescriptionsDto *[]dto.CitizenPrescription) error
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

func (c *citizenService) GetMedicalInfo(citizenId uuid.UUID, medicalInfo *dto.CitizenMedicalInfo) error {
	citizen := models.Citizen{}

	if err := c.citizenRepo.FindMedicalInfo(citizenId, &citizen); err != nil {
		return err
	}

	medicalInfo = &dto.CitizenMedicalInfo{
		FirstName:  citizen.FirstName,
		SecondName: citizen.SecondName,
		LastName:   citizen.LastName,
		BirthDate:  citizen.Birthday,
		Sex:        string(citizen.Sex),
		UCN:        citizen.UCN,
		PersonalDoctor: struct {
			FirstName  string `json:"first_name"`
			SecondName string `json:"second_name"`
			LastName   string `json:"last_name"`
			UIN        string `json:"uin"`
			Email      string `json:"email"`
		}{
			FirstName:  citizen.PersonalDoctor.FirstName,
			SecondName: citizen.PersonalDoctor.SecondName,
			LastName:   citizen.PersonalDoctor.LastName,
			UIN:        citizen.PersonalDoctor.UIN,
			Email:      citizen.PersonalDoctor.Email,
		},
	}

	return nil
}

func (c *citizenService) FindAllAvailablePharmacies(prescriptionId *dto.CitizenAvailablePharmacyGet, availablePharmacies *[]dto.CitizenAvailablePharmacy) error {
	branches := new([]models.PharmacyBranch)

	if err := c.citizenRepo.FindAvailablePharmacies(prescriptionId.PrescriptionId, branches); err != nil {
		return err
	}

	*availablePharmacies = make([]dto.CitizenAvailablePharmacy, len(*branches))

	for i, branch := range *branches {
		(*availablePharmacies)[i] = dto.CitizenAvailablePharmacy{
			Name:      branch.Name,
			Latitude:  branch.Latitude,
			Longitude: branch.Longitude,
		}
	}

	return nil
}

func (c *citizenService) ListPrescriptions(citizenId uuid.UUID, prescriptionsDto *[]dto.CitizenPrescription) error {
	prescriptions := new([]models.Prescription)

	if err := c.citizenRepo.FindAllPrescriptions(citizenId, prescriptions); err != nil {
		return err
	}

	*prescriptionsDto = make([]dto.CitizenPrescription, len(*prescriptions))

	for i, prescription := range *prescriptions {
		(*prescriptionsDto)[i] = dto.CitizenPrescription{
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
