package db

import "medico/models"

func Migrate(repository Repository) {

	if err := repository.DropTableIfExists(models.AuthorizationHolder{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.ActiveIngredient{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.ActiveIngredientInteraction{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.Medicament{}); err != nil {
		return
	}

	if err := repository.DropTableIfExists(models.PharmacyBrand{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.PharmacyBranch{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.Pharmacist{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.PharmacyBranchStorage{}); err != nil {
		return
	}

	if err := repository.DropTableIfExists(models.Hospital{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.Doctor{}); err != nil {
		return
	}

	if err := repository.DropTableIfExists(models.Citizen{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.CitizenAddress{}); err != nil {
		return
	}

	if err := repository.DropTableIfExists(models.Prescription{}); err != nil {
		return
	}
	if err := repository.DropTableIfExists(models.PrescriptionMedicament{}); err != nil {
		return
	}

	if err := repository.AutoMigrate(models.AuthorizationHolder{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.ActiveIngredient{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.ActiveIngredientInteraction{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.Medicament{}); err != nil {
		return
	}

	if err := repository.AutoMigrate(models.PharmacyBrand{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.PharmacyBranch{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.Pharmacist{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.PharmacyBranchStorage{}); err != nil {
		return
	}

	if err := repository.AutoMigrate(models.Hospital{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.Doctor{}); err != nil {
		return
	}

	if err := repository.AutoMigrate(models.Citizen{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.CitizenAddress{}); err != nil {
		return
	}

	if err := repository.AutoMigrate(models.Prescription{}); err != nil {
		return
	}
	if err := repository.AutoMigrate(models.PrescriptionMedicament{}); err != nil {
		return
	}
}
