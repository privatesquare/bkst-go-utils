package config

import (
	"fmt"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var (
	mockCfgFilePath  = "./config.env"
	mockCfgCreateMsg = "Mock Configuration file %s created"

	url      = "https://test.com"
	username = "test"
	password = "test123"

	cfgFormat = `MOCK_URL=%s
MOCK_USERNAME=%s
MOCK_PASSWORD=%s`

	validCfg      = fmt.Sprintf(cfgFormat, url, username, password)
	emptyCfg      = fmt.Sprintf(cfgFormat, "", "", "")
	incompleteCfg = fmt.Sprintf(cfgFormat, url, username, "")
)

type MockCfg struct {
	Url      string `mapstructure:"MOCK_URL" required:"true"`
	Username string `mapstructure:"MOCK_USERNAME" required:"true"`
	Password string `mapstructure:"MOCK_PASSWORD" required:"true"`
}

func TestAddConfigPath(t *testing.T) {
	AddConfigPath("./config")
	assert.Equal(t, "./config", configPathValue)
}

func TestSetConfigName(t *testing.T) {
	SetConfigName("test")
	assert.Equal(t, "test", configNameValue)
}

func TestSetConfigType(t *testing.T) {
	SetConfigType("yaml")
	assert.Equal(t, "yaml", configTypeValue)
}

func TestLoad(t *testing.T) {
	resetConfig()
	if err := mockConfig(mockCfgFilePath, validCfg); err != nil {
		assert.NoError(t, err)
	}
	t.Logf(mockCfgCreateMsg, mockCfgFilePath)
	defer removeMockCfg(mockCfgFilePath)

	cfg := &MockCfg{}
	err := Load(cfg)
	if assert.NoError(t, err) {
		assert.Equal(t, url, cfg.Url)
		assert.Equal(t, username, cfg.Username)
		assert.Equal(t, password, cfg.Password)
	}
}

func TestLoad_NoCfgFile(t *testing.T) {
	resetConfig()
	cfg := &MockCfg{}
	err := Load(cfg)
	if assert.Error(t, err) {
		assert.True(t, strings.Contains(err.Error(), "Config File \"config\" Not Found"))
	}
}

func TestLoad_UnmarshalError(t *testing.T) {
	resetConfig()
	if err := mockConfig(mockCfgFilePath, validCfg); err != nil {
		assert.NoError(t, err)
	}
	t.Logf(mockCfgCreateMsg, mockCfgFilePath)
	defer removeMockCfg(mockCfgFilePath)

	cfg := MockCfg{}
	err := Load(cfg)
	if assert.Error(t, err) {
		assert.EqualError(t, err, "result must be a pointer")
	}
}

func TestLoad_EmptyCgf(t *testing.T) {
	resetConfig()
	if err := mockConfig(mockCfgFilePath, emptyCfg); err != nil {
		assert.NoError(t, err)
	}
	t.Logf(mockCfgCreateMsg, mockCfgFilePath)
	defer removeMockCfg(mockCfgFilePath)

	cfg := &MockCfg{}
	err := Load(cfg)
	if assert.Error(t, err) {
		assert.EqualError(t, err, errors.MissingStartupConfigurationError([]string{"MOCK_URL", "MOCK_USERNAME", "MOCK_PASSWORD"}).Error())
	}
}

func TestLoad_IncompleteCgf(t *testing.T) {
	resetConfig()
	if err := mockConfig(mockCfgFilePath, incompleteCfg); err != nil {
		assert.NoError(t, err)
	}
	t.Logf(mockCfgCreateMsg, mockCfgFilePath)
	defer removeMockCfg(mockCfgFilePath)

	cfg := &MockCfg{}
	err := Load(cfg)
	if assert.Error(t, err) {
		assert.EqualError(t, err, errors.MissingStartupConfigurationError([]string{"MOCK_PASSWORD"}).Error())
	}
}

func resetConfig() {
	AddConfigPath(".")
	SetConfigName("config")
	SetConfigType("env")
}

func mockConfig(filePath string, config string) error {
	var (
		file *os.File
		err  error
	)

	// check if file exists
	_, err = os.Stat(filePath)
	if os.IsExist(err) {
		err := removeMockCfg(filePath)
		if err != nil {
			return err
		}
	}

	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// write default config to the file
	_, err = file.Write([]byte(config))
	if err != nil {
		return err
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func removeMockCfg(filePath string) error {
	var err = os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
