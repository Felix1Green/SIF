package profile_storage

import(
	"fmt"
	"github.com/Felix1Green/SIF/backend/ProfileService/internal/entities"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

type PostgresProfileStorage struct{
	pool *pgx.ConnPool
	log *logrus.Logger
}

var(
	ProfileNotFoundError = fmt.Errorf("profile not found")
	ProfileAlreadyExists = fmt.Errorf("profile already exists")
	InternalServiceError = fmt.Errorf("internal service error")
)

func NewPostgresProfileStorage(pool *pgx.ConnPool, logger *logrus.Logger) *PostgresProfileStorage{
	return &PostgresProfileStorage{
		pool: pool,
		log: logger,
	}
}

func (r *PostgresProfileStorage) CreateProfile(profile *entities.Profile) (*entities.Profile, error){
	if r.pool == nil{
		r.log.Error("no connection pool provided")
		return nil, InternalServiceError
	}

	query := "INSERT INTO Profile (user_id, user_mail, username, user_surname, user_role) VALUES ($1, $2, $3, $4, $5) RETURNING user_id"
	_, err := r.pool.Exec(
		query, profile.UserID, profile.UserMail, profile.UserName, profile.UserSurname, profile.UserRole,
	)

	if err != nil{
		switch err.(type){
		case pgx.PgError:
			r.log.Errorf("postgresql server error: %s", err.Error())
			return nil, InternalServiceError
		default:
			return nil, ProfileAlreadyExists
		}
	}

	return profile, nil
}

func (r *PostgresProfileStorage) GetProfileByID(userID int64) (*entities.Profile, error){
	if r.pool == nil{
		r.log.Error("no connection pool provided")
		return nil, InternalServiceError
	}

	profile := &entities.Profile{}
	query := "SELECT user_id, user_mail, username, user_surname, user_role from Profile where user_id = $1"
	result := r.pool.QueryRow(query, userID)

	err := result.Scan(&profile)
	if err != nil{
		switch err.(type){
		case pgx.PgError:
			r.log.Errorf("postgresql server error: %s", err.Error())
			return nil, InternalServiceError
		default:
			return nil, ProfileNotFoundError
		}
	}

	return profile, nil
}

func (r *PostgresProfileStorage) GetAllProfiles()([]*entities.Profile, error){
	if r.pool == nil{
		r.log.Error("no connection pool provided")
		return nil, InternalServiceError
	}

	profileArray := make([]*entities.Profile, 0)

	query := "SELECT user_id, user_mail, username, user_surname, user_role from Profile"
	result, err := r.pool.Query(query)
	if err != nil{
		r.log.Errorf("postgresql server error: %s", err.Error())
		return nil, InternalServiceError
	}

	for result.Next(){
		profile := &entities.Profile{}
		err = result.Scan(profile)
		if err != nil{
			r.log.Errorf("postgresql server error: %s", err.Error())
			return nil, InternalServiceError
		}

		profileArray = append(profileArray, profile)
	}

	return profileArray, nil
}