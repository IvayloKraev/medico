package repo

import (
	"github.com/google/uuid"
	"medico/config"
	"medico/models"
)

// DOCTOR

type DoctorModeratorRepo interface {
	FindAuthByEmail(email string, moderator *models.ModeratorAuth) error
	FindById(id uuid.UUID, moderator *models.Moderator) error

	CreateDoctor(doctorAuth *models.DoctorAuth) error
	DeleteDoctor(doctorId uuid.UUID) error
	FindAllDoctors(doctors *[]models.Doctor) error
}

type doctorModeratorRepo struct {
	repo Repository
}

func NewDoctorModeratorRepo() DoctorModeratorRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &doctorModeratorRepo{
		repo: CreateNewRepository(databaseConfig),
	}
}

func (m *doctorModeratorRepo) FindAuthByEmail(email string, moderator *models.ModeratorAuth) error {
	return m.repo.First(moderator, "email = ?", email).Error
}
func (m *doctorModeratorRepo) FindById(id uuid.UUID, moderator *models.Moderator) error {
	return m.repo.First(&moderator, "id = ? AND type = ?", id, models.DoctorMod).Error
}

func (m *doctorModeratorRepo) CreateDoctor(doctorAuth *models.DoctorAuth) error {
	return m.repo.Create(doctorAuth).Error
}
func (m *doctorModeratorRepo) DeleteDoctor(doctorId uuid.UUID) error {
	return m.repo.Where("id = ?", doctorId.String()).Delete(models.DoctorAuth{}).Error
}
func (m *doctorModeratorRepo) FindAllDoctors(doctors *[]models.Doctor) error {
	return m.repo.Find(doctors).Error
}

// PHARMA

type PharmaModeratorRepo interface {
	FindAuthByEmail(email string, moderator *models.ModeratorAuth) error
	FindById(id uuid.UUID, moderator *models.Moderator) error

	CreatePharmacyOwner(owner *models.PharmacyOwnerAuth) error
	CreatePharmacy(pharmacy *models.PharmacyBrand) error
	DeletePharmacyOwner(pharmacyOwnerId uuid.UUID) error
	DeletePharmacy(pharmacyId uuid.UUID) error
	FindAllPharmacies(pharmacies *[]models.PharmacyBrand) error
}

type pharmaModeratorRepo struct {
	repo Repository
}

func NewPharmaModeratorRepo() PharmaModeratorRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &pharmaModeratorRepo{
		repo: CreateNewRepository(databaseConfig),
	}
}

func (m *pharmaModeratorRepo) FindAuthByEmail(email string, moderator *models.ModeratorAuth) error {
	return m.repo.First(moderator, "email = ?", email).Error
}
func (m *pharmaModeratorRepo) FindById(id uuid.UUID, moderator *models.Moderator) error {
	return m.repo.First(&moderator, "id = ? AND type = ?", id, models.PharmacyMod).Error
}

func (m *pharmaModeratorRepo) CreatePharmacyOwner(owner *models.PharmacyOwnerAuth) error {
	return m.repo.Create(owner).Error
}
func (m *pharmaModeratorRepo) CreatePharmacy(pharmacy *models.PharmacyBrand) error {
	return m.repo.Create(pharmacy).Error
}
func (m *pharmaModeratorRepo) DeletePharmacyOwner(pharmacyOwnerId uuid.UUID) error {
	return m.repo.Where("id = ?", pharmacyOwnerId.String()).Delete(models.PharmacyOwner{}).Error
}
func (m *pharmaModeratorRepo) DeletePharmacy(pharmacyId uuid.UUID) error {
	return m.repo.Where("id = ?", pharmacyId.String()).Delete(models.PharmacyBrand{}).Error
}
func (m *pharmaModeratorRepo) FindAllPharmacies(pharmacies *[]models.PharmacyBrand) error {
	return m.repo.Preload("Owner").Find(pharmacies).Error
}

// MEDICAMENT

type MedicamentModeratorRepo interface {
	FindAuthByEmail(email string, moderator *models.ModeratorAuth) error
	FindById(id uuid.UUID, moderator *models.Moderator) error

	CreateMedicament(medicament *models.Medicament) error
	DeleteMedicament(medicamentId uuid.UUID) error
	FindAllMedicaments(medicaments *[]models.Medicament) error
}

type medicamentModeratorRepo struct {
	repo Repository
}

func NewMedicamentModeratorRepo() MedicamentModeratorRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &medicamentModeratorRepo{
		repo: CreateNewRepository(databaseConfig),
	}
}

func (m *medicamentModeratorRepo) FindAuthByEmail(email string, moderator *models.ModeratorAuth) error {
	return m.repo.First(moderator, "email = ?", email).Error
}
func (m *medicamentModeratorRepo) FindById(id uuid.UUID, moderator *models.Moderator) error {
	return m.repo.First(&moderator, "id = ? AND type = ?", id, models.MedicamentMod).Error
}

func (m *medicamentModeratorRepo) CreateMedicament(medicament *models.Medicament) error {
	return m.repo.Create(medicament).Error
}
func (m *medicamentModeratorRepo) DeleteMedicament(medicamentId uuid.UUID) error {
	return m.repo.Where("id = ?", medicamentId.String()).Delete(models.Medicament{}).Error
}
func (m *medicamentModeratorRepo) FindAllMedicaments(medicaments *[]models.Medicament) error {
	return m.repo.Find(medicaments).Error
}

// CITIZEN

type CitizenModeratorRepo interface {
	FindAuthByEmail(email string, moderator *models.ModeratorAuth) error
	FindById(id uuid.UUID, moderator *models.Moderator) error

	CreateCitizen(citizenAuth *models.CitizenAuth) error
	DeleteCitizen(citizenId uuid.UUID) error
	FindAllCitizens(citizens *[]models.Citizen) error
}

type citizenModeratorRepo struct {
	repo Repository
}

func NewCitizenModeratorRepo() CitizenModeratorRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &citizenModeratorRepo{
		repo: CreateNewRepository(databaseConfig),
	}
}

func (m *citizenModeratorRepo) FindAuthByEmail(email string, moderator *models.ModeratorAuth) error {
	return m.repo.First(moderator, "email = ?", email).Error
}
func (m *citizenModeratorRepo) FindById(id uuid.UUID, moderator *models.Moderator) error {
	return m.repo.First(&moderator, "id = ? AND type = ?", id, models.CitizenMod).Error
}

func (m *citizenModeratorRepo) CreateCitizen(citizenAuth *models.CitizenAuth) error {
	return m.repo.Create(citizenAuth).Error
}
func (m *citizenModeratorRepo) DeleteCitizen(doctorId uuid.UUID) error {
	return m.repo.Where("id = ?", doctorId.String()).Delete(models.CitizenAuth{}).Error
}
func (m *citizenModeratorRepo) FindAllCitizens(citizens *[]models.Citizen) error {
	return m.repo.Find(citizens).Error
}
