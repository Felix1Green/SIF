package login

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/generated/clients/auth"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/models/handlerErrors"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/models/user"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/rpc"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type handler struct {
	authServiceClient auth.AuthClient
	log               *logrus.Logger
}

func NewLoginHandler(authServiceClient auth.AuthClient, logger *logrus.Logger) *handler {
	return &handler{
		authServiceClient: authServiceClient,
		log:               logger,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.HandlePostRequest(w, r)
		return
	case http.MethodGet:
		h.HandlerGetRequest(w, r)
		return
	}
}

func (h *handler) HandlerGetRequest(w http.ResponseWriter, r *http.Request) {
	authToken := ""
	if val, err := r.Cookie(rpc.CookieName); err == nil && val != nil {
		authToken = val.Value
	} else {
		h.log.Error("no auth token found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authDto := &auth.AuthIn{
		AuthToken: &authToken,
	}

	result, err := h.authServiceClient.Auth(context.Background(), authDto)
	if err != nil {
		errInfo := handlerErrors.AuthError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "no authentication data provided",
		}
		bytesError, _ := json.Marshal(errInfo)
		h.log.Info(errInfo)
		_, _ = w.Write(bytesError)
		return
	}

	if result.Error != nil {
		switch *result.Error {
		case auth.Errors_IncorrectUser:
			w.WriteHeader(http.StatusUnauthorized)
			outputErr := handlerErrors.AuthError{
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: "incorrect username or password",
			}
			bytes, _ := json.Marshal(outputErr)
			_, err = w.Write(bytes)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			return
		default:
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}

	if result.Success {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (h *handler) HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	inputCredentials := user.User{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &inputCredentials)

	if val, err := r.Cookie(rpc.CookieName); err != nil && val != nil {
		inputCredentials.AuthToken = &val.Value
	}

	if inputCredentials.Username == nil && inputCredentials.AuthToken == nil {
		w.WriteHeader(http.StatusBadRequest)
		errInfo := handlerErrors.AuthError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "no authentication data provided",
		}
		bytesError, _ := json.Marshal(errInfo)
		h.log.Info(errInfo)
		_, _ = w.Write(bytesError)
		return
	}

	result, err := h.authServiceClient.Auth(context.Background(), &auth.AuthIn{
		AuthToken: inputCredentials.AuthToken,
		Username:  inputCredentials.Username,
		Password:  inputCredentials.Password,
	})

	if err != nil {
		h.log.Warning(fmt.Sprintf("auth service returned error: %s", err))
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if result.Error != nil {
		switch *result.Error {
		case auth.Errors_IncorrectUser:
			outputErr := handlerErrors.AuthError{
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: "incorrect username or password",
			}
			bytes, _ := json.Marshal(outputErr)
			_, err = w.Write(bytes)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}

	if result.Success {
		cookie := &http.Cookie{
			Name:     rpc.CookieName,
			Value:    *result.UserToken,
			Expires:  time.Now().AddDate(rpc.CookieExpiresYear, rpc.CookieExpiresMonth, rpc.CookieExpiresDay),
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}
