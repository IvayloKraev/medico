package repo

import (
	"github.com/google/uuid"
	"medico/config"
	"medico/models"
)

type DoctorRepo interface {
	FindAuthByEmail(email string, doctorAuth *models.DoctorAuth) error
	FindCitizenByUcn(doctorId uuid.UUID, citizenUcn string, citizen *models.Citizen) error
	FindPrescriptionsByCitizenId(citizenId uuid.UUID, prescriptions *[]models.Prescription) error

	FindMedicamentByName(name string, medicament *models.Medicament) error

	CreatePrescription(prescription *models.Prescription) error
}

type doctorRepo struct {
	repo Repository
}

func NewDoctorRepo() DoctorRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &doctorRepo{
		repo: CreateNewRepository(databaseConfig),
	}
}

func (d doctorRepo) FindAuthByEmail(email string, doctorAuth *models.DoctorAuth) error {
	return d.repo.First(doctorAuth, "email = ?", email).Error
}

func (d doctorRepo) FindCitizenByUcn(doctorId uuid.UUID, citizenUcn string, citizen *models.Citizen) error {
	return d.repo.First(citizen, "personal_doctor_id = ? AND ucn = ?", doctorId, citizenUcn).Error
}

func (d doctorRepo) FindPrescriptionsByCitizenId(citizenId uuid.UUID, prescriptions *[]models.Prescription) error {
	return d.repo.Preload("Medicaments.Medicament").Find(prescriptions, "citizen_id = ?", citizenId).Error
}

func (d doctorRepo) FindMedicamentByName(name string, medicament *models.Medicament) error {
	return d.repo.First(medicament, "official_name = ?", name).Error
}

func (d doctorRepo) CreatePrescription(prescription *models.Prescription) error {
	return d.repo.Create(prescription).Error
}
