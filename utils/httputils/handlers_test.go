package httputils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	healthApiPath = "/health"
	notFoundApiPath = "/notFound"
)

var (
	successResponse = `{"message":"OK"}`

	pathNotfoundResponse = `{"message":"Path Not Found"}`

	methodNotAllowedResponse = `{"message":"Method Not Allowed"}`
)

func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET(healthApiPath, Health)
	r.NoRoute(NoRoute)
	r.HandleMethodNotAllowed = true
	r.NoMethod(MethodNotAllowed)
	return r
}

func readResponseBody(body *bytes.Buffer, t *testing.T) string {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		t.Error("unexpected test utils error")
	}
	return string(byt)
}

func TestHealth(t *testing.T) {
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, healthApiPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, successResponse, readResponseBody(w.Body, t))
}

func TestNoRoute(t *testing.T) {
	r := newRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, notFoundApiPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, pathNotfoundResponse, readResponseBody(w.Body, t))
}

func TestMethodNotAllowed(t *testing.T) {
	r := newRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, healthApiPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedResponse, readResponseBody(w.Body, t))
}

