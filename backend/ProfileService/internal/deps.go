package internal

import "fmt"

var (
	ProfileNotFoundError        = fmt.Errorf("profile not found")
	ProfileAlreadyExists        = fmt.Errorf("profile already exists")
	ServiceInternalError        = fmt.Errorf("internal service error")
	ProfileDataNotProvidedError = fmt.Errorf("profile data not provided")
)
