package register

import (
	"UserFacade/internal/generated/clients/auth"
	"UserFacade/internal/generated/clients/profile"
	"UserFacade/internal/models/handlerErrors"
	"UserFacade/internal/models/handlersDto"
	"UserFacade/internal/models/user"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type handler struct {
	authServiceClient    auth.AuthClient
	profileServiceClient profile.ProfileClient
	log                  *logrus.Logger
}

func NewHandler(authServiceClient auth.AuthClient, profileServiceClient profile.ProfileClient, log *logrus.Logger) *handler {
	return &handler{
		authServiceClient,
		profileServiceClient,
		log,
	}
}

// Handle godoc
// @Summary Register
// @Description Register user
// @ID register-id
// @Param User_info body user.RegisterUser true "user credentials"
// @Success 200
// @Failure 503 {object} handlerErrors.Error
// @Failure 400 {object} handlerErrors.Error
// @Failure 403 {object} handlerErrors.Error
// @Router /register [post]
func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	inputData := user.RegisterUser{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &inputData)
	if inputData.UserMail == "" || inputData.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	RegisterIn := auth.RegisterIn{
		UserName: inputData.UserMail,
		Password: inputData.Password,
	}

	response, err := h.authServiceClient.Register(context.Background(), &RegisterIn)
	if err != nil {
		h.log.Warning(fmt.Sprintf("auth service returned error: %s", err))
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if response.Error != nil {
		h.log.Error(response.Error)
		switch *response.Error {
		case auth.Errors_NotEnoughRightsToCreateUser:
			w.WriteHeader(http.StatusForbidden)
			return
		case auth.Errors_UserAlreadyRegistered:
			body := handlerErrors.Error{
				ErrorMessage: "user already registered",
				ErrorCode:    http.StatusBadRequest,
			}
			rawDto, _ := json.Marshal(body)
			_, err = w.Write(rawDto)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		case auth.Errors_NoAuthDataProvided:
			body := handlerErrors.Error{
				ErrorMessage: "no auth data",
				ErrorCode:    http.StatusBadRequest,
			}
			rawDto, _ := json.Marshal(body)
			_, err = w.Write(rawDto)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		case auth.Errors_IncorrectUser:
			body := handlerErrors.Error{
				ErrorMessage: "user already registered",
				ErrorCode:    http.StatusBadRequest,
			}
			rawDto, _ := json.Marshal(body)
			_, err = w.Write(rawDto)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		default:
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}

	if !response.Success {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profileDto := &profile.CreateProfileIn{
		Profile: &profile.ProfileData{
			UserID:      *response.UserId,
			UserSurname: inputData.UserSurname,
			UserMail:    inputData.UserMail,
			UserName:    inputData.UserName,
			UserRole:    inputData.UserRole,
		},
	}

	profileResponse, err := h.profileServiceClient.CreateProfile(context.Background(), profileDto)
	if err != nil {
		h.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !profileResponse.Success || profileResponse.Error != nil {
		h.log.Error(profileResponse.Error)
		switch *profileResponse.Error {
		case profile.Errors_ProfileDataNotProvided:
			w.WriteHeader(http.StatusBadRequest)
			outputErr := handlerErrors.Error{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "incorrect username or password",
			}
			bytes, _ := json.Marshal(outputErr)
			_, _ = w.Write(bytes)
		case profile.Errors_ProfileAlreadyExists:
			w.WriteHeader(http.StatusBadRequest)
			outputErr := handlerErrors.Error{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "profile already exists",
			}
			bytes, _ := json.Marshal(outputErr)
			_, _ = w.Write(bytes)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	outDto := handlersDto.RegisterOutDto{
		UserID:      *response.UserId,
		Username:    profileResponse.Profile.UserName,
		UserMail:    profileResponse.Profile.UserMail,
		UserRole:    profileResponse.Profile.UserRole,
		UserSurname: profileResponse.Profile.UserSurname,
	}

	rawOut, _ := json.Marshal(outDto)
	_, _ = w.Write(rawOut)
}
