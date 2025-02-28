package repo

import (
	"github.com/google/uuid"
	"medico/config"
	"medico/models"
)

type CitizenRepo interface {
	FindAuthByEmail(email string, citizenAuth *models.CitizenAuth) error
	FindMedicalInfo(citizenId uuid.UUID, citizen *models.Citizen) error
	FindAllPrescriptions(citizenId uuid.UUID, prescriptions *[]models.Prescription) error
	FindPersonalDoctor(citizenId uuid.UUID, doctor *models.Doctor) error
	FindAvailablePharmacies(prescriptionId uuid.UUID, branches *[]models.PharmacyBranch) error
}

type citizenRepo struct {
	repo Repository
}

func NewCitizenRepo() CitizenRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &citizenRepo{repo: CreateNewRepository(databaseConfig)}
}

func (c *citizenRepo) FindAuthByEmail(email string, citizenAuth *models.CitizenAuth) error {
	return c.repo.First(citizenAuth, "email = ?", email).Error
}

func (c *citizenRepo) FindMedicalInfo(citizenId uuid.UUID, citizen *models.Citizen) error {
	return c.repo.First(citizen, "id = ?", citizenId).Error
}

func (c *citizenRepo) FindPersonalDoctor(citizenId uuid.UUID, doctor *models.Doctor) error {
	return c.repo.Model(models.Doctor{}).Joins("JOIN citizens ON citizens.personal_doctor_id = doctors.id").First(doctor, "citizens.id = ?", citizenId).Error
}

func (c *citizenRepo) FindAllPrescriptions(citizenId uuid.UUID, prescriptions *[]models.Prescription) error {
	return c.repo.Preload("Doctor").Preload("Medicaments.Medicament").Find(prescriptions, "citizen_id = ?", citizenId).Error
}

func (c *citizenRepo) FindAvailablePharmacies(prescriptionId uuid.UUID, branches *[]models.PharmacyBranch) error {
	return c.repo.Model(models.PharmacyBranchStorage{}).
		Joins("LEFT JOIN pharmacy_branch ON pharmacy_branch.id = pharmacy_branch_storage.pharmacy_branch_id").
		Joins("LEFT JOIN prescription_medicament ON pharmacy_branch_storage.medicament_id = prescription_medicament.medicament_id AND prescription_medicament.prescription_id = ?", prescriptionId).
		Where("pharmacy_branch_storage.quantity > prescription_medicament.quantity").
		Find(branches).Error
}
