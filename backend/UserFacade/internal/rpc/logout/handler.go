package logout

import (
	"UserFacade/internal/generated/clients/auth"
	"UserFacade/internal/rpc"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	authServiceClient auth.AuthClient
	log               *logrus.Logger
}

func NewHandler(authServiceClient auth.AuthClient, log *logrus.Logger) *handler {
	return &handler{
		authServiceClient: authServiceClient,
		log:               log,
	}
}

func getSessionCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}

func deleteSessionCookie(w http.ResponseWriter, name string) {
	cookie := http.Cookie{
		Name:   name,
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}

// Handle godoc
// @Summary log out
// @Description log out user
// @ID logout-id
// @Success 200
// @Failure 503 {object} handlerErrors.Error
// @Router /logout [post]
func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	token := getSessionCookie(r, rpc.CookieName)
	logoutIn := auth.LogoutIn{
		AuthToken: token,
	}

	response, err := h.authServiceClient.LogOut(context.Background(), &logoutIn)
	if err != nil || response == nil {
		h.log.Warning(fmt.Sprintf("auth service handled with error: %e", err))
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	deleteSessionCookie(w, rpc.CookieName)
	w.WriteHeader(http.StatusOK)
	return
}
