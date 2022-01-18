package entities

type Profile struct{
	UserID int64 `db:"user_id"`
	UserMail string `db:"user_mail"`
	UserName string `db:"username"`
	UserSurname string `db:"user_surname"`
	UserRole string `db:"user_role"`
}
