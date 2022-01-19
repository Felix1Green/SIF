package create_profile

import (
	"context"
	"github.com/Felix1Green/SIF/backend/ProfileService/internal"
	"github.com/Felix1Green/SIF/backend/ProfileService/internal/entities"
	"github.com/Felix1Green/SIF/backend/ProfileService/internal/generated/service/profile"
	"github.com/Felix1Green/SIF/backend/ProfileService/internal/interactor"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type handler struct{
	profileDomain interactor.ProfileInteract
	log *logrus.Logger
}

func NewHandler(profileDomain interactor.ProfileInteract, logger *logrus.Logger) *handler{
	return &handler{
		profileDomain: profileDomain,
		log: logger,
	}
}

func (h *handler) CreateProfile(ctx context.Context, in *profile.CreateProfileIn, opts ...grpc.CallOption) (*profile.CreateProfileOut, error){
	var(
		handlerError profile.Errors
	)
	if in.Profile == nil{
		handlerError = profile.Errors_ProfileDataNotProvided
		return &profile.CreateProfileOut{
			Success: false,
			Error: &handlerError,
		}, nil
	}

	createdProfile, err := h.profileDomain.CreateProfile(&entities.Profile{
		UserID: in.Profile.UserID,
	})
	if err != nil{
		switch err{
		case internal.ProfileDataNotProvidedError:
			handlerError = profile.Errors_ProfileDataNotProvided
		case internal.ProfileAlreadyExists:
			handlerError = profile.Errors_ProfileAlreadyExists
		default:
			handlerError = profile.Errors_InternalServiceError
		}
		return &profile.CreateProfileOut{
			Success: false,
			Error: &handlerError,
		}, nil
	}

	return &profile.CreateProfileOut{
		Success: true,
		Profile: &profile.ProfileData{
			UserID: createdProfile.UserID,
			UserName: createdProfile.UserName,
			UserSurname: createdProfile.UserSurname,
			UserRole: createdProfile.UserRole,
			UserMail: createdProfile.UserMail,
		},
	}, nil
}
