package profile

import (
	"github.com/Felix1Green/SIF/backend/ProfileService/internal"
	"github.com/Felix1Green/SIF/backend/ProfileService/internal/components"
	"github.com/Felix1Green/SIF/backend/ProfileService/internal/entities"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
)

type useCase struct{
	sanitizer        *bluemonday.Policy
	userStorage      components.UserStorage
	log              *logrus.Logger
}

func NewInteractor(userStorage components.UserStorage, log *logrus.Logger) *useCase{
	return &useCase{
		sanitizer: bluemonday.UGCPolicy(),
		userStorage: userStorage,
		log: log,
	}
}

func (u *useCase) CreateProfile(profile *entities.Profile) (*entities.Profile, error){
	if !u.isDataCorrect(profile){
		return nil, internal.ProfileDataNotProvidedError
	}

	profile, err := u.userStorage.CreateProfile(profile)
	if err != nil{
		return nil, err
	}

	return profile, nil
}

func (u *useCase) GetProfileByUserID(userID int64) (*entities.Profile, error){
	profile, err := u.userStorage.GetProfileByID(userID)
	if err != nil{
		return nil, err
	}

	return profile, nil
}

func (u *useCase) GetAllProfiles() ([]*entities.Profile, error){
	profiles, err := u.userStorage.GetAllProfiles()
	if err != nil{
		return nil, err
	}

	return profiles, err
}

func (u *useCase) isDataCorrect(profile *entities.Profile) (bool){
	internal.SanitizeInput(
		u.sanitizer,
		&profile.UserMail,
		profile.UserName,
		profile.UserRole,
		profile.UserSurname,
	)

	if profile.UserMail == "" || profile.UserID == 0{
		return false
	}

	return true
}