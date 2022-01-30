package user

type User struct {
	Username  *string
	Password  *string
	AuthToken *string
}

type RegisterUser struct {
	Password    string
	UserMail    string
	UserName    *string
	UserSurname *string
	UserRole    *string
}
