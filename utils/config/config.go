package config

import (
	"fmt"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

const (
	missingConfigErrMsg  = "Missing mandatory configuration: %v"
	configLoadSuccessMsg = "Configuration '%s' loaded successfully"
	configLoadErrMsg     = "Error loading configuration '%s'"
)

var (
	configPathValue = "."
	configNameValue = "config"
	configTypeValue = "env"
)

// AddConfigPath adds a new path where configuration can be found. The path is set in viper's configuration along
// with the default current path ".".
func AddConfigPath(configPath string) {
	viper.AddConfigPath(configPathValue)
}

// SetConfigName sets the configNameValue global variable.
// Use this function to set a custom config file name.
func SetConfigName(configName string) {
	configNameValue = configName
}

// SetConfigType sets the configTypeValue global variable.
// Use this function to set a custom config file type.
func SetConfigType(configType string) {
	configTypeValue = configType
}

// Load reads the configuration from the config file and the environment. The function transforms the configuration
// into the input struct. The input should be an address to a valid config struct.
func Load(c interface{}) error {
	viper.AddConfigPath(configPathValue)
	viper.SetConfigName(configNameValue)
	viper.SetConfigType(configTypeValue)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error(fmt.Sprintf(configLoadErrMsg, reflect.TypeOf(c).Elem()), err)
		return err
	}

	err = viper.Unmarshal(c)
	if err != nil {
		logger.Error(fmt.Sprintf(configLoadErrMsg, reflect.TypeOf(c).Elem()), err)
		return err
	}

	if err := Validate(c); err != nil {
		return err
	}

	logger.Info(fmt.Sprintf(configLoadSuccessMsg, reflect.TypeOf(c).Elem()))
	return nil
}

// Validate check if all the required configuration is not empty.
// The struct filed that is required should have a tag "required: true"
// If a required struct field value is "" then the tag value of mapstructure will be returned.
func Validate(c interface{}) error {
	var missingParams []string
	sr := reflect.ValueOf(c).Elem()

	for i := 0; i < sr.NumField(); i++ {
		if strings.TrimSpace(sr.Field(i).String()) == "" && sr.Type().Field(i).Tag.Get("required") == "true" {
			missingParams = append(missingParams, sr.Type().Field(i).Tag.Get("mapstructure"))
		}
	}
	if len(missingParams) > 0 {
		return errors.NewError(fmt.Sprintf(missingConfigErrMsg, missingParams))
	}
	return nil
}
