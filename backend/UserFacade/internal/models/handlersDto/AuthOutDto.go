package handlersDto

type AuthOutDto struct {
	UserID      int64
	Username    *string
	UserMail    string
	UserSurname *string
	UserRole    *string
}

type RegisterOutDto struct {
	UserID      int64
	Username    *string
	UserMail    string
	UserSurname *string
	UserRole    *string
}
