package errors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	MissingMandatoryParamErrMsg = "Missing mandatory parameter(s) : %v"
	internalServerErrMsg        = "Unable to process the request due to an internal error. Please contact the system administrator"
	invalidPasswordErrMsg       = "password should be at least 8 characters long with at least one number, one uppercase letter, one lowercase letter and one special character"
	passwordEncryptionErrMsg    = "password encryption error: %v"
	passwordDecryptionErrMsg    = "password decryption error: %v"
	regexCompileErrMsg          = "Unable to compile regex : %v"
	jsonMarshalErrMsg           = "JSON marshal error : %v"
	jsonUnMarshalErrMsg         = "JSON unmarshal error : %v"
	yamlMarshalErrMsg           = "YAML marshal error : %v"
	yamlUnmarshalErrMsg         = "YAML unmarshal error : %v"
)

var (
	InvalidPasswordError = errors.New(invalidPasswordErrMsg)
)

type RestErr struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Error      string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      http.StatusText(http.StatusBadRequest),
	}
}

func BadRequestErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusBadRequest,
		Error:      http.StatusText(http.StatusBadRequest),
	}
}

func UnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Error:      http.StatusText(http.StatusUnauthorized),
	}
}

func UnauthorizedErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusUnauthorized,
		Error:      http.StatusText(http.StatusUnauthorized),
	}
}

func ForbiddenError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusForbidden,
		Error:      http.StatusText(http.StatusForbidden),
	}
}

func ForbiddenErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusForbidden,
		Error:      http.StatusText(http.StatusForbidden),
	}
}

func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      http.StatusText(http.StatusNotFound),
	}
}

func NotFoundErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusNotFound,
		Error:      http.StatusText(http.StatusNotFound),
	}
}

func ConflictError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusConflict,
		Error:      http.StatusText(http.StatusConflict),
	}
}

func ConflictErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusConflict,
		Error:      http.StatusText(http.StatusConflict),
	}
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message:    internalServerErrMsg,
		StatusCode: http.StatusInternalServerError,
		Error:      message,
	}
}

func InternalServerErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    internalServerErrMsg,
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Sprintf(format, a...),
	}
}

// MissingMandatoryParamError represents an error when a mandatory parameter is missing
type MissingMandatoryParamError []string

// Error returns the formatted MissingMandatoryParamError
func (e MissingMandatoryParamError) Error() string {
	return fmt.Sprintf(MissingMandatoryParamErrMsg, []string(e))
}

type PasswordEncryptionError struct {
	Err error
}

func (e PasswordEncryptionError) Error() string {
	return fmt.Sprintf(passwordEncryptionErrMsg, e.Err)
}

type PasswordDecryptionError struct {
	Err error
}

func (e PasswordDecryptionError) Error() string {
	return fmt.Sprintf(passwordDecryptionErrMsg, e.Err)
}

// RegexCompileError represents an error when a regex compilation fails
type RegexCompileError struct {
	Err error
}

// Error returns the formatted RegexCompileError
func (rc RegexCompileError) Error() string {
	return fmt.Sprintf(regexCompileErrMsg, rc.Err)
}

// JSONMarshalError represents an error when json marshal fails
type JSONMarshalError struct {
	Err error
}

// Error returns the formatted JSONMarshalError
func (jm JSONMarshalError) Error() string {
	return fmt.Sprintf(jsonMarshalErrMsg, jm.Err)
}

// JSONUnMarshalError represents an error when json unmarshal fails
type JSONUnMarshalError struct {
	Err error
}

// Error returns the formatted JSONUnMarshalError
func (jum JSONUnMarshalError) Error() string {
	return fmt.Sprintf(jsonUnMarshalErrMsg, jum.Err)
}

// YAMLMarshalError represents an error when yaml marshal fails
type YAMLMarshalError struct {
	Err error
}

// Error returns the formatted YAMLMarshalError
func (ym YAMLMarshalError) Error() string {
	return fmt.Sprintf(yamlMarshalErrMsg, ym.Err)
}

// YAMLUnMarshalError represents an error when yaml unmarshal fails
type YAMLUnMarshalError struct {
	Err error
}

// Error returns the formatted YAMLUnMarshalError
func (yum YAMLUnMarshalError) Error() string {
	return fmt.Sprintf(yamlUnmarshalErrMsg, yum.Err)
}
