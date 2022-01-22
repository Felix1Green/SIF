package register

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/generated/clients/auth"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/generated/clients/profile"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/models/handlerErrors"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/models/handlersDto"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/models/user"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/rpc"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
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
			body := handlerErrors.AuthError{
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
			body := handlerErrors.AuthError{
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
			body := handlerErrors.AuthError{
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

	if !profileResponse.Success {
		switch *profileResponse.Error {
		case profile.Errors_ProfileDataNotProvided:
			w.WriteHeader(http.StatusBadRequest)
			outputErr := handlerErrors.AuthError{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "incorrect username or password",
			}
			bytes, _ := json.Marshal(outputErr)
			_, err = w.Write(bytes)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
		case profile.Errors_ProfileAlreadyExists:
			w.WriteHeader(http.StatusBadRequest)
			outputErr := handlerErrors.AuthError{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "profile already exists",
			}
			bytes, _ := json.Marshal(outputErr)
			_, err = w.Write(bytes)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	outDto := handlersDto.RegisterOutDto{
		UserID:      *response.UserId,
		Username:    profileResponse.Profile.UserName,
		UserMail:    profileResponse.Profile.UserMail,
		UserRole:    profileResponse.Profile.UserRole,
		UserSurname: profileResponse.Profile.UserSurname,
	}

	cookie := &http.Cookie{
		Name:    rpc.CookieName,
		Value:   *response.UserToken,
		Expires: time.Now().AddDate(rpc.CookieExpiresYear, rpc.CookieExpiresMonth, rpc.CookieExpiresDay),
		Path:    "/",
		Secure:  true,
	}
	http.SetCookie(w, cookie)

	rawOut, _ := json.Marshal(outDto)
	_, _ = w.Write(rawOut)
}
