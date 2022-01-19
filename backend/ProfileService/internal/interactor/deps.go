package interactor

import "github.com/Felix1Green/SIF/backend/ProfileService/internal/entities"

type ProfileInteract interface{
	CreateProfile(profile *entities.Profile) (*entities.Profile, error)
	GetProfileByUserID(userID int64) (*entities.Profile, error)
	GetAllProfiles() ([]*entities.Profile, error)
}
