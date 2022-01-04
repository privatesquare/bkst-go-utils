package structutils

import (
	"reflect"
	"strings"
)

const (
	JsonStructFieldKey         = "json"
	YamlStructFieldKey         = "yaml"
	MapstructureStructFieldKey = "mapstructure"
)

func GetJsonFieldTagValue(structPointer interface{}, fieldPointer interface{}) string {
	return GetFieldTagValue(structPointer, fieldPointer, JsonStructFieldKey)
}

func GetYamlFieldTagValue(structPointer interface{}, fieldPointer interface{}) string {
	return GetFieldTagValue(structPointer, fieldPointer, YamlStructFieldKey)
}

func GetMapstructureFieldTagValue(structPointer interface{}, fieldPointer interface{}) string {
	return GetFieldTagValue(structPointer, fieldPointer, MapstructureStructFieldKey)
}

func GetFieldTagValue(structPointer interface{}, fieldPointer interface{}, fieldKey string) string {
	var tagValue string
	structReflect := reflect.ValueOf(structPointer).Elem()
	fieldReflect := reflect.ValueOf(fieldPointer).Elem()
	for i := 0; i < structReflect.NumField(); i++ {
		fieldValue := structReflect.Field(i)
		if fieldValue.Addr().Interface() == fieldReflect.Addr().Interface() {
			tagValue = structReflect.Type().Field(i).Tag.Get(fieldKey)
		}
	}
	return strings.Split(tagValue, ",")[0]
}
