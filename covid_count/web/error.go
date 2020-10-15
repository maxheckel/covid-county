package web

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	ErrorTypeNotFound           ErrorType = "not-found"
	ErrorTypeUnexpected         ErrorType = "unexpected-error"
	ErrorTypeNotImplemented     ErrorType = "not-implemented"
	ErrorTypeBadRequest         ErrorType = "bad-request"
	ErrorTypeAuth               ErrorType = "auth-error"
	ErrorTypeConflict           ErrorType = "conflict"
	ErrorTypeMethodNotAllowed   ErrorType = "method-not-allowed"
	ErrorTypeServiceUnavailable ErrorType = "server-unavailable"
)

// Type represents the body of an error response (e.g. non-2xx statuses)
type Error struct {
	Type     ErrorType `json:"error"`
	Code     int       `json:"code"`
	Messages []string  `json:"messages"`
}

type ErrorType string

func (e *Error) Error() string {
	return fmt.Sprintf(`[type=%s code=%d messages=[%s]]`, e.Type, e.Code, strings.Join(e.Messages, "; "))
}

// UnexpectedError returns an Error that represents a "not found" error
func UnexpectedError(messages ...string) *Error {
	return &Error{
		Type:     ErrorTypeUnexpected,
		Code:     http.StatusInternalServerError,
		Messages: messages,
	}
}

// NotImplemented returns an Error that represents an attempt to call an
// endpoint that has not yet been implemented
func NotImplemented() *Error {
	return &Error{
		Type: ErrorTypeNotImplemented,
		Code: http.StatusNotImplemented,
		Messages: []string{
			"This endpoint has not been implemented",
		},
	}
}

// MethodNotAllowedError returns an Error that represents an attempt to call an existing path w/ an http method
// that has not been implemented
func MethodNotAllowedError() *Error {
	return &Error{
		Type: ErrorTypeMethodNotAllowed,
		Code: http.StatusMethodNotAllowed,
		Messages: []string{
			"this http method is not allowed",
		},
	}
}

// NotFound returns an Error that represents an unexpected app error
func NotFound(messages ...string) *Error {
	return &Error{
		Type:     ErrorTypeNotFound,
		Code:     http.StatusNotFound,
		Messages: messages,
	}
}

// Unauthorized returns an Error that represents an authentication issue
func Unauthorized(messages ...string) *Error {
	return newAuthError(messages, http.StatusUnauthorized)
}

// Forbidden returns an Error that represents an authorization issue
func Forbidden(messages ...string) *Error {
	return newAuthError(messages, http.StatusForbidden)
}

// BadRequest returns an Error that represents an invalid request
func BadRequest(messages ...string) *Error {
	return &Error{
		Type:     ErrorTypeBadRequest,
		Code:     http.StatusBadRequest,
		Messages: messages,
	}
}

// Conflict returns an Error that represents a conflict
func Conflict(messages ...string) *Error {
	return &Error{
		Type:     ErrorTypeConflict,
		Code:     http.StatusConflict,
		Messages: messages,
	}
}

func newAuthError(messages []string, statusCode int) *Error {
	return &Error{
		Type:     ErrorTypeAuth,
		Code:     statusCode,
		Messages: messages,
	}
}

// ServiceUnavailable returns an Error that represents as service unavailable request
func ServiceUnavailable(messages ...string) *Error {
	return &Error{
		Type:     ErrorTypeServiceUnavailable,
		Code:     http.StatusServiceUnavailable,
		Messages: messages,
	}
}
