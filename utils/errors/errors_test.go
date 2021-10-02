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
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestUnauthorizedError(t *testing.T) {
	msg := "some message"
	err := UnauthorizedError(msg)
	assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestForbiddenError(t *testing.T) {
	msg := "some message"
	err := ForbiddenError(msg)
	assert.Equal(t, http.StatusForbidden, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusForbidden), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestNotFoundError(t *testing.T) {
	msg := "some message"
	err := NotFoundError(msg)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusNotFound), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestConflictError(t *testing.T) {
	msg := "some message"
	err := ConflictError(msg)
	assert.Equal(t, http.StatusConflict, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusConflict), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestInternalServerError(t *testing.T) {
	err := InternalServerError()
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), err.Error)
	assert.Equal(t, internalServerErrMsg, err.Message)
}

func TestMissingMandatoryParamError(t *testing.T) {
	missingParams := []string{"username", "password"}
	err := MissingMandatoryParamError(missingParams)
	assert.Equal(t, fmt.Sprintf(MissingMandatoryParamErrMsg, missingParams), err.Error())
}

func TestMissingStartupConfigurationError(t *testing.T) {
	missingParams := []string{"username", "password"}
	err := MissingStartupConfigurationError(missingParams)
	assert.Equal(t, fmt.Sprintf(missingStartupConfigurationErrMsg, missingParams), err.Error())
}

func TestPasswordEncryptionError(t *testing.T) {
	err := errors.New("some error message")
	assert.Equal(t, fmt.Sprintf(passwordEncryptionErrMsg, err), PasswordEncryptionError{Err: err}.Error())
}

func TestPasswordDecryptionError_Error(t *testing.T) {
	err := errors.New("some error message")
	assert.Equal(t, fmt.Sprintf(passwordDecryptionErrMsg, err), PasswordDecryptionError{Err: err}.Error())
}

func TestMissingMandatoryParamError_Error(t *testing.T) {
	err := MissingMandatoryParamError{}
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Sprintf(MissingMandatoryParamErrMsg, []string{}), err.Error())
	}
}

func TestRegexCompileError_Error(t *testing.T) {
	errMsg := fmt.Errorf("invalid regex")
	err := RegexCompileError{errMsg}
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Sprintf(regexCompileErrMsg, errMsg), err.Error())
	}
}

func TestJSONMarshalError_Error(t *testing.T) {
	errMsg := fmt.Errorf("invalid json format")
	err := JSONMarshalError{errMsg}
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Sprintf(jsonMarshalErrMsg, errMsg), err.Error())
	}
}

func TestJSONUnMarshalError_Error(t *testing.T) {
	errMsg := fmt.Errorf("invalid json format")
	err := JSONUnMarshalError{errMsg}
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Sprintf(jsonUnMarshalErrMsg, errMsg), err.Error())
	}
}

func TestYAMLMarshalError_Error(t *testing.T) {
	errMsg := fmt.Errorf("invalid yaml format")
	err := YAMLMarshalError{errMsg}
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Sprintf(yamlMarshalErrMsg, errMsg), err.Error())
	}
}

func TestYAMLUnMarshalError_Error(t *testing.T) {
	errMsg := fmt.Errorf("invalid yaml format")
	err := YAMLUnMarshalError{errMsg}
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Sprintf(yamlUnmarshalErrMsg, errMsg), err.Error())
	}
}
