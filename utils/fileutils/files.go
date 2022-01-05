package fileutils

import (
	"encoding/json"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"io/ioutil"
	"os"
)

// File util constants
const (
	fileNotFoundErrMsg = "File '%s' was not found"
	fileCreateErrMsg   = "Unable to create file '%s' : %v"
	fileOpenErrMsg     = "Unable to open file '%s' : %v"
	fileReadErrMsg     = "Unable to read file '%s' : %v"
	fileWriteErrMsg    = "Unable to write to the file '%s' : %v"
)

// FileExists checks if a file exists and returns an error if the file was not found
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

// OpenFile opens a file
// The method returns an error if there is an issue with opening the file
func OpenFile(file string) (*os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return f, errors.Newf(fileOpenErrMsg, file, err)
	}
	return f, nil
}

// CreateFile creates a new file
// The method returns an error if there was an issue with creating an new file
func CreateFile(file string) (*os.File, error) {
	f, err := os.Create(file)
	if err != nil {
		return f, errors.Newf(fileCreateErrMsg, file, err)
	}
	return f, nil
}

// ReadFile checks if a file exists and if it does try to read the contents of the
// file and returns the data back
// The method returns an error the file does not exist or if there was an error in reading the contents of the file
func ReadFile(file string) ([]byte, error) {
	if !FileExists(file) {
		return nil, errors.Newf(fileNotFoundErrMsg, file)
	} else if data, err := ioutil.ReadFile(file); err != nil {
		return nil, errors.Newf(fileReadErrMsg, file, err)
	} else {
		return data, err
	}
}

// WriteFile creates a new file if the file does not exists and writes data into the file
// The method returns an error if there was an issue creating a new file
// or while writing data into the file
func WriteFile(file string, data []byte) error {
	var (
		err error
	)
	if !FileExists(file) {
		_, err = CreateFile(file)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return errors.Newf(fileWriteErrMsg, file, err)
	}
	return nil
}

// RemoveFile removes files from the provided valid filePath.
func RemoveFile(filePath string) error {
	if FileExists(filePath) {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}
	return nil
}

// ReadJsonFile reads a yaml file and puts the contents into the out variables
// out variable should be a pointer to a valid struct
// The method returns and error if reading a file or the unmarshal process fails
func ReadJsonFile(filePath string, out interface{}) error {
	data, err := ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, out)
	if err != nil {
		return errors.New(err.Error())
	}
	return err
}

// WriteJSONFile encodes the data from an input interface into json format
// and writes the data into a file
// The in interface should be an address to a valid struct
// The method returns an error if there is an error with the json encode
// or with writing to the file
func WriteJSONFile(filePath string, in interface{}) error {
	data, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return errors.New(err.Error())
	}
	err = WriteFile(filePath, data)
	if err != nil {
		return err
	}
	return nil
}
