package get_all_profiles

import (
	"context"
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

func (h *handler) GetAllProfiles(ctx context.Context, in *profile.GetAllProfilesIn, opts ...grpc.CallOption) (*profile.GetAllProfilesOut, error){
	var(
		handlerError profile.Errors
	)

	profiles, err := h.profileDomain.GetAllProfiles()
	if err != nil{
		switch err{
		default:
			handlerError = profile.Errors_InternalServiceError
			h.log.Errorf("internal service error: %s", err.Error())
			return &profile.GetAllProfilesOut{
				Success: false,
				Error: &handlerError,
			}, nil
		}
	}

	ProfilesData := make([]*profile.ProfileData, len(profiles))
	for _, profileData := range profiles{
		ProfilesData = append(ProfilesData, &profile.ProfileData{
			UserID: profileData.UserID,
			UserMail: profileData.UserMail,
			UserName: profileData.UserName,
			UserSurname: profileData.UserSurname,
			UserRole: profileData.UserRole,
		})
	}

	return &profile.GetAllProfilesOut{
		Profiles: ProfilesData,
	}, nil
}
