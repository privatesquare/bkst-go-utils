package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEntryExists(t *testing.T) {

	slice := []string{"one", "two", "three"}

	result := EntryExists(slice, "one")

	if result == false {
		t.Errorf("Test Failed!, expected: %v, got: %v", true, result)
	}

	result = EntryExists(slice, "four")

	if result == true {
		t.Errorf("Test Failed!, expected: %v, got: %v", false, result)
	}
}

func ExampleEntryExists() {
	slice := []string{"one", "two", "three"}
	fmt.Println(EntryExists(slice, "two"))
	// Output: true
}

func TestGetSliceEntryIndex(t *testing.T) {

	slice := []string{"one", "two", "three"}

	result := GetSliceEntryIndex(slice, "one")

	if result != 0 {
		t.Errorf("Test Failed!, expected: %v, got: %v", 0, result)
	}

	result = GetSliceEntryIndex(slice, "doesNotExist")

	if result != -1 {
		t.Errorf("Test Failed!, expected: %v, got: %v", -1, result)
	}
}

func ExampleGetSliceEntryIndex() {
	slice := []string{"one", "two", "three"}
	fmt.Println(GetSliceEntryIndex(slice, "one"))
	// Output: 0
}

func TestRemoveEntryFromSlice(t *testing.T) {
	slice := []string{"one", "two", "three"}
	expectedResult := []string{"one", "three"}

	result := RemoveEntryFromSlice(slice, "two")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}

	slice = []string{"one", "two", "three"}
	result = RemoveEntryFromSlice(slice, "four")

	if reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test Failed!, expected: %v, got: %v", slice, result)
	}

	slice = []string{"one", "two", "three", "two"}
	expectedResult = []string{"one", "three", "two"}

	result = RemoveEntryFromSlice(slice, "two")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}

	slice = []string{"one", "two", "two", "three"}
	expectedResult = []string{"one", "two", "three"}

	result = RemoveEntryFromSlice(slice, "two")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}
}

func ExampleRemoveEntryFromSlice() {
	slice := []string{"one", "two", "three"}
	fmt.Println(RemoveEntryFromSlice(slice, "two"))
	// Output: [one three]
}

func TestRemoveDuplicateEntries(t *testing.T) {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}
	expectedResult := []string{"one", "two", "three", "four"}

	result := RemoveDuplicateEntries(slice)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}
}

func ExampleRemoveDuplicateEntries() {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}
	fmt.Println(RemoveDuplicateEntries(slice))
	// output: [one two three four]
}

func TestCountDuplicateEntries(t *testing.T) {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}

	expectedResult := make(map[string]int)
	expectedResult["one"] = 1
	expectedResult["two"] = 2
	expectedResult["three"] = 3
	expectedResult["four"] = 4

	result := CountDuplicateEntries(slice)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}
}

func ExampleCountDuplicateEntries() {
	slice := []string{"one", "two", "two", "three", "three", "three"}
	fmt.Println(CountDuplicateEntries(slice))
	// Output: map[one:1 three:3 two:2]
}

func TestDuplicateEntryExists(t *testing.T) {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}

	result := DuplicateEntryExists(slice)

	if result == false {
		t.Errorf("Test Failed!, expected: %v, got: %v", true, result)
	}

	slice = []string{"one", "two", "three", "four"}

	result = DuplicateEntryExists(slice)

	if result == true {
		t.Errorf("Test Failed!, expected: %v, got: %v", false, result)
	}
}

func ExampleDuplicateEntryExists() {
	slice := []string{"one", "two", "two"}
	fmt.Println(DuplicateEntryExists(slice))
	// Output: true
}
