package fileutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	validFilePath       = "test/file.txt"
	invalidFilePath     = "test/invalid.txt"
	validJsonFilePath   = "test/valid.json"
	invalidJsonFilePath = "test/invalid.json"
)

type Out struct {
	Foo string `json:"foo"`
}

func TestFileExists(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		assert.True(t, FileExists(validFilePath))
	})

	t.Run("error", func(t *testing.T) {
		assert.False(t, FileExists(invalidFilePath))
	})
}

func TestOpenFile(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		_, err := OpenFile(validFilePath)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		_, err := OpenFile(invalidFilePath)
		assert.Error(t, err)
	})
}

func TestCreateAndRemoveFile(t *testing.T) {

	filePath := "test/test.txt"

	t.Run("create", func(t *testing.T) {
		_, err := CreateFile(filePath)
		assert.NoError(t, err)
		assert.True(t, FileExists(filePath))
	})

	t.Run("remove", func(t *testing.T) {
		err := RemoveFile(filePath)
		assert.NoError(t, err)
		assert.False(t, FileExists(filePath))
	})
}

func TestReadFile(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		data, err := ReadFile(validFilePath)
		assert.NoError(t, err)
		assert.Equal(t, "something", string(data))
	})

	t.Run("invalid", func(t *testing.T) {
		_, err := ReadFile(invalidFilePath)
		assert.Error(t, err)
	})
}

func TestReadJsonFile(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		out := new(Out)
		err := ReadJsonFile(validJsonFilePath, out)
		assert.NoError(t, err)
		assert.Equal(t, "bar", out.Foo)
	})

	t.Run("invalid", func(t *testing.T) {
		out := new(Out)
		err := ReadJsonFile(invalidJsonFilePath, out)
		assert.Error(t, err)
	})
}

func TestWriteFile(t *testing.T) {
	filePath := "test/test.txt"
	err := WriteFile(filePath, []byte("something"))
	assert.NoError(t, err)

	err = RemoveFile(filePath)
	assert.NoError(t, err)
}

func TestWriteJSONFile(t *testing.T) {
	filePath := "test/test.txt"
	in := &Out{Foo: "bar"}
	err := WriteJSONFile(filePath, in)
	assert.NoError(t, err)

	err = RemoveFile(filePath)
	assert.NoError(t, err)
}
