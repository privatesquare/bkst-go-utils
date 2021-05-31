package structutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	FirstName string `yaml:"first_name_yaml" json:"first_name_json" mapstructure:"first_name_ms"`
	LastName string `yaml:"last_name_yaml" json:"last_name_json"`
	Age int64 `json:"age,omitempty"`
}

func TestGetFieldTagValue(t *testing.T) {
	user := new(User)

	// based on precedence test if mapstructure tag value is returned
	tag := GetFieldTagValue(user, &user.FirstName)
	assert.Equal(t, "first_name_ms", tag)

	// based on precedence test if yaml tag value is returned
	tag = GetFieldTagValue(user, &user.LastName)
	assert.Equal(t, "last_name_yaml", tag)

	// test if the first value is returned when a tag has multiple values
	tag = GetFieldTagValue(user, &user.Age)
	assert.Equal(t, "age", tag)
}