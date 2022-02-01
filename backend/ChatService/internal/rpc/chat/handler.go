package chat

import (
	"ChatService/internal/entities/handler_errors"
	"ChatService/internal/generated/clients/auth"
	"ChatService/internal/rpc"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	log               *logrus.Logger
	authServiceClient auth.AuthClient
}

func NewHandler(logger *logrus.Logger, authClient auth.AuthClient) *handler {
	return &handler{
		log:               logger,
		authServiceClient: authClient,
	}
}

var connUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	_, err := h.getUserIDFromToken(r)
	if err != nil {
		var (
			statusCode = 200
			outErr     handler_errors.Error
		)

		switch err {
		case rpc.NoUserCredentialsProvided, rpc.IncorrectUserCredentials:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusServiceUnavailable
		}

		outErr = handler_errors.Error{
			ErrorMessage: err.Error(),
			ErrorCode:    statusCode,
		}

		rawErr, _ := json.Marshal(outErr)
		w.WriteHeader(statusCode)
		_, _ = w.Write(rawErr)
		return
	}

	_, err = connUpgrade.Upgrade(w, r, nil)
	if err != nil {
		h.log.Errorf("err with upgrading connection: %s", err.Error())
		outErr := handler_errors.Error{
			ErrorCode:    http.StatusServiceUnavailable,
			ErrorMessage: "connection could not be established",
		}
		rawErr, _ := json.Marshal(outErr)
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write(rawErr)
		return
	}

	// send connection peer to interactor
}

func (h *handler) getUserIDFromToken(r *http.Request) (int64, error) {
	authToken := ""
	if val, err := r.Cookie(rpc.CookieName); err == nil && val != nil {
		authToken = val.Value
	} else {
		h.log.Error("no auth token found")
		return 0, rpc.NoUserCredentialsProvided
	}

	authDto := &auth.AuthIn{
		AuthToken: &authToken,
	}

	result, err := h.authServiceClient.Auth(context.Background(), authDto)
	if err != nil {
		return 0, rpc.NoUserCredentialsProvided
	}

	if result.Error != nil || !result.Success {
		switch *result.Error {
		case auth.Errors_IncorrectUser:
			return 0, rpc.IncorrectUserCredentials
		default:
			return 0, rpc.InternalServiceError
		}
	}

	return *result.UserId, nil
}
