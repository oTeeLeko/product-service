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

func ParseError(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return http.StatusConflict
	}

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return http.StatusBadRequest
	}

	if httpErr, ok := err.(interface{ StatusCode() int }); ok {
		return httpErr.StatusCode()
	}

	return http.StatusInternalServerError
}
