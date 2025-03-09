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
	return c.repo.Model(models.PharmacyBranch{}).
		Joins("LEFT JOIN pharmacy_branch_storages ON pharmacy_branches.id = pharmacy_branch_storages.pharmacy_branch_id").
		Joins("LEFT JOIN prescription_medicaments ON pharmacy_branch_storages.medicament_id = prescription_medicaments.medicament_id AND prescription_medicaments.prescription_id = ?", prescriptionId).
		Where("pharmacy_branch_storages.quantity >= prescription_medicaments.quantity").
		Find(branches).Error
}
