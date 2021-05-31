package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryExists(t *testing.T) {
	slice := []string{"one", "two", "three"}
	assert.True(t, EntryExists(slice, "one"))
	assert.False(t, EntryExists(slice, "four"))
}