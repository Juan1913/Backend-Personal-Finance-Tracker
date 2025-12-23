package errors

type EmailAlreadyExistsError struct {
	Email string
}

func (e *EmailAlreadyExistsError) Error() string {
	return "email already exists: " + e.Email
}

type UserNotFoundError struct {
	UserID string
}

func (e *UserNotFoundError) Error() string {
	return "user not found: " + e.UserID
}
