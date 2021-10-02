package httputils

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	noAuthResponse = `{"error":"401 unauthorized: Basic authentication is required"}`

	invalidAuthResponse = `{"error":"401 unauthorized: username or password is incorrect"}`
)

func setupMockRouter(authMiddleware gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(authMiddleware)
	r.GET("/login", func(c *gin.Context) {
		c.String(200, "OK")
	})

	return r
}

func getAccount() map[string]string {
	accounts := make(map[string]string)
	accounts["admin"] = "admin"
	accounts["user"] = "user"
	return accounts
}

func TestBasicAuthRequiredError_Error(t *testing.T) {
	assert.Equal(t, BasicAuthRequiredErrMsg, BasicAuthRequiredError{}.Error())
}

func TestBasicAuthRequired_Success(t *testing.T) {
	router := setupMockRouter(BasicAuthRequired())

	// correct auth
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("admin", "admin")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestBasicAuthRequired_NoAuth(t *testing.T) {
	router := setupMockRouter(BasicAuthRequired())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())
}

func TestBasicAuthRequired_InvalidAuth(t *testing.T) {

	// invalid basic auth
	router := setupMockRouter(BasicAuthRequired())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.Header.Add(AuthorizationHeaderKey, "YWRtaW4=")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())

	// invalid base64
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login", nil)
	req.Header.Add(AuthorizationHeaderKey, "Basic YWRtaW4")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())

	// no : in encoded auth
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login", nil)
	req.Header.Add(AuthorizationHeaderKey, "Basic YWRtaW4=")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())
}

func TestBasicAuth_Success(t *testing.T) {
	router := setupMockRouter(BasicAuth(getAccount()))

	// admin
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("admin", "admin")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())

	// user
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("user", "user")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestBasicAuth_Success_NoAuth(t *testing.T) {
	router := setupMockRouter(BasicAuth(getAccount()))

	// admin
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())
}

func TestBasicAuth_Success_InvalidAuth(t *testing.T) {
	router := setupMockRouter(BasicAuth(getAccount()))

	// admin
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("admin", "")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, invalidAuthResponse, w.Body.String())

	// non-existing user
	// admin
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("admi", "admin")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, invalidAuthResponse, w.Body.String())
}
