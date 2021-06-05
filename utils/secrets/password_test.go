package secrets

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRandomPassword(t *testing.T) {
	p0 := GetRandomPassword()
	assert.NotEmpty(t, p0)

	p1 := GetRandomPassword()
	p2 := GetRandomPassword()
	assert.NotEqual(t, p1, p2)
}

func TestVerifyPassword(t *testing.T) {
	// Test if password has a number in it
	password := "@Password"
	err := VerifyPassword(password)
	assert.EqualError(t, err, errors.InvalidPasswordError.Error())

	// Test if password has a upper case letter
	password = "@password123"
	err = VerifyPassword(password)
	assert.EqualError(t, err, errors.InvalidPasswordError.Error())

	// Test if password has a lower case letter
	password = "@PASSWORD123"
	err = VerifyPassword(password)
	assert.EqualError(t, err, errors.InvalidPasswordError.Error())

	// Test if password has a special letter
	password = "PASSWORD123"
	err = VerifyPassword(password)
	assert.EqualError(t, err, errors.InvalidPasswordError.Error())

	// Test password length
	password = "@123"
	err = VerifyPassword(password)
	assert.EqualError(t, err, errors.InvalidPasswordError.Error())

	// Test valid password
	password = "somePass@123"
	err = VerifyPassword(password)
	assert.NoError(t, err)
}

func TestEncryptPassword(t *testing.T) {
	password := "somepassword@123"
	ePass, err := EncryptPassword(password, "")
	if assert.NoError(t, err) {
		assert.NotEqual(t, password, ePass)
	}
}

func TestDecryptPassword(t *testing.T) {
	password := "somepassword@123"
	passphrase := "something"

	ePass, err := EncryptPassword(password, "")
	assert.NoError(t, err)
	dPass, err := DecryptPassword(ePass, "")
	assert.NoError(t, err)
	assert.Equal(t, password, dPass)

	ePass, err = EncryptPassword(password, passphrase)
	assert.NoError(t, err)
	dPass, err = DecryptPassword(ePass, passphrase)
	assert.NoError(t, err)
	assert.Equal(t, password, dPass)

	dPass, err = DecryptPassword("====ddwf=", "notValid")
	assert.Error(t, err)

	dPass, err = DecryptPassword(ePass, "notValid")
	assert.Error(t, err)
}
