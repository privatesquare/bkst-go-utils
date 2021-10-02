package config

import (
	"fmt"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerConfig_Set(t *testing.T) {
	cnf := ServerConfig{
		Protocol: "",
		Host:     "",
		Port:     "",
		LogLevel: "",
		ProxyUrl: "",
	}
	err := cnf.Set()
	assert.NoError(t, err)
	assert.Equal(t, defaultServerProtocol, ServerCnf.Protocol)
	assert.Equal(t, logger.DefaultLogLevel, ServerCnf.LogLevel)

	cnf.Protocol = "htt"
	err = cnf.Set()
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf(invalidServerProtocolErrMsg, cnf.Protocol), err.Error())

	cnf.Protocol = "http"
	cnf.LogLevel = "INF"
	err = cnf.Set()
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf(invalidServerLogLevelErrMsg, cnf.LogLevel), err.Error())
}
