package logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/private-square/bkst-users-api/utils/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	Sink *MemorySink
)

// MemorySink implements zap.Sink by writing all messages to a buffer.
type MemorySink struct {
	*bytes.Buffer
}

// Implement Close and Sync as no-ops to satisfy the interface. The Write
// method is provided by the embedded buffer.

func (s *MemorySink) Close() error { return nil }
func (s *MemorySink) Sync() error  { return nil }

func init() {
	configureMockLogger()
}

func configureMockLogger() {
	// Create a sink instance, and register it with zap for the "memory"
	// protocol.
	Sink = &MemorySink{new(bytes.Buffer)}
	zap.RegisterSink("memory", func(*url.URL) (zap.Sink, error) {
		return Sink, nil
	})

	// Redirect all messages to the MemorySink.
	config := getLoggerConfig()
	config.OutputPaths = []string{"memory://"}

	setLoggerConfig(config)
}

func TestInfo(t *testing.T) {
	msg := "some info message"
	Info(msg)

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"info\""))
	assert.True(t, strings.Contains(output, msg))
}

func TestWarn(t *testing.T) {
	msg := "some warning message"
	Warn(msg)

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"warn\""))
	assert.True(t, strings.Contains(output, msg))
}

func TestError(t *testing.T) {
	msg := "some error message"
	Error(msg, nil)

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"error\""))
	assert.True(t, strings.Contains(output, msg))

	err := errors.NewError(msg)
	Error(msg, err)

	// Assert sink contents
	output = Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"error\""))
	assert.True(t, strings.Contains(output, msg))
}

func TestGinZap(t *testing.T) {
	r := newRouter()

	apiPath := "/test"
	msg := "some message"
	r.GET(apiPath, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, msg)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)
	r.ServeHTTP(w, req)

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"info\""))
	assert.True(t, strings.Contains(output, "\"path\":\"/test\""))
	assert.True(t, strings.Contains(output, "\"status\":200"))
}

func TestGinZapError(t *testing.T) {
	r := newRouter()

	apiPath := "/test"
	msg := "some message"
	r.GET(apiPath, func(ctx *gin.Context) {
		ctx.Errors = append(ctx.Errors, &gin.Error{
			Err: errors.NewError(""),
		})
		ctx.JSON(http.StatusOK, msg)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)
	r.ServeHTTP(w, req)

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"error\""))
	assert.True(t, strings.Contains(output, "\"caller\":\"gin"))
}

func newRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(GinZap())
	r.Use(gin.Recovery())
	return r
}
