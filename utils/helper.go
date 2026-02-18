package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type APIResponse[T any] struct {
	Successful bool   `json:"successful"`
	ErrorCode  string `json:"error_code"`
	Data       any    `json:"data"`
}

func NewSuccessResponse[T any](data T) APIResponse[T] {
	return APIResponse[T]{
		Successful: true,
		ErrorCode:  "",
		Data:       data,
	}
}

func NewErrorResponse(status int) APIResponse[any] {
	return APIResponse[any]{
		Successful: false,
		ErrorCode:  strconv.Itoa(status),
		Data:       nil,
	}
}

// ParseError returns the appropriate HTTP status code based on the error type
func ParseError(err error) int {
	if err == nil {
		return http.StatusOK
	}

	// 1. GORM / Database Errors
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return http.StatusConflict
	}

	// 2. Validation / Binding Errors
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return http.StatusBadRequest
	}

	// 3. Known HTTP errors (if any)
	if httpErr, ok := err.(interface{ StatusCode() int }); ok {
		return httpErr.StatusCode()
	}

	// Default to 500
	return http.StatusInternalServerError
}
