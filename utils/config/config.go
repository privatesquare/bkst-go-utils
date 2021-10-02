package config

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

const (
	configLoadSuccessMsg = "Configuration loaded successfully"
	configLoadErrMsg     = "Unable to load the required configuration"
)

var (
	configPathValue = "."
	configNameValue = "config"
	configTypeValue = "env"
)

func AddConfigPath(configPath string) {
	configPathValue = configPath
}

func SetConfigName(configName string) {
	configNameValue = configName
}

func SetConfigType(configType string) {
	configTypeValue = configType
}

func Load(c interface{}) error {
	viper.AddConfigPath(configPathValue)
	viper.SetConfigName(configNameValue)
	viper.SetConfigType(configTypeValue)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error(configLoadErrMsg, err)
		return err
	}

	err = viper.Unmarshal(c)
	if err != nil {
		logger.Error(configLoadErrMsg, err)
		return err
	}

	if err := validate(c); err != nil {
		return err
	}

	logger.Info(configLoadSuccessMsg)
	return nil
}

func validate(c interface{}) error {
	var missingParams []string
	sr := reflect.ValueOf(c).Elem()

	for i := 0; i < sr.NumField(); i++ {
		if strings.TrimSpace(sr.Field(i).String()) == "" && sr.Type().Field(i).Tag.Get("required") == "true" {
			missingParams = append(missingParams, sr.Type().Field(i).Tag.Get("mapstructure"))
		}
	}
	if len(missingParams) > 0 {
		err := errors.MissingStartupConfigurationError(missingParams)
		logger.Error(err.Error(), nil)
		return err
	}
	return nil
}
