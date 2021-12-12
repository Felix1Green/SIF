package interactor

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"

	"github.com/Felix1Green/SIF/backend/AuthService/internal"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/components"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/entities"
)

type AuthInteractor struct {
	sanitizer        *bluemonday.Policy
	tokenStorage     components.TokenStorage
	userStorage      components.UserStorage
	cacheUserStorage components.UserStorage
	salt             string
}

var (
	NoAuthenticationDataProvidedError = fmt.Errorf("no authentication data provided")
	InternalServiceError              = fmt.Errorf("internal service error: some dependencies are not available")
)

func (s *AuthInteractor) createHashPassword(password, salt string) (string, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), 7)
	return string(hashedPassword), err == nil
}

func (s *AuthInteractor) Auth(user *entities.User) (*int64, error) {
	if user.AuthToken != nil {
		internal.SanitizeInput(s.sanitizer, user.AuthToken)
		id, err := s.tokenStorage.Get(*user.AuthToken)
		if err != nil {
			return nil, InternalServiceError
		}
		return &id, nil
	} else if user.Password != nil && user.Username != nil {
		internal.SanitizeInput(s.sanitizer, user.Username, user.Password)
		password, ok := s.createHashPassword(*user.Password, s.salt)
		if !ok {
			return nil, InternalServiceError
		}
		currentUser, err := s.cacheUserStorage.GetUser(*user.Username, password)
		if err != nil {
			currentUser, err = s.userStorage.GetUser(*user.Username, password)
			if err != nil || currentUser.Password == nil || currentUser.Username == nil {
				return nil, err
			}

			_, err = s.cacheUserStorage.CreateUser(*currentUser.Username, *currentUser.Password, currentUser.UserID)
			if err != nil {
				return nil, err
			}
		} else {
			return currentUser.UserID, nil
		}
	}

	return nil, NoAuthenticationDataProvidedError
}
