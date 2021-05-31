package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	msg := "some message"
	assert.EqualError(t, NewError(msg), msg)
}

func TestBadRequestError(t *testing.T) {
	msg := "some message"
	err := BadRequestError(msg)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestUnauthorizedError(t *testing.T) {
	msg := "some message"
	err := UnauthorizedError(msg)
	assert.Equal(t, http.StatusUnauthorized, err.Status)
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestForbiddenError(t *testing.T) {
	msg := "some message"
	err := ForbiddenError(msg)
	assert.Equal(t, http.StatusForbidden, err.Status)
	assert.Equal(t, http.StatusText(http.StatusForbidden), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestNotFoundError(t *testing.T) {
	msg := "some message"
	err := NotFoundError(msg)
	assert.Equal(t, http.StatusNotFound, err.Status)
	assert.Equal(t, http.StatusText(http.StatusNotFound), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestInternalServerError(t *testing.T) {
	msg := "some message"
	err := InternalServerError(msg)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestMissingMandatoryParamError(t *testing.T) {
	missingParams := []string{"username", "password"}
	err := MissingMandatoryParamError(missingParams)
	assert.Equal(t, fmt.Sprintf(missingMandatoryParamErrMsg, missingParams), err.Error())
}

func TestPasswordEncryptionError(t *testing.T) {
	err := errors.New("some error message")
	assert.Equal(t, fmt.Sprintf(passwordEncryptionErrMsg, err), PasswordEncryptionError{Err: err}.Error())
}

func TestPasswordDecryptionError_Error(t *testing.T) {
	err := errors.New("some error message")
	assert.Equal(t, fmt.Sprintf(passwordDecryptionErrMsg, err), PasswordDecryptionError{Err: err}.Error())
}