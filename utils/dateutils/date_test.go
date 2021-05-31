package dateutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetDateTimeNow(t *testing.T) {
	assert.Equal(t, time.Now().UTC(), GetDateTimeNow())
}

func TestGetDateTimeNowFormat(t *testing.T) {
	assert.Equal(t, time.Now().UTC().Format(dateTimeFormat), GetDateTimeNowFormat())
}
