package structutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	FirstName string `yml:"first_name_yml" json:"first_name_json" mapstructure:"first_name_ms"`
	LastName  string `yml:"last_name_yml" json:"last_name_json"`
	Age       int64  `json:"age_json,omitempty"`
}

func TestGetMapstructureFieldTagValue(t *testing.T) {
	user := new(User)
	tag := GetMapstructureFieldTagValue(user, &user.FirstName)
	assert.Equal(t, "first_name_ms", tag)
}

func TestGetJsonFieldTagValue(t *testing.T) {
	user := new(User)
	tag := GetJsonFieldTagValue(user, &user.FirstName)
	assert.Equal(t, "first_name_json", tag)

	tag = GetJsonFieldTagValue(user, &user.Age)
	assert.Equal(t, "age_json", tag)
}

func TestGetYamlFieldTagValue(t *testing.T) {
	user := new(User)
	tag := GetYmlFieldTagValue(user, &user.FirstName)
	assert.Equal(t, "first_name_yml", tag)
}

func TestGetFieldTagValue(t *testing.T) {
	user := new(User)

	tag := GetFieldTagValue(user, &user.FirstName, "mapstructure")
	assert.Equal(t, "first_name_ms", tag)

	tag = GetFieldTagValue(user, &user.LastName, "yml")
	assert.Equal(t, "last_name_yml", tag)

	// test if the first value is returned when a tag has multiple values
	tag = GetFieldTagValue(user, &user.Age, "json")
	assert.Equal(t, "age_json", tag)
}
