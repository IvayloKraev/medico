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

type DoctorService interface {
	AuthenticateByEmailAndPassword(email string, password string, doctorAuth *models.DoctorAuth) error

	CreateAuthenticationSession(doctorId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error

	GetCitizenInfo(doctorId uuid.UUID, citizenUcn string, citizenDto *dto.DoctorCitizenInfo) error
	GetCitizensPrescriptions(doctorId, citizenId uuid.UUID, citizenPrescriptionDto *[]dto.DoctorGetCitizenPrescriptionResponse) error
	CreatePrescription(doctorId, citizenId uuid.UUID, newPrescriptionDto *dto.DoctorCreatePrescription) error
}

type doctorService struct {
	authSession session.AuthSession
	repo        repo.DoctorRepo
}

func NewDoctorService() DoctorService {
	return &doctorService{
		authSession: session.NewAuthSession("doctor"),
		repo:        repo.NewDoctorRepo()}
}

func (d *doctorService) AuthenticateByEmailAndPassword(email string, password string, doctorAuth *models.DoctorAuth) error {
	if err := d.repo.FindAuthByEmail(email, doctorAuth); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(doctorAuth.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (d *doctorService) CreateAuthenticationSession(doctorId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return d.authSession.CreateAuthSession(doctorId)
}

func (d *doctorService) GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error) {
	return d.authSession.GetAuthSession(sessionID)
}

func (d *doctorService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return d.authSession.DeleteAuthSession(sessionID)
}

func (d *doctorService) GetCitizenInfo(doctorId uuid.UUID, citizenUcn string, citizenDto *dto.DoctorCitizenInfo) error {
	citizen := models.Citizen{}

	if err := d.repo.FindCitizenByUcn(doctorId, citizenUcn, &citizen); err != nil {
		return err
	}

	citizenDto.ID = citizen.ID
	citizenDto.FirstName = citizen.FirstName
	citizenDto.SecondName = citizen.SecondName
	citizenDto.LastName = citizen.LastName
	citizenDto.Email = citizen.Email
	citizenDto.BirthDate = citizen.Birthday

	return nil
}

func (d *doctorService) GetCitizensPrescriptions(doctorId, citizenId uuid.UUID, citizenPrescriptionDto *[]dto.DoctorGetCitizenPrescriptionResponse) error {
	var prescriptions []models.Prescription

	if err := d.repo.FindPrescriptionsByCitizenId(citizenId, &prescriptions); err != nil {
		return err
	}

	*citizenPrescriptionDto = make([]dto.DoctorGetCitizenPrescriptionResponse, len(prescriptions))

	for i, prescription := range prescriptions {
		(*citizenPrescriptionDto)[i] = dto.DoctorGetCitizenPrescriptionResponse{
			Id:          prescription.ID,
			Name:        prescription.Name,
			State:       string(prescription.State),
			CreatedDate: prescription.CreationDate,
			StartDate:   prescription.StartDate,
			EndDate:     prescription.EndDate,
		}

		(*citizenPrescriptionDto)[i].Medicaments = []struct {
			OfficialName string `json:"official_name"`
			Quantity     uint   `json:"quantity"`
		}(make([]struct {
			OfficialName string
			Quantity     uint
		}, len(prescription.Medicaments)))

		for k, medicament := range prescription.Medicaments {
			(*citizenPrescriptionDto)[i].Medicaments[k].OfficialName = medicament.Medicament.OfficialName
			(*citizenPrescriptionDto)[i].Medicaments[k].Quantity = medicament.Quantity
		}
	}

	return nil
}

func (d *doctorService) CreatePrescription(doctorId uuid.UUID, citizenId uuid.UUID, newPrescriptionDto *dto.DoctorCreatePrescription) error {
	medicaments := make([]models.PrescriptionMedicament, len(newPrescriptionDto.Medicaments))

	for i, medicament := range newPrescriptionDto.Medicaments {
		med := &models.Medicament{}

		if err := d.repo.FindMedicamentByName(medicament.OfficialName, med); err != nil {
			return err
		}
		medicaments[i].MedicamentID = med.ID
		medicaments[i].Quantity = medicament.Quantity
		medicaments[i].Fulfilled = false
	}

	newPrescription := models.Prescription{
		ID:           uuid.New(),
		DoctorID:     doctorId,
		CitizenID:    citizenId,
		Medicaments:  medicaments,
		State:        "active",
		Name:         newPrescriptionDto.Name,
		CreationDate: time.Now(),
		StartDate:    time.Now(),
		EndDate:      newPrescriptionDto.EndDate,
	}

	return d.repo.CreatePrescription(&newPrescription)
}
