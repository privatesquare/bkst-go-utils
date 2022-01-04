package fileutils

import (
	"encoding/json"
	"fmt"
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

// FileNotFoundError represents an error when the file is not found
type FileNotFoundError string

// Error returns the formatted FileNotFoundError
func (fnf FileNotFoundError) Error() string {
	return fmt.Sprintf(fileNotFoundErrMsg, string(fnf))
}

// FileCreateError represents an error when the code is not able to create a file
type FileCreateError struct {
	File string
	Err  error
}

// Error returns the formatted FileCreateError
func (fc FileCreateError) Error() string {
	return fmt.Sprintf(fileCreateErrMsg, fc.File, fc.Err)
}

// FileOpenError represents an error when the code is not able to open the file
type FileOpenError struct {
	File string
	Err  error
}

// Error returns the formatted FileOpenError
func (fo FileOpenError) Error() string {
	return fmt.Sprintf(fileOpenErrMsg, fo.File, fo.Err)
}

// FileReadError represents an error when the code is not able to read the file
type FileReadError struct {
	File string
	Err  error
}

// Error returns the formatted FileReadError
func (fr FileReadError) Error() string {
	return fmt.Sprintf(fileReadErrMsg, fr.File, fr.Err)
}

// FileWriteError represents an error when the code is not able to write to the file
type FileWriteError struct {
	File string
	Err  error
}

// Error returns the formatted FileWriteError
func (fw FileWriteError) Error() string {
	return fmt.Sprintf(fileWriteErrMsg, fw.File, fw.Err)
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
		return errors.JSONUnMarshalError{Err: err}
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
		return errors.JSONMarshalError{Err: err}
	}
	err = WriteFile(filePath, data)
	if err != nil {
		return err
	}
	return nil
}

// CreateFile creates a new file
// The method returns an error if there was an issue with creating an new file
func CreateFile(file string) (*os.File, error) {
	f, err := os.Create(file)
	if err != nil {
		return f, FileCreateError{File: file, Err: err}
	}
	return f, nil
}

// OpenFile opens a file
// The method returns an error if there is an issue with opening the file
func OpenFile(file string) (*os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return f, FileOpenError{File: file, Err: err}
	}
	return f, nil
}

// ReadFile checks if a file exists and if it does tries to reads the contents of the
// file and returns the data back
// The method returns an error the file does not exist or if there was an error in reading the contents of the file
func ReadFile(file string) ([]byte, error) {
	if !FileExists(file) {
		return nil, FileNotFoundError(file)
	} else if data, err := ioutil.ReadFile(file); err != nil {
		return nil, FileReadError{File: file, Err: err}
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
		return FileWriteError{File: file, Err: err}
	}
	return nil
}

// FileExists checks if a file exists and returns an error if the file was not found
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

func RemoveFile(filePath string) error {
	if FileExists(filePath) {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}
	return nil
}
