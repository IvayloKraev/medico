package repo

import (
	"github.com/google/uuid"
	"medico/config"
	"medico/models"
)

type ModeratorRepo interface {
	FindAuthByEmail(email string, moderatorAuth *models.ModeratorAuth) error

	CreateDoctor(doctorAuth *models.DoctorAuth) error
	DeleteDoctor(doctorId uuid.UUID) error
	FindAllDoctors(doctors *[]models.Doctor) error

	CreateMedicament(medicament *models.Medicament) error
	DeleteMedicament(medicamentId uuid.UUID) error
	FindAllMedicaments(medicaments *[]models.Medicament) error

	CreatePharmacyOwner(owner *models.PharmacyOwnerAuth) error
	CreatePharmacy(pharmacy *models.PharmacyBrand) error
	DeletePharmacyOwner(pharmacyOwnerId uuid.UUID) error
	DeletePharmacy(pharmacyId uuid.UUID) error
	FindAllPharmacies(pharmacies *[]models.PharmacyBrand) error

	CreateCitizen(citizenAuth *models.CitizenAuth) error
	DeleteCitizen(citizenId uuid.UUID) error
	FindAllCitizens(citizens *[]models.Citizen) error
}

type moderatorRepo struct {
	repo Repository
}

func NewModeratorRepo() ModeratorRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &moderatorRepo{repo: CreateNewRepository(databaseConfig)}
}

func (m *moderatorRepo) FindAuthByEmail(email string, auth *models.ModeratorAuth) error {
	return m.repo.First(&auth, "email = ?", email).Error
}

func (m *moderatorRepo) CreateDoctor(doctorAuth *models.DoctorAuth) error {
	return m.repo.Create(doctorAuth).Error
}

func (m *moderatorRepo) DeleteDoctor(doctorId uuid.UUID) error {
	return m.repo.Where("id = ?", doctorId.String()).Delete(models.DoctorAuth{}).Error
}

func (m *moderatorRepo) FindAllDoctors(doctors *[]models.Doctor) error {
	return m.repo.Find(doctors).Error
}

func (m *moderatorRepo) CreateMedicament(medicament *models.Medicament) error {
	return m.repo.Create(medicament).Error
}

func (m *moderatorRepo) DeleteMedicament(medicamentId uuid.UUID) error {
	return m.repo.Where("id = ?", medicamentId.String()).Delete(models.Medicament{}).Error
}

func (m *moderatorRepo) FindAllMedicaments(medicaments *[]models.Medicament) error {
	return m.repo.Find(medicaments).Error
}

func (m *moderatorRepo) CreatePharmacyOwner(owner *models.PharmacyOwnerAuth) error {
	return m.repo.Create(owner).Error
}

func (m *moderatorRepo) CreatePharmacy(pharmacy *models.PharmacyBrand) error {
	return m.repo.Create(pharmacy).Error
}

func (m *moderatorRepo) DeletePharmacyOwner(pharmacyOwnerId uuid.UUID) error {
	return m.repo.Where("id = ?", pharmacyOwnerId.String()).Delete(models.PharmacyOwner{}).Error
}

func (m *moderatorRepo) DeletePharmacy(pharmacyId uuid.UUID) error {
	return m.repo.Where("id = ?", pharmacyId.String()).Delete(models.PharmacyBrand{}).Error
}

func (m *moderatorRepo) FindAllPharmacies(pharmacies *[]models.PharmacyBrand) error {
	return m.repo.Find(pharmacies).Error
}

func (m *moderatorRepo) CreateCitizen(citizenAuth *models.CitizenAuth) error {
	return m.repo.Create(citizenAuth).Error
}

func (m *moderatorRepo) DeleteCitizen(doctorId uuid.UUID) error {
	return m.repo.Where("id = ?", doctorId.String()).Delete(models.CitizenAuth{}).Error
}

func (m *moderatorRepo) FindAllCitizens(citizens *[]models.Citizen) error {
	return m.repo.Find(citizens).Error
}
