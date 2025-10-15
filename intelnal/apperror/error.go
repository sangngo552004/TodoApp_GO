package apperror

import "net/http"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func BadRequest(message string, err error) *AppError {
	return New(http.StatusBadRequest, message, err)
}

func InternalServerError(message string, err error) *AppError {
	return New(http.StatusInternalServerError, message, err)
}

func Unauthorized(message string, err error) *AppError {
	return New(http.StatusUnauthorized, message, err)
}

func Forbidden(message string, err error) *AppError {
	return New(http.StatusForbidden, message, err)
}

func NotFound(message string, err error) *AppError {
	return New(http.StatusNotFound, message, err)
}

func Conflict(message string, err error) *AppError {
	return New(http.StatusConflict, message, err)
}

func UnprocessableEntity(message string, err error) *AppError {
	return New(http.StatusUnprocessableEntity, message, err)
}

func NotImplemented(message string, err error) *AppError {
	return New(http.StatusNotImplemented, message, err)
}

func ServiceUnavailable(message string, err error) *AppError {
	return New(http.StatusServiceUnavailable, message, err)
}
