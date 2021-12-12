package interactor

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

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
	log              internal.Logger
	tokenSize        int
	salt             string
}

var (
	NoAuthenticationDataProvidedError = fmt.Errorf("no authentication data provided")
	InternalServiceError              = fmt.Errorf("internal service error: some dependencies are not available")
	letters                           = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func (s *AuthInteractor) Auth(user *entities.User) (requestedUser *entities.User, err error) {
	defer func() {
		if err == nil && requestedUser != nil && requestedUser.AuthToken == nil {
			authToken := s.createAuthToken(s.tokenSize)
			requestedUser.AuthToken = &authToken
			err = s.tokenStorage.Set(authToken, *requestedUser.UserID)
		}
	}()

	if user.AuthToken != nil {
		internal.SanitizeInput(s.sanitizer, user.AuthToken)
		id, err := s.tokenStorage.Get(*user.AuthToken)
		if err != nil {
			s.log.Warning(context.Background(), err)
			return nil, InternalServiceError
		}

		requestedUser = &entities.User{
			UserID:    &id,
			AuthToken: user.AuthToken,
		}
		return requestedUser, err
	} else if user.Password != nil && user.Username != nil {
		internal.SanitizeInput(s.sanitizer, user.Username, user.Password)
		password, ok := s.createHashPassword(*user.Password)
		if !ok {
			return nil, InternalServiceError
		}

		currentUser, err := s.cacheUserStorage.GetUser(*user.Username, password)

		if err != nil {
			currentUser, err = s.userStorage.GetUser(*user.Username, password)
			if err != nil || (currentUser.Password == nil && currentUser.Username == nil && currentUser.UserID == nil) {
				return nil, err
			}

			_, cacheErr := s.cacheUserStorage.CreateUser(*currentUser.Username, *currentUser.Password, currentUser.UserID)
			if cacheErr != nil {
				s.log.Warning(context.Background(), cacheErr)
			}

			requestedUser = &entities.User{
				Username: currentUser.Username,
				Password: currentUser.Password,
				UserID:   currentUser.UserID,
			}
			return requestedUser, err

		} else {
			return currentUser, nil
		}
	}

	return nil, NoAuthenticationDataProvidedError
}

func (s *AuthInteractor) createHashPassword(password string) (string, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+s.salt), 7)
	return string(hashedPassword), err == nil
}

func (s *AuthInteractor) createAuthToken(amount int) string {
	b := make([]rune, amount)
	for i := range b {
		randInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[randInt.Int64()]
	}
	return string(b)
}
