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

var(
	cookieExpiresYear = 0
	cookieExpiresMonth = 1
	cookieExpiresDay = 0
)

type handler struct{
	authServiceClient auth.AuthClient
	log *logrus.Logger
}

func NewLoginHandler(authServiceClient auth.AuthClient, logger *logrus.Logger) *handler {
	return &handler{
		authServiceClient: authServiceClient,
		log: logger,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodPost:
		h.HandlePostRequest(w, r)
		return
	case http.MethodGet:
		return
	}
}

func (h *handler) HandlePostRequest(w http.ResponseWriter, r *http.Request){
	inputCredentials := user.User{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &inputCredentials)

	if val, err := r.Cookie(rpc.CookieName); err != nil && val != nil{
		inputCredentials.AuthToken = &val.Value
	}

	if inputCredentials.Username == nil && inputCredentials.AuthToken == nil{
		w.WriteHeader(http.StatusBadRequest)
		errInfo := handlerErrors.AuthError{
			ErrorCode: http.StatusBadRequest,
			ErrorMessage: "no authentication data provided",
		}
		bytesError, _ := json.Marshal(errInfo)
		h.log.Info(errInfo)
		_, _ = w.Write(bytesError)
		return
	}

	result, err := h.authServiceClient.Auth(context.Background(), &auth.AuthIn{
		AuthToken: inputCredentials.AuthToken,
		Username: inputCredentials.Username,
		Password: inputCredentials.Password,
	})

	if err != nil{
		h.log.Warning(fmt.Sprintf("auth service returned error: %s", err))
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if result.Error != nil{
		switch *result.Error{
		case auth.Errors_IncorrectUser:
			outputErr := handlerErrors.AuthError{
				ErrorCode: http.StatusUnauthorized,
				ErrorMessage: "incorrect username or password",
			}
			bytes, _ := json.Marshal(outputErr)
			_, err = w.Write(bytes)
			if err != nil{
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

	if result.Success{
		cookie := &http.Cookie{
			Name: rpc.CookieName,
			Value: *result.UserToken,
			Expires: time.Now().AddDate(cookieExpiresYear,cookieExpiresMonth, cookieExpiresDay),
			Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}