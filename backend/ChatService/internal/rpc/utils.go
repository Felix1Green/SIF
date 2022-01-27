package rpc

import "fmt"

var(
	CookieName = "sif_token"
	NoUserCredentialsProvided = fmt.Errorf("no user credentials provided")
	IncorrectUserCredentials = fmt.Errorf("incorrect username or password")
	InternalServiceError = fmt.Errorf("internal service error")
)
