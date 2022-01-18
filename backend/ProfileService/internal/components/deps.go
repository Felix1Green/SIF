package components

import "github.com/Felix1Green/SIF/backend/ProfileService/internal/entities"

type UserStorage interface{
	CreateProfile(profile *entities.Profile) (*entities.Profile, error)
	GetProfileByID(userID int64) (*entities.Profile, error)
	GetAllProfiles()([]*entities.Profile, error)
}