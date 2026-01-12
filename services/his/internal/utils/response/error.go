package response

import (
	"net/http"
)

var (
	RecordNotFound = "record not found"
)

type AppError struct {
	HTTPStatus int
	Code       string
	Message    string
	Errors     []FieldError
}

func (e *AppError) Error() string {
	return e.Message
}

// Constructors
func NewAppError(status int, code, message string) *AppError {
	return &AppError{
		HTTPStatus: status,
		Code:       code,
		Message:    message,
	}
}

func BadRequestError(message string) *AppError {
	return &AppError{
		HTTPStatus: http.StatusBadRequest,
		Code:       "BAD_REQUEST",
		Message:    message,
	}
}
func NotFoundError(message string) *AppError {
	return &AppError{
		HTTPStatus: http.StatusNotFound,
		Code:       "NOT_FOUND",
		Message:    message,
	}
}

func UnprocessableEntityError(message string) *AppError {
	return &AppError{
		HTTPStatus: http.StatusUnprocessableEntity,
		Code:       "UNPROCESSABLE_ENTITY",
		Message:    message,
	}
}
