package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SerivceError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type HTTPResponse struct {
	Error *SerivceError `json:"error,omitempty"`
	Data  interface{}   `json:"data,omitempty"`
}

func mapHTTPtoGRPCStatus(status int) codes.Code {
	switch status {
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusTooManyRequests:
	case http.StatusBadGateway:
	case http.StatusServiceUnavailable:
	case http.StatusGatewayTimeout:
		return codes.Unavailable
	}
	return codes.Unknown
}

func (err *SerivceError) ToGRPCError() error {
	errMessage := ""
	if err.Error != nil {
		errMessage += fmt.Sprintf("error: %v", err.Error)

	}
	if err.Message != "" {
		errMessage = fmt.Sprintf("message: %s\n", err.Message) + errMessage
	}

	return status.Error(mapHTTPtoGRPCStatus(err.Code), errMessage)
}

func (err *SerivceError) ToHTTPError(ctx echo.Context) error {
	sErr := &SerivceError{Code: err.Code, Message: err.Message}
	if err.Error != nil {
		sErr.Error = fmt.Sprint(err.Error)
	}
	return ctx.JSON(err.Code, HTTPResponse{Error: sErr})
}
