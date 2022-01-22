package user

type User struct {
	Username, Password, AuthToken *string
}

type RegisterUser struct {
	Password, UserMail              string
	UserName, UserSurname, UserRole *string
}
