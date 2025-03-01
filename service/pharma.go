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

type PharmacyOwnerService interface {
	AuthenticateByEmailAndPassword(email string, password string, pharmacyOwnerAuth *models.PharmacyOwnerAuth) error
	CreateAuthenticationSession(pharmacyOwnerId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error

	GetAllBranches(pharmacyOwnerId uuid.UUID, branches *[]dto.ResponsePharmacyOwnerBranches) error
	GetAllPharmacists(pharmacyOwnerId uuid.UUID, pharmacists *[]dto.ResponsePharmacyOwnerPharmacist) error

	NewPharmacyBranch(pharmacyOwnerId uuid.UUID, branch *dto.RequestPharmacyOwnerNewBranch) error
	NewPharmacist(pharmacyOwnerId uuid.UUID, pharmacist *dto.RequestPharmacyOwnerNewPharmacist) error
}

type pharmacyOwnerService struct {
	authSession session.AuthSession
	repo        repo.PharmacyOwnerRepo
}

func NewPharmacyOwnerService() PharmacyOwnerService {
	return &pharmacyOwnerService{
		authSession: session.NewAuthSession("pharmacy:owner"),
		repo:        repo.NewPharmacyOwnerRepo(),
	}
}

func (p pharmacyOwnerService) AuthenticateByEmailAndPassword(email string, password string, pharmacyOwnerAuth *models.PharmacyOwnerAuth) error {
	if err := p.repo.FindAuthByEmail(email, pharmacyOwnerAuth); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(pharmacyOwnerAuth.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (p pharmacyOwnerService) CreateAuthenticationSession(pharmacyOwnerId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return p.authSession.CreateAuthSession(pharmacyOwnerId)
}

func (p pharmacyOwnerService) GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error) {
	return p.authSession.GetAuthSession(sessionID)
}

func (p pharmacyOwnerService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return p.authSession.DeleteAuthSession(sessionID)
}

func (p pharmacyOwnerService) GetAllBranches(pharmacyOwnerId uuid.UUID, branches *[]dto.ResponsePharmacyOwnerBranches) error {
	pharmacyBranches := new([]models.PharmacyBranch)

	if err := p.repo.FindPharmacyBranchesByOwnerId(pharmacyOwnerId, pharmacyBranches); err != nil {
		return err
	}

	*branches = make([]dto.ResponsePharmacyOwnerBranches, len(*pharmacyBranches))

	for i, pharmacyBranch := range *pharmacyBranches {
		(*branches)[i] = dto.ResponsePharmacyOwnerBranches{
			ID:   pharmacyBranch.ID,
			Name: pharmacyBranch.Name,
		}
	}

	return nil
}

func (p pharmacyOwnerService) GetAllPharmacists(pharmacyOwnerId uuid.UUID, pharmacistsDto *[]dto.ResponsePharmacyOwnerPharmacist) error {
	pharmacists := new([]models.Pharmacist)

	if err := p.repo.FindPharmacistsByPharmacyOwnerId(pharmacyOwnerId, pharmacists); err != nil {
		return err
	}

	*pharmacistsDto = make([]dto.ResponsePharmacyOwnerPharmacist, len(*pharmacists))

	for i, pharmacist := range *pharmacists {
		(*pharmacistsDto)[i] = dto.ResponsePharmacyOwnerPharmacist{
			ID:        pharmacist.ID,
			FirstName: pharmacist.FirstName,
			LastName:  pharmacist.Surname,
		}
	}

	return nil
}

func (p pharmacyOwnerService) NewPharmacyBranch(pharmacyOwnerId uuid.UUID, branch *dto.RequestPharmacyOwnerNewBranch) error {
	pharmacyBrand := models.PharmacyBrand{}

	if err := p.repo.FindPharmacyBrandByOwnerId(pharmacyOwnerId, &pharmacyBrand); err != nil {
		return err
	}

	newBranch := models.PharmacyBranch{
		ID:              uuid.New(),
		Name:            branch.Name,
		PharmacyBrandID: pharmacyBrand.ID,
		Latitude:        branch.Latitude,
		Longitude:       branch.Longitude,
	}

	if err := p.repo.CreatePharmacyBranch(&newBranch); err != nil {
		return err
	}

	return nil
}

func (p pharmacyOwnerService) NewPharmacist(pharmacyOwnerId uuid.UUID, pharmacist *dto.RequestPharmacyOwnerNewPharmacist) error {
	password, err := bcrypt.GenerateFromPassword([]byte(pharmacist.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newPharmacist := models.PharmacistAuth{
		ID:       uuid.New(),
		Email:    pharmacist.Email,
		Password: string(password),
		Pharmacist: models.Pharmacist{
			FirstName:        pharmacist.FirstName,
			Surname:          pharmacist.LastName,
			PharmacyBranchID: pharmacist.WorkingBranch,
		},
	}

	if err := p.repo.CreatePharmacist(&newPharmacist); err != nil {
		return err
	}

	return nil
}

type PharmacistService interface {
	AuthenticateByEmailAndPassword(email string, password string, pharmacistAuth *models.PharmacistAuth) error
	CreateAuthenticationSession(pharmacyOwnerId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error

	GetCitizensActivePrescriptions(citizenUcn *dto.QueryPharmacistCitizenPrescriptionGet, prescriptions *[]dto.ResponsePharmacistCitizenPrescription) error
	FulfillWholePrescription(data *dto.RequestPharmacistCitizenFulfillWholePrescription) error
	FulfillMedicamentFromPrescription(data *dto.RequestPharmacistCitizenFulfillMedicamentFromPrescription) error

	AddMedicamentToBranchStorage(data *dto.RequestPharmacistBranchAddMedicament) error
}

type pharmacistService struct {
	authSession session.AuthSession
	repo        repo.PharmacistRepo
}

func NewPharmacistService() PharmacistService {
	return &pharmacistService{
		authSession: session.NewAuthSession("pharmacy:pharmacist"),
		repo:        repo.NewPharmacistRepo(),
	}
}

func (p pharmacistService) AuthenticateByEmailAndPassword(email string, password string, pharmacyOwnerAuth *models.PharmacistAuth) error {
	if err := p.repo.FindAuthByEmail(email, pharmacyOwnerAuth); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(pharmacyOwnerAuth.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (p pharmacistService) CreateAuthenticationSession(pharmacyOwnerId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return p.authSession.CreateAuthSession(pharmacyOwnerId)
}

func (p pharmacistService) GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error) {
	return p.authSession.GetAuthSession(sessionID)
}

func (p pharmacistService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return p.authSession.DeleteAuthSession(sessionID)
}

func (p pharmacistService) GetCitizensActivePrescriptions(citizenUcn *dto.QueryPharmacistCitizenPrescriptionGet, prescriptionsDto *[]dto.ResponsePharmacistCitizenPrescription) error {
	prescriptions := new([]models.Prescription)

	if err := p.repo.FindActivePrescriptionsByCitizenUcn(citizenUcn.CitizenUCN, prescriptions); err != nil {
		return err
	}

	*prescriptionsDto = make([]dto.ResponsePharmacistCitizenPrescription, len(*prescriptions))

	for i, prescription := range *prescriptions {
		(*prescriptionsDto)[i] = dto.ResponsePharmacistCitizenPrescription{
			ID:           prescription.ID,
			Name:         prescription.Name,
			CreationDate: prescription.CreationDate,
			StartDate:    prescription.StartDate,
			EndDate:      prescription.EndDate,
			Medicaments: make([]struct {
				MedicamentName string `json:"medicament_name"`
				Quantity       uint   `json:"quantity"`
				Fulfilled      bool   `json:"fulfilled"`
			}, len(prescription.Medicaments)),
		}

		for k, medicament := range prescription.Medicaments {
			(*prescriptionsDto)[i].Medicaments[k] = struct {
				MedicamentName string `json:"medicament_name"`
				Quantity       uint   `json:"quantity"`
				Fulfilled      bool   `json:"fulfilled"`
			}{
				MedicamentName: medicament.Medicament.OfficialName,
				Quantity:       medicament.Quantity,
				Fulfilled:      medicament.Fulfilled,
			}
		}
	}

	return nil
}

func (p pharmacistService) FulfillWholePrescription(data *dto.RequestPharmacistCitizenFulfillWholePrescription) error {
	return p.repo.FulfillWholePrescription(data.BranchId, data.PrescriptionId)
}

func (p pharmacistService) FulfillMedicamentFromPrescription(data *dto.RequestPharmacistCitizenFulfillMedicamentFromPrescription) error {
	return p.repo.FulfillMedicamentFromPrescription(data.PrescriptionId, data.MedicamentId)
}

func (p pharmacistService) AddMedicamentToBranchStorage(data *dto.RequestPharmacistBranchAddMedicament) error {
	for _, medicament := range data.Medicaments {
		err := p.repo.AddMedicamentToBranchStorage(data.BranchId, medicament.MedicamentId, medicament.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}
