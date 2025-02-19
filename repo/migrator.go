package repo

import (
	"medico/config"
	"medico/models"
)

type MigratorRepo interface {
	MigrateAll() error
}

type migratorRepo struct {
	repo Repository
}

func NewMigratorRepo() MigratorRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return migratorRepo{repo: CreateNewRepository(databaseConfig)}
}

func (m migratorRepo) MigrateAll() error {
	if err := m.repo.DropTableIfExists(models.AuthorizationHolder{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.ActiveIngredient{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.ActiveIngredientInteraction{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.Medicament{}); err != nil {
		return err
	}

	if err := m.repo.DropTableIfExists(models.PharmacyOwnerAuth{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.PharmacyOwner{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.PharmacyBrand{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.PharmacyBranch{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.Pharmacist{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.PharmacyBranchStorage{}); err != nil {
		return err
	}

	if err := m.repo.DropTableIfExists(models.Hospital{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.Doctor{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.DoctorAuth{}); err != nil {
		return err
	}

	if err := m.repo.DropTableIfExists(models.Citizen{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.CitizenAddress{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.CitizenAuth{}); err != nil {
		return err
	}

	if err := m.repo.DropTableIfExists(models.Prescription{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.PrescriptionMedicament{}); err != nil {
		return err
	}

	if err := m.repo.DropTableIfExists(models.ModeratorAuth{}); err != nil {
		return err
	}
	if err := m.repo.DropTableIfExists(models.Moderator{}); err != nil {
		return err
	}

	if err := m.repo.DropTableIfExists(models.AdminAuth{}); err != nil {
		return err
	}

	if err := m.repo.AutoMigrate(models.AuthorizationHolder{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.ActiveIngredient{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.ActiveIngredientInteraction{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.Medicament{}); err != nil {
		return err
	}

	if err := m.repo.AutoMigrate(models.PharmacyOwnerAuth{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.PharmacyOwner{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.PharmacyBrand{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.PharmacyBranch{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.Pharmacist{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.PharmacyBranchStorage{}); err != nil {
		return err
	}

	if err := m.repo.AutoMigrate(models.Hospital{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.Doctor{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.DoctorAuth{}); err != nil {
		return err
	}

	if err := m.repo.AutoMigrate(models.Citizen{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.CitizenAddress{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.CitizenAuth{}); err != nil {
		return err
	}

	if err := m.repo.AutoMigrate(models.Prescription{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.PrescriptionMedicament{}); err != nil {
		return err
	}

	if err := m.repo.AutoMigrate(models.ModeratorAuth{}); err != nil {
		return err
	}
	if err := m.repo.AutoMigrate(models.Moderator{}); err != nil {
		return err
	}

	if err := m.repo.AutoMigrate(models.AdminAuth{}); err != nil {
		return err
	}

	return nil
}
