package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"medico/dto"
	"medico/models"
	"medico/repo"
	"medico/session"
	"strings"
	"time"
)

// DOCTORS

type DoctorModeratorService interface {
	AuthenticateWithEmailAndPassword(email, password string) (uuid.UUID, error)
	GetModeratorDetails(moderatorID uuid.UUID, moderator *models.Moderator) error

	CreateAuthenticationSession(moderatorId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error

	CreateDoctor(createDoctor *dto.RequestModeratorCreateDoctor) error
	DeleteDoctor(doctorId *dto.QueryModeratorDeleteDoctor) error
	FindAllDoctors(dtoDoctors *[]dto.ResponseModeratorGetDoctors) error
}

type doctorModeratorService struct {
	authSession session.AuthSession
	repo        repo.DoctorModeratorRepo
}

func NewDoctorModeratorService() DoctorModeratorService {
	return &doctorModeratorService{
		authSession: session.NewAuthSession("moderator:doctor"),
		repo:        repo.NewDoctorModeratorRepo(),
	}
}

func (m *doctorModeratorService) AuthenticateWithEmailAndPassword(email, password string) (moderator uuid.UUID, err error) {
	moderatorAuth := models.ModeratorAuth{}

	if err := m.repo.FindAuthByEmail(email, &moderatorAuth); err != nil {
		return uuid.Nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(moderatorAuth.Password), []byte(password)); err != nil {
		return uuid.Nil, err
	}

	return moderatorAuth.ID, nil
}
func (m *doctorModeratorService) GetModeratorDetails(moderatorID uuid.UUID, moderator *models.Moderator) error {
	return m.repo.FindById(moderatorID, moderator)
}

func (m *doctorModeratorService) CreateAuthenticationSession(moderatorId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return m.authSession.CreateAuthSession(moderatorId)
}
func (m *doctorModeratorService) GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error) {
	return m.authSession.GetAuthSession(sessionID)

}
func (m *doctorModeratorService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return m.authSession.DeleteAuthSession(sessionID)
}

func (m *doctorModeratorService) CreateDoctor(createDoctor *dto.RequestModeratorCreateDoctor) error {
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
func (m *doctorModeratorService) DeleteDoctor(doctorId *dto.QueryModeratorDeleteDoctor) error {
	return m.repo.DeleteDoctor(doctorId.DoctorId)
}
func (m *doctorModeratorService) FindAllDoctors(dtoDoctors *[]dto.ResponseModeratorGetDoctors) error {
	var doctors []models.Doctor

	if err := m.repo.FindAllDoctors(&doctors); err != nil {
		return err
	}

	*dtoDoctors = make([]dto.ResponseModeratorGetDoctors, len(doctors))

	for i, doc := range doctors {
		(*dtoDoctors)[i] = dto.ResponseModeratorGetDoctors{
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

// PHARMA

type PharmaModeratorService interface {
	AuthenticateWithEmailAndPassword(email, password string) (uuid.UUID, error)
	GetModeratorDetails(moderatorID uuid.UUID, moderator *models.Moderator) error

	CreateAuthenticationSession(moderatorId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error

	CreatePharmacyAndOwner(createPharmacy *dto.RequestModeratorCreatePharmacy) error
	DeletePharmacy(pharmacyId *dto.QueryModeratorDeletePharmacy) error
	FindAllPharmacies(dtoPharmacies *[]dto.ResponseModeratorGetPharmacies) error
}

type pharmaModeratorService struct {
	authSession session.AuthSession
	repo        repo.PharmaModeratorRepo
}

func NewPharmaModeratorService() PharmaModeratorService {
	return &pharmaModeratorService{
		authSession: session.NewAuthSession("moderator:pharma"),
		repo:        repo.NewPharmaModeratorRepo(),
	}
}

func (m *pharmaModeratorService) AuthenticateWithEmailAndPassword(email, password string) (uuid.UUID, error) {
	moderatorAuth := models.ModeratorAuth{}

	if err := m.repo.FindAuthByEmail(email, &moderatorAuth); err != nil {
		return uuid.Nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(moderatorAuth.Password), []byte(password)); err != nil {
		return uuid.Nil, err
	}

	return moderatorAuth.ID, nil
}
func (m *pharmaModeratorService) GetModeratorDetails(moderatorID uuid.UUID, moderator *models.Moderator) error {
	return m.repo.FindById(moderatorID, moderator)
}

func (m *pharmaModeratorService) CreateAuthenticationSession(moderatorId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return m.authSession.CreateAuthSession(moderatorId)
}
func (m *pharmaModeratorService) GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error) {
	return m.authSession.GetAuthSession(sessionID)

}
func (m *pharmaModeratorService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return m.authSession.DeleteAuthSession(sessionID)
}

func (m *pharmaModeratorService) CreatePharmacyAndOwner(createPharmacy *dto.RequestModeratorCreatePharmacy) error {
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
		ID:      uuid.New(),
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
func (m *pharmaModeratorService) DeletePharmacy(pharmacyId *dto.QueryModeratorDeletePharmacy) error {
	return m.repo.DeletePharmacy(pharmacyId.PharmacyId)
}
func (m *pharmaModeratorService) FindAllPharmacies(dtoPharmacies *[]dto.ResponseModeratorGetPharmacies) error {
	var pharmacies []models.PharmacyBrand

	if err := m.repo.FindAllPharmacies(&pharmacies); err != nil {
		return err
	}

	*dtoPharmacies = make([]dto.ResponseModeratorGetPharmacies, len(pharmacies))

	for i, pharmacy := range pharmacies {
		(*dtoPharmacies)[i] = dto.ResponseModeratorGetPharmacies{
			ID:        pharmacy.ID,
			Name:      pharmacy.Name,
			OwnerName: pharmacy.Owner.Name,
		}
	}

	return nil
}

// MEDICAMENT

type MedicamentModeratorService interface {
	AuthenticateWithEmailAndPassword(email, password string) (uuid.UUID, error)
	GetModeratorDetails(moderatorID uuid.UUID, moderator *models.Moderator) error

	CreateAuthenticationSession(moderatorId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error

	CreateMedicament(createMedicament *dto.RequestModeratorCreateMedicament) error
	DeleteMedicament(medicamentId *dto.QueryModeratorDeleteMedicament) error
	FindAllMedicaments(dtoMedicaments *[]dto.ResponseModeratorGetMedicaments) error
}

type medicamentModeratorService struct {
	authSession session.AuthSession
	repo        repo.MedicamentModeratorRepo
}

func NewMedicamentModeratorService() MedicamentModeratorService {
	return &medicamentModeratorService{
		authSession: session.NewAuthSession("moderator:medicament"),
		repo:        repo.NewMedicamentModeratorRepo(),
	}
}

func (m *medicamentModeratorService) AuthenticateWithEmailAndPassword(email, password string) (uuid.UUID, error) {
	moderatorAuth := models.ModeratorAuth{}

	if err := m.repo.FindAuthByEmail(email, &moderatorAuth); err != nil {
		return uuid.Nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(moderatorAuth.Password), []byte(password)); err != nil {
		return uuid.Nil, err
	}

	return moderatorAuth.ID, nil
}
func (m *medicamentModeratorService) GetModeratorDetails(moderatorID uuid.UUID, moderator *models.Moderator) error {
	return m.repo.FindById(moderatorID, moderator)
}

func (m *medicamentModeratorService) CreateAuthenticationSession(moderatorId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return m.authSession.CreateAuthSession(moderatorId)
}
func (m *medicamentModeratorService) GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error) {
	return m.authSession.GetAuthSession(sessionID)

}
func (m *medicamentModeratorService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return m.authSession.DeleteAuthSession(sessionID)
}

func (m *medicamentModeratorService) CreateMedicament(createMedicament *dto.RequestModeratorCreateMedicament) error {
	newMedicament := models.Medicament{
		ID:                uuid.New(),
		OfficialName:      createMedicament.OfficialName,
		ActiveIngredients: strings.Join(createMedicament.ActiveIngredients, ","),
		ATC:               createMedicament.ATC,
	}

	return m.repo.CreateMedicament(&newMedicament)
}
func (m *medicamentModeratorService) DeleteMedicament(medicamentId *dto.QueryModeratorDeleteMedicament) error {
	return m.repo.DeleteMedicament(medicamentId.MedicamentId)
}
func (m *medicamentModeratorService) FindAllMedicaments(dtoMedicaments *[]dto.ResponseModeratorGetMedicaments) error {
	var medicaments []models.Medicament

	if err := m.repo.FindAllMedicaments(&medicaments); err != nil {
		return err
	}

	*dtoMedicaments = make([]dto.ResponseModeratorGetMedicaments, len(medicaments))

	for i, medicament := range medicaments {
		(*dtoMedicaments)[i] = dto.ResponseModeratorGetMedicaments{
			ID:                medicament.ID,
			OfficialName:      medicament.OfficialName,
			ATC:               medicament.ATC,
			ActiveIngredients: strings.Split(medicament.ActiveIngredients, ","),
		}
	}

	return nil
}

// CITIZEN

type CitizenModeratorService interface {
	AuthenticateWithEmailAndPassword(email, password string) (uuid.UUID, error)
	GetModeratorDetails(moderatorID uuid.UUID, moderator *models.Moderator) error

	CreateAuthenticationSession(moderatorId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error

	CreateCitizen(createCitizen *dto.RequestModeratorCreateCitizen) error
	DeleteCitizen(citizenId *dto.QueryModeratorDeleteCitizen) error
	FindAllCitizens(dtoCitizens *[]dto.ResponseModeratorGetCitizens) error
}

type citizenModeratorService struct {
	authSession session.AuthSession
	repo        repo.CitizenModeratorRepo
}

func NewCitizenModeratorService() CitizenModeratorService {
	return &citizenModeratorService{
		authSession: session.NewAuthSession("moderator:citizen"),
		repo:        repo.NewCitizenModeratorRepo(),
	}
}

func (m *citizenModeratorService) AuthenticateWithEmailAndPassword(email, password string) (uuid.UUID, error) {
	moderatorAuth := models.ModeratorAuth{}

	if err := m.repo.FindAuthByEmail(email, &moderatorAuth); err != nil {
		return uuid.Nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(moderatorAuth.Password), []byte(password)); err != nil {
		return uuid.Nil, err
	}

	return moderatorAuth.ID, nil
}
func (m *citizenModeratorService) GetModeratorDetails(moderatorID uuid.UUID, moderator *models.Moderator) error {
	return m.repo.FindById(moderatorID, moderator)
}

func (m *citizenModeratorService) CreateAuthenticationSession(moderatorId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return m.authSession.CreateAuthSession(moderatorId)
}
func (m *citizenModeratorService) GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error) {
	return m.authSession.GetAuthSession(sessionID)

}
func (m *citizenModeratorService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return m.authSession.DeleteAuthSession(sessionID)
}

func (m *citizenModeratorService) CreateCitizen(createCitizen *dto.RequestModeratorCreateCitizen) error {
	password, err := bcrypt.GenerateFromPassword([]byte(createCitizen.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newCitizenAuth := models.CitizenAuth{
		ID:       uuid.New(),
		Email:    createCitizen.Email,
		Password: string(password),
		Citizen: models.Citizen{
			FirstName:        createCitizen.FirstName,
			SecondName:       createCitizen.SecondName,
			LastName:         createCitizen.LastName,
			UCN:              createCitizen.UCN,
			Email:            createCitizen.Email,
			PersonalDoctorID: createCitizen.PersonalDoctorId,
		},
	}

	if err := m.repo.CreateCitizen(&newCitizenAuth); err != nil {
		return err
	}

	return nil
}
func (m *citizenModeratorService) DeleteCitizen(citizenId *dto.QueryModeratorDeleteCitizen) error {
	return m.repo.DeleteCitizen(citizenId.CitizenId)
}
func (m *citizenModeratorService) FindAllCitizens(dtoCitizens *[]dto.ResponseModeratorGetCitizens) error {
	var citizens []models.Citizen

	if err := m.repo.FindAllCitizens(&citizens); err != nil {
		return err
	}

	*dtoCitizens = make([]dto.ResponseModeratorGetCitizens, len(citizens))

	for i, citizen := range citizens {
		(*dtoCitizens)[i] = dto.ResponseModeratorGetCitizens{
			ID:         citizen.ID,
			FirstName:  citizen.FirstName,
			SecondName: citizen.SecondName,
			LastName:   citizen.LastName,
			UCN:        citizen.UCN,
		}
	}

	return nil
}
