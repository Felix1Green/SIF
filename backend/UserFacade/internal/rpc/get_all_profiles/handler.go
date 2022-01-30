package get_all_profiles

import (
	"UserFacade/internal/generated/clients/profile"
	"UserFacade/internal/models/handlerErrors"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	profileServiceClient profile.ProfileClient
	log                  *logrus.Logger
}

func NewHandler(profileServiceClient profile.ProfileClient, logger *logrus.Logger) *handler {
	return &handler{
		profileServiceClient: profileServiceClient,
		log:                  logger,
	}
}

// Handle godoc
// @Summary GetAllProfiles
// @Description Get all profiles of the service
// @ID get-all-profiles-id
// @Success 200 {array} profile.ProfileData
// @Failure 503 {object} handlerErrors.Error
// @Router /profiles [get]
func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handle(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	return
}

func (h *handler) handle(w http.ResponseWriter, r *http.Request) {
	getProfileDto := &profile.GetAllProfilesIn{}

	response, err := h.profileServiceClient.GetAllProfiles(
		context.Background(),
		getProfileDto,
	)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if !response.Success {
		hError := handlerErrors.Error{
			ErrorCode:    http.StatusServiceUnavailable,
			ErrorMessage: response.Error.String(),
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		rawOut, _ := json.Marshal(hError)
		_, _ = w.Write(rawOut)
		return
	}

	rawOut, _ := json.Marshal(response.Profiles)

	_, _ = w.Write(rawOut)
	return
}
