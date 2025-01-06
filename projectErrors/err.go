package projectErrors

import "fmt"

type UserError struct {
	Code    int
	Message string
	Context string
}

func (e UserError) Error() string {
	return fmt.Sprintf("The %v error occured, details: %v \n Subdetails: %v", e.Code, e.Message, e.Context)
}

var (
	// errors
	ErrInvalidName = UserError{Code: 1, Message: "Incorrect name", Context: "Name must be shorter then 10 and longer than 2 symbols!"}
)
