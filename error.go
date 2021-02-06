package main

import (
	"fmt"
	"net/http"
)

const (
	ErrorCodePermission     = 1002
	ErrorCodeSessionExpired = 1003
	ErrorCodeNoPlayLink     = 1004
	ErrorCodeBadRequest     = 1005
	ErrorCodeBackend        = 1006
	ErrorCodeNotFound       = 2001
	ErrorCodeInternalServer = 3001
	ErrorCodeTokenNotFound  = 4001
	ErrorCodeTokenExpired   = 4002
	SuccessCode             = 0
)

type (
	ServiceError struct {
		Code           int    `json:"errorCode"`
		Message        string `json:"message"`
		HTTPStatusCode int    `json:"-"`
	}

	ServiceSuccess struct {
		Code    int    `json:"resultCode"`
		Message string `json:"message"`
	}
)

var (
	successResponse = ServiceSuccess{
		Message: "success",
	}

	errBlank = ServiceError{
		Message:        "",
		HTTPStatusCode: http.StatusBadRequest,
	}

	errBadRequest = ServiceError{
		Message:        "bad request",
		HTTPStatusCode: http.StatusBadRequest,
	}

	errNotFound = ServiceError{
		Message:        "not found",
		Code:           ErrorCodeNotFound,
		HTTPStatusCode: http.StatusBadRequest,
	}

	errInternalServer = ServiceError{
		Message:        "something went wrong",
		Code:           ErrorCodeInternalServer,
		HTTPStatusCode: http.StatusInternalServerError,
	}
)

func (re *ServiceError) Err(errorCode int, message string) ResponseData {
	return ResponseData{
		Code:    errorCode,
		Message: fmt.Sprintf("%s: %s", re.Message, message),
	}
}

func (r *ServiceSuccess) Msg(message string) ResponseData {
	return ResponseData{
		Code:    SuccessCode,
		Message: fmt.Sprintf("%s: %s", r.Message, message),
	}
}
