package user

type User struct {
	Username, Password, AuthToken *string
}

type RegisterUser struct{
	Username, Password string
}