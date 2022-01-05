package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	msg    = "some message"
	format = "some message format %s"
	arg    = "abc"
)

func TestNewError(t *testing.T) {
	assert.EqualError(t, New(msg), msg)
}

func TestNewErrorf(t *testing.T) {
	assert.EqualError(t, Newf(format, arg), fmt.Sprintf(format, arg))
}

func TestBadRequestError(t *testing.T) {
	err := BadRequestError(msg)
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestBadRequestErrorf(t *testing.T) {
	err := BadRequestErrorf(format, arg)
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusBadRequest), err.Error)
	assert.Equal(t, fmt.Sprintf(format, arg), err.Message)
}

func TestUnauthorizedError(t *testing.T) {
	err := UnauthorizedError(msg)
	assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestUnauthorizedErrorf(t *testing.T) {
	err := UnauthorizedErrorf(format, arg)
	assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), err.Error)
	assert.Equal(t, fmt.Sprintf(format, arg), err.Message)
}

func TestForbiddenError(t *testing.T) {
	err := ForbiddenError(msg)
	assert.Equal(t, http.StatusForbidden, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusForbidden), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestForbiddenErrorf(t *testing.T) {
	err := ForbiddenErrorf(format, arg)
	assert.Equal(t, http.StatusForbidden, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusForbidden), err.Error)
	assert.Equal(t, fmt.Sprintf(format, arg), err.Message)
}

func TestNotFoundError(t *testing.T) {
	err := NotFoundError(msg)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusNotFound), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestNotFoundErrorf(t *testing.T) {
	err := NotFoundErrorf(format, arg)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusNotFound), err.Error)
	assert.Equal(t, fmt.Sprintf(format, arg), err.Message)
}

func TestConflictError(t *testing.T) {
	err := ConflictError(msg)
	assert.Equal(t, http.StatusConflict, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusConflict), err.Error)
	assert.Equal(t, msg, err.Message)
}

func TestConflictErrorf(t *testing.T) {
	err := ConflictErrorf(format, arg)
	assert.Equal(t, http.StatusConflict, err.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusConflict), err.Error)
	assert.Equal(t, fmt.Sprintf(format, arg), err.Message)
}

func TestInternalServerError(t *testing.T) {
	err := InternalServerError(msg)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, msg, err.Error)
	assert.Equal(t, internalServerErrMsg, err.Message)
}

func TestInternalServerErrorf(t *testing.T) {
	err := InternalServerErrorf(format, arg)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, internalServerErrMsg, err.Message)
	assert.Equal(t, fmt.Sprintf(format, arg), err.Error)
}
