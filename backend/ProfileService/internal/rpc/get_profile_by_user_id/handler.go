package get_profile_by_user_id

import (
	"context"
	"github.com/Felix1Green/SIF/backend/ProfileService/internal"
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

func (h *handler) GetProfileByUserID(ctx context.Context, in *profile.GetProfileByUserIDIn, opts ...grpc.CallOption) (*profile.GetProfileByUserIDOut, error){
	var(
		handlerError profile.Errors
	)

	profileData, err := h.profileDomain.GetProfileByUserID(in.UserID)
	if err != nil{
		switch err{
		case internal.ProfileNotFoundError:
			handlerError = profile.Errors_ProfileNotFound
		default:
			handlerError = profile.Errors_InternalServiceError
		}
		return &profile.GetProfileByUserIDOut{
			Success: false,
			Error: &handlerError,
		}, nil
	}

	return &profile.GetProfileByUserIDOut{
		Profile: &profile.ProfileData{
			UserID: profileData.UserID,
			UserMail: profileData.UserMail,
			UserRole: profileData.UserRole,
			UserSurname: profileData.UserSurname,
			UserName: profileData.UserName,
		},
	}, nil
}