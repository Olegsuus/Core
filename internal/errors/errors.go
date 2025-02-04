package errors

type AppError struct {
	BusinessError string
	UserError     string
	Status        int
}

func (a AppError) Error() string {
	return a.UserError
}
