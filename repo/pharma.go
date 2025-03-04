package repo

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"medico/config"
	"medico/models"
)

type PharmacyOwnerRepo interface {
	FindAuthByEmail(email string, pharmacyOwner *models.PharmacyOwnerAuth) error
	FindDataById(pharmacyOwnerId uuid.UUID, pharmacyOwner *models.PharmacyOwner) error

	FindPharmacyBrandByOwnerId(pharmacyOwnerId uuid.UUID, pharmacyBrand *models.PharmacyBrand) error
	FindPharmacyBranchesByBrandId(pharmacyBrandId uuid.UUID, pharmacyBranches *[]models.PharmacyBranch) error
	FindPharmacyBranchesByOwnerId(pharmacyOwnerId uuid.UUID, pharmacyBranches *[]models.PharmacyBranch) error
	FindPharmacistsByBranchID(pharmacyBranchId uuid.UUID, pharmacists *[]models.Pharmacist) error
	FindPharmacistsByPharmacyOwnerId(pharmacyOwnerId uuid.UUID, pharmacists *[]models.Pharmacist) error

	CreatePharmacyBranch(pharmacyBranch *models.PharmacyBranch) error
	CreatePharmacist(pharmacist *models.PharmacistAuth) error
}

type pharmacyOwnerRepo struct {
	repo Repository
}

func NewPharmacyOwnerRepo() PharmacyOwnerRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &pharmacyOwnerRepo{
		repo: CreateNewRepository(databaseConfig),
	}
}

func (p pharmacyOwnerRepo) FindAuthByEmail(email string, pharmacyOwner *models.PharmacyOwnerAuth) error {
	return p.repo.First(pharmacyOwner, "email = ?", email).Error
}

func (p pharmacyOwnerRepo) FindDataById(pharmacyOwnerId uuid.UUID, pharmacyOwner *models.PharmacyOwner) error {
	return p.repo.First(pharmacyOwner, "id = ?", pharmacyOwnerId).Error
}

func (p pharmacyOwnerRepo) FindPharmacyBrandByOwnerId(pharmacyOwnerId uuid.UUID, pharmacyBrand *models.PharmacyBrand) error {
	return p.repo.First(pharmacyBrand, "owner_id = ?", pharmacyOwnerId).Error
}

func (p pharmacyOwnerRepo) FindPharmacyBranchesByBrandId(pharmacyBrandId uuid.UUID, pharmacyBranches *[]models.PharmacyBranch) error {
	return p.repo.Find(pharmacyBranches, "pharmacy_brand_id = ?", pharmacyBrandId).Error
}

func (p pharmacyOwnerRepo) FindPharmacyBranchesByOwnerId(pharmacyOwnerId uuid.UUID, pharmacyBranches *[]models.PharmacyBranch) error {
	return p.repo.Model(models.PharmacyBranch{}).
		InnerJoins("INNER JOIN pharmacy_brands ON pharmacy_branches.pharmacy_brand_id = pharmacy_brands.id").
		Where("pharmacy_brands.owner_id = ?", pharmacyOwnerId).
		Find(pharmacyBranches).Error
}

func (p pharmacyOwnerRepo) FindPharmacistsByBranchID(pharmacyBranchId uuid.UUID, pharmacists *[]models.Pharmacist) error {
	return p.repo.Find(pharmacists, "pharmacy_branch_id = ?", pharmacyBranchId).Error
}

func (p pharmacyOwnerRepo) FindPharmacistsByPharmacyOwnerId(pharmacyOwnerId uuid.UUID, pharmacists *[]models.Pharmacist) error {
	return p.repo.Model(models.Pharmacist{}).
		InnerJoins("INNER JOIN pharmacy_branches ON pharmacists.pharmacy_branch_id = pharmacy_branches.id").
		InnerJoins("INNER JOIN pharmacy_brands ON pharmacy_branches.pharmacy_brand_id = pharmacy_brands.id").
		Where("pharmacy_brands.owner_id = ?", pharmacyOwnerId).
		Find(pharmacists).Error
}

func (p pharmacyOwnerRepo) CreatePharmacyBranch(pharmacyBranch *models.PharmacyBranch) error {
	return p.repo.Create(pharmacyBranch).Error
}

func (p pharmacyOwnerRepo) CreatePharmacist(pharmacist *models.PharmacistAuth) error {
	return p.repo.Create(pharmacist).Error
}

type PharmacistRepo interface {
	FindAuthByEmail(email string, pharmacist *models.PharmacistAuth) error

	FindActivePrescriptionsByCitizenUcn(citizenUcn string, activePrescriptions *[]models.Prescription) error
	FulfillWholePrescription(branchId, prescriptionId uuid.UUID) error
	FulfillMedicamentFromPrescription(prescriptionId uuid.UUID, medicamentId uuid.UUID) error

	AddMedicamentToBranchStorage(branchId uuid.UUID, medicamentId uuid.UUID, quantity uint) error
}

type pharmacistRepo struct {
	repo Repository
}

func NewPharmacistRepo() PharmacistRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &pharmacistRepo{
		repo: CreateNewRepository(databaseConfig),
	}
}

func (p pharmacistRepo) FindAuthByEmail(email string, pharmacist *models.PharmacistAuth) error {
	return p.repo.First(pharmacist, "email = ?", email).Error
}

func (p pharmacistRepo) FindActivePrescriptionsByCitizenUcn(citizenUcn string, activePrescriptions *[]models.Prescription) error {
	return p.repo.Preload("Medicaments.Medicament").
		Where("citizen_id IN (?)", p.repo.
			Model(models.Citizen{}).
			Select("id").
			Where("ucn = ?", citizenUcn)).
		Where("state = ?", "active").
		Find(activePrescriptions).Error
}

func (p pharmacistRepo) FulfillWholePrescription(branchId, prescriptionId uuid.UUID) error {
	return p.repo.Transaction(func(tx Repository) error {
		return errors.Join(
			tx.Model(models.Prescription{}).
				Where("id = ?", prescriptionId).
				Update("state", "fulfilled").Error,

			tx.Model(models.PrescriptionMedicament{}).
				Where("prescription_id = ?", prescriptionId).
				Update("fulfilled", true).Error,

			tx.Model(models.PharmacyBranchStorage{}).
				Where("pharmacy_branch_storages.pharmacy_branch_id = ?", branchId).
				Update("pharmacy_branch_storages.quantity",
					gorm.Expr("pharmacy_branch_storages.quantity - (?)",
						tx.Model(models.PrescriptionMedicament{}).
							Select("quantity").
							Where("prescription_id = ?", prescriptionId))).Error,
		)
	})
}

func (p pharmacistRepo) FulfillMedicamentFromPrescription(prescriptionId uuid.UUID, medicamentId uuid.UUID) error {
	return p.repo.Model(models.PrescriptionMedicament{}).
		Where("prescription_id = ? AND medicament_id = ?", prescriptionId, medicamentId).
		Update("fulfilled", true).Error
}

func (p pharmacistRepo) AddMedicamentToBranchStorage(branchId uuid.UUID, medicamentId uuid.UUID, quantity uint) error {
	pharmacyBranchStorage := models.PharmacyBranchStorage{
		PharmacyBranchID: branchId,
		MedicamentID:     medicamentId,
		Quantity:         quantity,
	}

	if err := p.repo.First(&pharmacyBranchStorage, "pharmacy_branch_id = ? AND medicament_id = ?", branchId, medicamentId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.repo.Create(&pharmacyBranchStorage)
		} else {
			return err
		}
	} else {
		p.repo.Model(&pharmacyBranchStorage).Update("quantity", gorm.Expr("quantity + ?", quantity))
	}

	return nil
}
