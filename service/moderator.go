package service

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"medico/dto"
	"medico/models"
	"medico/repo"
	"medico/session"
	"strings"
	"time"
)

type ModeratorService interface {
	AuthenticateWithEmailAndPassword(email, password string, moderatorAuth *models.ModeratorAuth) error

	CreateAuthenticationSession(moderatorId uuid.UUID, moderatorType models.ModeratorType) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID, moderatorType models.ModeratorType) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID, moderatorType models.ModeratorType) error

	CreateDoctor(createDoctor *dto.ModeratorCreateDoctor) error
	DeleteDoctor(doctorId *dto.ModeratorDeleteDoctor) error
	FindAllDoctors(dtoDoctors *[]dto.ModeratorGetDoctors) error

	CreateMedicament(createMedicament *dto.ModeratorCreateMedicament) error
	DeleteMedicament(medicamentId *dto.ModeratorDeleteMedicament) error
	FindAllMedicaments(dtoMedicaments *[]dto.ModeratorGetMedicaments) error

	CreatePharmacyAndOwner(createPharmacy *dto.ModeratorCreatePharmacy) error
	DeletePharmacy(pharmacyId *dto.ModeratorDeletePharmacy) error
	FindAllPharmacies(dtoPharmacies *[]dto.ModeratorGetPharmacies) error

	CreateCitizen(createCitizen *dto.ModeratorCreateCitizen) error
	DeleteCitizen(citizenId *dto.ModeratorDeleteCitizen) error
	FindAllCitizens(dtoCitizens *[]dto.ModeratorGetCitizens) error
}

type moderatorService struct {
	doctorModAuthSession     session.AuthSession
	citizenModAuthSession    session.AuthSession
	medicamentModAuthSession session.AuthSession
	pharmacyModAuthSession   session.AuthSession
	repo                     repo.ModeratorRepo
}

func NewModeratorService() ModeratorService {
	return &moderatorService{
		doctorModAuthSession:     session.NewAuthSession("moderator:doctor"),
		citizenModAuthSession:    session.NewAuthSession("moderator:citizen"),
		medicamentModAuthSession: session.NewAuthSession("moderator:medicament"),
		pharmacyModAuthSession:   session.NewAuthSession("moderator:pharmacy"),
		repo:                     repo.NewModeratorRepo(),
	}
}

func (m *moderatorService) AuthenticateWithEmailAndPassword(email, password string, moderatorAuth *models.ModeratorAuth) error {
	if err := m.repo.FindAuthByEmail(email, moderatorAuth); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(moderatorAuth.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (m *moderatorService) CreateAuthenticationSession(moderatorId uuid.UUID, moderatorType models.ModeratorType) (uuid.UUID, time.Duration, error) {
	switch moderatorType {
	case "doctorMod":
		return m.doctorModAuthSession.CreateAuthSession(moderatorId)
	case "citizenMod":
		return m.citizenModAuthSession.CreateAuthSession(moderatorId)
	case "medicamentMod":
		return m.medicamentModAuthSession.CreateAuthSession(moderatorId)
	case "pharmacyMod":
		return m.pharmacyModAuthSession.CreateAuthSession(moderatorId)
	default:
		return uuid.Nil, 0, errors.New("invalid type")
	}
}
func (m *moderatorService) GetAuthenticationSession(sessionID uuid.UUID, moderatorType models.ModeratorType) (uuid.UUID, error) {
	switch moderatorType {
	case "doctorMod":
		return m.doctorModAuthSession.GetAuthSession(sessionID)
	case "citizenMod":
		return m.citizenModAuthSession.GetAuthSession(sessionID)
	case "medicamentMod":
		return m.medicamentModAuthSession.GetAuthSession(sessionID)
	case "pharmacyMod":
		return m.pharmacyModAuthSession.GetAuthSession(sessionID)
	default:
		return uuid.Nil, errors.New("invalid type")
	}
}
func (m *moderatorService) DeleteAuthenticationSession(sessionID uuid.UUID, moderatorType models.ModeratorType) error {
	switch moderatorType {
	case "doctorMod":
		return m.doctorModAuthSession.DeleteAuthSession(sessionID)
	case "citizenMod":
		return m.citizenModAuthSession.DeleteAuthSession(sessionID)
	case "medicamentMod":
		return m.medicamentModAuthSession.DeleteAuthSession(sessionID)
	case "pharmacyMod":
		return m.pharmacyModAuthSession.DeleteAuthSession(sessionID)
	default:
		return errors.New("invalid type")
	}
}

func (m *moderatorService) CreateDoctor(createDoctor *dto.ModeratorCreateDoctor) error {
	password, err := bcrypt.GenerateFromPassword([]byte(createDoctor.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newDoctorAuth := models.DoctorAuth{
		ID:       uuid.New(),
		Email:    createDoctor.Email,
		Password: string(password),
		Doctor: models.Doctor{
			FirstName:  createDoctor.FirstName,
			SecondName: createDoctor.SecondName,
			LastName:   createDoctor.LastName,
			UIN:        createDoctor.UIN,
		},
	}

	if err := m.repo.CreateDoctor(&newDoctorAuth); err != nil {
		return err
	}

	return nil
}
func (m *moderatorService) DeleteDoctor(doctorId *dto.ModeratorDeleteDoctor) error {
	return m.repo.DeleteDoctor(doctorId.DoctorId)
}
func (m *moderatorService) FindAllDoctors(dtoDoctors *[]dto.ModeratorGetDoctors) error {
	var doctors []models.Doctor

	if err := m.repo.FindAllDoctors(&doctors); err != nil {
		return err
	}

	*dtoDoctors = make([]dto.ModeratorGetDoctors, len(doctors))

	for i, doc := range doctors {
		(*dtoDoctors)[i] = dto.ModeratorGetDoctors{
			ID:         doc.ID,
			FirstName:  doc.FirstName,
			SecondName: doc.SecondName,
			LastName:   doc.LastName,
			Email:      doc.Email,
			UIN:        doc.UIN,
		}
	}

	return nil
}

func (m *moderatorService) CreateMedicament(createMedicament *dto.ModeratorCreateMedicament) error {
	newMedicament := models.Medicament{
		ID:                uuid.New(),
		OfficialName:      createMedicament.OfficialName,
		ActiveIngredients: strings.Join(createMedicament.ActiveIngredients, ","),
		ATC:               createMedicament.ATC,
	}

	return m.repo.CreateMedicament(&newMedicament)
}
func (m *moderatorService) DeleteMedicament(medicamentId *dto.ModeratorDeleteMedicament) error {
	return m.repo.DeleteMedicament(medicamentId.MedicamentId)
}
func (m *moderatorService) FindAllMedicaments(dtoMedicaments *[]dto.ModeratorGetMedicaments) error {
	var medicaments []models.Medicament

	if err := m.repo.FindAllMedicaments(&medicaments); err != nil {
		return err
	}

	*dtoMedicaments = make([]dto.ModeratorGetMedicaments, len(medicaments))

	for i, medicament := range medicaments {
		(*dtoMedicaments)[i] = dto.ModeratorGetMedicaments{
			ID:                medicament.ID,
			OfficialName:      medicament.OfficialName,
			ATC:               medicament.ATC,
			ActiveIngredients: strings.Split(medicament.ActiveIngredients, ","),
		}
	}

	return nil
}

func (m *moderatorService) CreatePharmacyAndOwner(createPharmacy *dto.ModeratorCreatePharmacy) error {
	password, err := bcrypt.GenerateFromPassword([]byte(createPharmacy.OwnerPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newPharmacyOwnerAuth := models.PharmacyOwnerAuth{
		ID:       uuid.New(),
		Email:    createPharmacy.OwnerEmail,
		Password: string(password),
		PharmacyOwner: models.PharmacyOwner{
			Name: createPharmacy.OwnerEmail,
		},
	}

	newPharmacy := models.PharmacyBrand{
		Name:    createPharmacy.Name,
		OwnerID: newPharmacyOwnerAuth.ID,
	}

	if err := m.repo.CreatePharmacyOwner(&newPharmacyOwnerAuth); err != nil {
		return err
	}

	if err := m.repo.CreatePharmacy(&newPharmacy); err != nil {
		return err
	}

	return nil
}
func (m *moderatorService) DeletePharmacy(pharmacyId *dto.ModeratorDeletePharmacy) error {
	return m.repo.DeletePharmacy(pharmacyId.PharmacyId)
}
func (m *moderatorService) FindAllPharmacies(dtoPharmacies *[]dto.ModeratorGetPharmacies) error {
	var pharmacies []models.PharmacyBrand

	if err := m.repo.FindAllPharmacies(&pharmacies); err != nil {
		return err
	}

	*dtoPharmacies = make([]dto.ModeratorGetPharmacies, len(pharmacies))

	for i, pharmacy := range pharmacies {
		(*dtoPharmacies)[i] = dto.ModeratorGetPharmacies{
			ID:        pharmacy.ID,
			Name:      pharmacy.Name,
			OwnerName: pharmacy.Owner.Name,
		}
	}

	return nil
}

func (m *moderatorService) CreateCitizen(createCitizen *dto.ModeratorCreateCitizen) error {
	password, err := bcrypt.GenerateFromPassword([]byte(createCitizen.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newCitizenAuth := models.CitizenAuth{
		ID:       uuid.New(),
		Email:    createCitizen.Email,
		Password: string(password),
		Citizen: models.Citizen{
			FirstName:  createCitizen.FirstName,
			SecondName: createCitizen.SecondName,
			LastName:   createCitizen.LastName,
			UCN:        createCitizen.UCN,
			Email:      createCitizen.Email,
		},
	}

	if err := m.repo.CreateCitizen(&newCitizenAuth); err != nil {
		return err
	}

	return nil
}
func (m *moderatorService) DeleteCitizen(citizenId *dto.ModeratorDeleteCitizen) error {
	return m.repo.DeleteCitizen(citizenId.CitizenId)
}
func (m *moderatorService) FindAllCitizens(dtoCitizens *[]dto.ModeratorGetCitizens) error {
	var citizens []models.Citizen

	if err := m.repo.FindAllCitizens(&citizens); err != nil {
		return err
	}

	*dtoCitizens = make([]dto.ModeratorGetCitizens, len(citizens))

	for i, citizen := range citizens {
		(*dtoCitizens)[i] = dto.ModeratorGetCitizens{
			ID:         citizen.ID,
			FirstName:  citizen.FirstName,
			SecondName: citizen.SecondName,
			LastName:   citizen.LastName,
			UCN:        citizen.UCN,
		}
	}

	return nil
}
