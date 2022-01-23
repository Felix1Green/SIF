package rpc

import (
	"ProfileService/internal"
	"ProfileService/internal/entities"
	"ProfileService/internal/generated/service/profile"
	"ProfileService/internal/interactor"
	"context"
	"github.com/sirupsen/logrus"
)

type handler struct {
	profile.UnimplementedProfileServer
	profileDomain interactor.ProfileInteract
	log           *logrus.Logger
}

func NewHandler(profileDomain interactor.ProfileInteract, logger *logrus.Logger) *handler {
	return &handler{
		profileDomain: profileDomain,
		log:           logger,
	}
}

func (h *handler) CreateProfile(ctx context.Context, in *profile.CreateProfileIn) (*profile.CreateProfileOut, error) {
	var (
		handlerError profile.Errors
	)
	if in.Profile == nil {
		handlerError = profile.Errors_ProfileDataNotProvided
		return &profile.CreateProfileOut{
			Success: false,
			Error:   &handlerError,
		}, nil
	}

	createdProfile, err := h.profileDomain.CreateProfile(&entities.Profile{
		UserID:      in.Profile.UserID,
		UserMail:    in.Profile.UserMail,
		UserRole:    in.Profile.UserRole,
		UserName:    in.Profile.UserName,
		UserSurname: in.Profile.UserSurname,
	})
	if err != nil {
		switch err {
		case internal.ProfileDataNotProvidedError:
			handlerError = profile.Errors_ProfileDataNotProvided
		case internal.ProfileAlreadyExists:
			handlerError = profile.Errors_ProfileAlreadyExists
		default:
			handlerError = profile.Errors_InternalServiceError
		}
		return &profile.CreateProfileOut{
			Success: false,
			Error:   &handlerError,
		}, nil
	}

	return &profile.CreateProfileOut{
		Success: true,
		Profile: &profile.ProfileData{
			UserID:      createdProfile.UserID,
			UserName:    createdProfile.UserName,
			UserSurname: createdProfile.UserSurname,
			UserRole:    createdProfile.UserRole,
			UserMail:    createdProfile.UserMail,
		},
	}, nil
}

func (h *handler) GetAllProfiles(ctx context.Context, in *profile.GetAllProfilesIn) (*profile.GetAllProfilesOut, error) {
	var (
		handlerError profile.Errors
	)

	profiles, err := h.profileDomain.GetAllProfiles()
	if err != nil {
		switch err {
		default:
			handlerError = profile.Errors_InternalServiceError
			h.log.Errorf("internal service error: %s", err.Error())
			return &profile.GetAllProfilesOut{
				Success: false,
				Error:   &handlerError,
			}, nil
		}
	}

	ProfilesData := make([]*profile.ProfileData, 0)
	for _, profileData := range profiles {
		ProfilesData = append(ProfilesData, &profile.ProfileData{
			UserID:      profileData.UserID,
			UserMail:    profileData.UserMail,
			UserName:    profileData.UserName,
			UserSurname: profileData.UserSurname,
			UserRole:    profileData.UserRole,
		})
	}

	return &profile.GetAllProfilesOut{
		Success:  true,
		Profiles: ProfilesData,
	}, nil
}

func (h *handler) GetProfileByUserID(ctx context.Context, in *profile.GetProfileByUserIDIn) (*profile.GetProfileByUserIDOut, error) {
	var (
		handlerError profile.Errors
	)

	profileData, err := h.profileDomain.GetProfileByUserID(in.UserID)
	if err != nil {
		switch err {
		case internal.ProfileNotFoundError:
			handlerError = profile.Errors_ProfileNotFound
		default:
			handlerError = profile.Errors_InternalServiceError
		}
		return &profile.GetProfileByUserIDOut{
			Success: false,
			Error:   &handlerError,
		}, nil
	}

	return &profile.GetProfileByUserIDOut{
		Success: true,
		Profile: &profile.ProfileData{
			UserID:      profileData.UserID,
			UserMail:    profileData.UserMail,
			UserRole:    profileData.UserRole,
			UserSurname: profileData.UserSurname,
			UserName:    profileData.UserName,
		},
	}, nil
}
