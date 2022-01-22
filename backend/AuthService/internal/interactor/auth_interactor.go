package interactor

import (
	"context"
	"crypto/rand"
	"github.com/sirupsen/logrus"
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
	log              *logrus.Logger
	tokenSize        int
	salt             string
}

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewInteractor(
	tokenStorage components.TokenStorage,
	userStorage components.UserStorage,
	cacheUserStorage components.UserStorage,
	log *logrus.Logger,
	tokenSize int,
	salt string,
) *AuthInteractor {
	return &AuthInteractor{
		userStorage:      userStorage,
		tokenStorage:     tokenStorage,
		cacheUserStorage: cacheUserStorage,
		log:              log,
		tokenSize:        tokenSize,
		salt:             salt,
		sanitizer:        bluemonday.NewPolicy(),
	}
}

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
			return nil, internal.UserNotFoundError
		}

		requestedUser = &entities.User{
			UserID:    &id,
			AuthToken: user.AuthToken,
		}
		return requestedUser, nil
	} else if user.Password != nil && user.Username != nil {
		internal.SanitizeInput(s.sanitizer, user.Username, user.Password)

		currentUser, err := s.cacheUserStorage.GetUser(*user.Username)

		if err != nil {
			currentUser, err = s.userStorage.GetUser(*user.Username)
			if err != nil {
				return nil, err
			}

			passwordErr := bcrypt.CompareHashAndPassword([]byte(*currentUser.Password), []byte(*user.Password+s.salt))
			if passwordErr != nil {
				return nil, internal.UserNotFoundError
			}

			_, cacheErr := s.cacheUserStorage.CreateUser(*currentUser.Username, *currentUser.Password, currentUser.UserID)
			if cacheErr != nil {
				s.log.Warning(context.Background(), cacheErr)
			}

			requestedUser = &entities.User{
				Username: currentUser.Username,
				UserID:   currentUser.UserID,
			}
			return requestedUser, nil

		} else {
			return currentUser, nil
		}
	}

	return nil, internal.NoAuthenticationDataProvidedError
}

func (s *AuthInteractor) Logout(token string) error {
	return s.tokenStorage.Del(token)
}

func (s *AuthInteractor) Register(user *entities.User) (*entities.User, error) {
	internal.SanitizeInput(s.sanitizer, user.Username, user.Password)

	if user.Username != nil && user.Password != nil {
		password, ok := s.createHashPassword(*user.Password)
		if !ok {
			s.log.Error("cannot create hash password")
			return nil, internal.InternalServiceError
		}

		_, err := s.cacheUserStorage.GetUser(*user.Username)
		if err == nil {
			s.log.Error(err)
			return nil, internal.UserAlreadyRegistered
		}

		user, err = s.userStorage.CreateUser(*user.Username, password, nil)
		if err != nil {
			s.log.Error(err)
			return nil, err
		}

		_, cacheErr := s.cacheUserStorage.CreateUser(*user.Username, password, user.UserID)
		if cacheErr != nil {
			s.log.Warning(context.Background(), cacheErr)
		}

		authToken := s.createAuthToken(s.tokenSize)
		err = s.tokenStorage.Set(authToken, *user.UserID)
		if err != nil {
			s.log.Error(err)
			return nil, internal.InternalServiceError
		}

		requestedUser := &entities.User{
			Username:  user.Username,
			AuthToken: &authToken,
			UserID:    user.UserID,
		}
		return requestedUser, nil
	}

	return nil, internal.NoAuthenticationDataProvidedError
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
