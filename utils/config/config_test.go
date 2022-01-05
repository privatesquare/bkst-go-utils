package config

import (
	"fmt"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/fileutils"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var (
	mockCnfCreateMsg = "Mock Configuration file %s created"

	url      = "https://test.com"
	username = "test"
	password = "test123"

	envCnfFileFmt = `MOCK_URL=%s
MOCK_USERNAME=%s
MOCK_PASSWORD=%s`

	jsonCnfFileFmt = `{
	"mock_url": "%s",
	"mock_username": "%s",
	"mock_password": "%s"
}`

	ymlCnfFileFmt = `---
mock_url: "%s"
mock_username: "%s"
mock_password: "%s"`

	validEnvCnf   = fmt.Sprintf(envCnfFileFmt, url, username, password)
	emptyEnvCnf   = fmt.Sprintf(envCnfFileFmt, "", "", "")
	invalidEnvCnf = fmt.Sprintf(envCnfFileFmt, url, username, "")

	validJsonCnf   = fmt.Sprintf(jsonCnfFileFmt, url, username, password)
	emptyJsonCnf   = fmt.Sprintf(jsonCnfFileFmt, "", "", "")
	invalidJsonCnf = fmt.Sprintf(jsonCnfFileFmt, url, username, "")

	validYmlCnf   = fmt.Sprintf(ymlCnfFileFmt, url, username, password)
	emptyYmlCnf   = fmt.Sprintf(ymlCnfFileFmt, "", "", "")
	invalidYmlCnf = fmt.Sprintf(ymlCnfFileFmt, url, username, "")
)

type MockEnvConfig struct {
	Url      string `mapstructure:"MOCK_URL" required:"true"`
	Username string `mapstructure:"MOCK_USERNAME" required:"true"`
	Password string `mapstructure:"MOCK_PASSWORD" required:"true"`
}

type MockJsonConfig struct {
	Url      string `mapstructure:"mock_url" required:"true"`
	Username string `mapstructure:"mock_username" required:"true"`
	Password string `mapstructure:"mock_password" required:"true"`
}

type MockYmlConfig struct {
	Url      string `mapstructure:"mock_url" required:"true"`
	Username string `mapstructure:"mock_username" required:"true"`
	Password string `mapstructure:"mock_password" required:"true"`
}

func TestAddConfigPath(t *testing.T) {
	AddConfigPath("./tst")
}

func TestSetConfigName(t *testing.T) {
	SetConfigName("test")
	assert.Equal(t, "test", configNameValue)
	resetConfig()
}

func TestSetConfigType(t *testing.T) {
	SetConfigType("yml")
	assert.Equal(t, "yml", configTypeValue)
	resetConfig()
}

func TestLoad(t *testing.T) {

	t.Run("env config", func(t *testing.T) {

		t.Run("valid config", func(t *testing.T) {
			err := writeMockConfig(validEnvCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockEnvConfig)
			err = Load(cnf)
			assert.NoError(t, err)
			assert.Equal(t, url, cnf.Url)
			assert.Equal(t, username, cnf.Username)
			assert.Equal(t, password, cnf.Password)
		})

		t.Run("config file does not exist", func(t *testing.T) {
			cnf := new(MockEnvConfig)
			err := Load(cnf)
			assert.Error(t, err)
			assert.True(t, strings.Contains(err.Error(), "Config File \"config\" Not Found"))
		})

		t.Run("unmarshal error", func(t *testing.T) {
			err := writeMockConfig(validEnvCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := MockEnvConfig{}
			assert.Panics(t, func() {
				err = Load(cnf)
			})
		})

		t.Run("empty config", func(t *testing.T) {
			err := writeMockConfig(emptyEnvCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockEnvConfig)
			err = Load(cnf)
			assert.Error(t, err)
			assert.EqualError(t, err, errors.New(fmt.Sprintf(missingConfigErrMsg, []string{"MOCK_URL", "MOCK_USERNAME", "MOCK_PASSWORD"})).Error())
		})

		t.Run("invalid config", func(t *testing.T) {
			err := writeMockConfig(invalidEnvCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockEnvConfig)
			err = Load(cnf)
			assert.Error(t, err)
			assert.EqualError(t, err, errors.New(fmt.Sprintf(missingConfigErrMsg, []string{"MOCK_PASSWORD"})).Error())
		})
	})

	t.Run("json config", func(t *testing.T) {

		t.Run("valid config", func(t *testing.T) {
			SetConfigType("json")
			err := writeMockConfig(validJsonCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockJsonConfig)
			err = Load(cnf)
			assert.NoError(t, err)
			assert.Equal(t, url, cnf.Url)
			assert.Equal(t, username, cnf.Username)
			assert.Equal(t, password, cnf.Password)
		})

		t.Run("empty config", func(t *testing.T) {
			SetConfigType("json")
			err := writeMockConfig(emptyJsonCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockJsonConfig)
			err = Load(cnf)
			assert.Error(t, err)
			assert.EqualError(t, err, errors.New(fmt.Sprintf(missingConfigErrMsg, []string{"mock_url", "mock_username", "mock_password"})).Error())
		})

		t.Run("invalid config", func(t *testing.T) {
			SetConfigType("json")
			err := writeMockConfig(invalidJsonCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockJsonConfig)
			err = Load(cnf)
			assert.Error(t, err)
			assert.EqualError(t, err, errors.New(fmt.Sprintf(missingConfigErrMsg, []string{"mock_password"})).Error())
		})
	})

	t.Run("yml config", func(t *testing.T) {

		t.Run("valid config", func(t *testing.T) {
			SetConfigType("yml")
			err := writeMockConfig(validYmlCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockYmlConfig)
			err = Load(cnf)
			assert.NoError(t, err)
			assert.Equal(t, url, cnf.Url)
			assert.Equal(t, username, cnf.Username)
			assert.Equal(t, password, cnf.Password)
		})

		t.Run("empty config", func(t *testing.T) {
			SetConfigType("yml")
			err := writeMockConfig(emptyYmlCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockYmlConfig)
			err = Load(cnf)
			assert.Error(t, err)
			assert.EqualError(t, err, errors.New(fmt.Sprintf(missingConfigErrMsg, []string{"mock_url", "mock_username", "mock_password"})).Error())
		})

		t.Run("invalid config", func(t *testing.T) {
			SetConfigType("yml")
			err := writeMockConfig(invalidYmlCnf)
			assert.NoError(t, err)
			logger.Infof(mockCnfCreateMsg, getMockConfigFilePath())
			defer resetConfig()

			cnf := new(MockYmlConfig)
			err = Load(cnf)
			assert.Error(t, err)
			assert.EqualError(t, err, errors.New(fmt.Sprintf(missingConfigErrMsg, []string{"mock_password"})).Error())
		})
	})
}

func getMockConfigFilePath() string {
	return configPathValue + "/" + configNameValue + "." + configTypeValue
}

func resetConfig() {
	if err := fileutils.RemoveFile(getMockConfigFilePath()); err != nil {
		logger.Panic(err.Error())
	}
	AddConfigPath(".")
	SetConfigName("config")
	SetConfigType("env")
}

func writeMockConfig(data string) error {
	var (
		filePath = getMockConfigFilePath()
	)
	if err := fileutils.RemoveFile(filePath); err != nil {
		return err
	}
	if err := fileutils.WriteFile(filePath, []byte(data)); err != nil {
		return err
	}
	return nil
}
