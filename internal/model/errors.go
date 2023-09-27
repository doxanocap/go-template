package model

import (
	"net/http"

	"github.com/doxanocap/pkg/errs"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// default http error responses
var (
	HttpInternalServerError = errs.NewHttp(http.StatusInternalServerError, "internal server error")
	HttpConflictError = errs.NewHttp(http.StatusConflict, "conflict")
	HttpBadRequest = errs.NewHttp(http.StatusBadRequest, "bad request")
	HttpUnauthorized = errs.NewHttp(http.StatusUnauthorized, "unauthorized")
)

// custom errors for special cases
var (

	
	ErrUserAlreadyExist = errs.NewHttp(http.StatusConflict, "user already exist")
	ErrUserNotFound = errs.NewHttp(http.StatusConflict, "user not found")

	ErrTokenNotFound = errs.NewHttp(http.StatusConflict, "token not found")
	ErrInvalidToken = errs.NewHttp(http.StatusConflict, "invalid token")

	ErrInvalidFileFormat       = errs.NewHttp(http.StatusConflict, "invalid file formate")
	ErrSuchServiceAlreadyExist = errs.NewHttp(http.StatusConflict, "service with such name already exists")
	ErrServiceIdNotFound       = errs.NewHttp(http.StatusNotFound, "service with such id not found")
)
