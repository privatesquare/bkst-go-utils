package config

import (
	"fmt"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-go-utils/utils/slice"
)

var (
	ServerCnf            = ServerConfig{}
	validServerProtocols = []string{"http", "https"}
	validServerLoglevels = []string{"INFO", "DEBUG"}
)

const (
	defaultServerProtocol = "https"

	invalidServerProtocolErrMsg = "Invalid Server HTTP protocol : %s"
	invalidServerLogLevelErrMsg = "Invalid Server Log Level : %s"
)

// ServerConfig represents the required configuration to run a http server.
type ServerConfig struct {
	Protocol string `mapstructure:"SERVER_PROTOCOL"`
	Host     string `mapstructure:"SERVER_HOST" required:"true"`
	Port     string `mapstructure:"SERVER_PORT" required:"true"`
	LogLevel string `mapstructure:"SERVER_LOG_LEVEL"`
	ProxyUrl string `mapstructure:"SERVER_PROXY_URL"`
}

// Set sets the server configuration in the global variable ServerCnf.
func (cnf *ServerConfig) Set() error {
	ServerCnf.Protocol = cnf.Protocol
	if err := ServerCnf.validateServerProtocol(); err != nil {
		return err
	}

	ServerCnf.Host = cnf.Host
	ServerCnf.Port = cnf.Port

	ServerCnf.LogLevel = cnf.LogLevel
	if err := ServerCnf.validateServerLogLevel(); err != nil {
		return err
	}
	logger.SetLoggerConfig(logger.GetLoggerConfig(ServerCnf.LogLevel))

	ServerCnf.ProxyUrl = cnf.ProxyUrl

	return nil
}

// validateServerProtocol checks if the server protocol is set to a valid value.
// if the protocol is "", the value is set to a default value (https).
// The method returns an error if the server protocol is not valid.
func (cnf *ServerConfig) validateServerProtocol() error {
	if ServerCnf.Protocol == "" {
		ServerCnf.Protocol = defaultServerProtocol
	}
	if !slice.EntryExists(validServerProtocols, cnf.Protocol) {
		return errors.New(fmt.Sprintf(invalidServerProtocolErrMsg, cnf.Protocol))
	}
	return nil
}

// validateServerLogLevel checks if the server log level is valid.
// If the log level is "", the value is set to a default value (INFO).
func (cnf *ServerConfig) validateServerLogLevel() error {
	if ServerCnf.LogLevel == "" {
		ServerCnf.LogLevel = logger.DefaultLogLevel
	}
	if !slice.EntryExists(validServerLoglevels, cnf.LogLevel) {
		return errors.New(fmt.Sprintf(invalidServerLogLevelErrMsg, cnf.LogLevel))
	}
	return nil
}
