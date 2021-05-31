package dateutils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestGetDateTimeNow(t *testing.T) {
	assert.Equal(t, time.Now().UTC(), GetDateTimeNow())
}

func TestGetDateTimeNowFormat(t *testing.T) {
	assert.Equal(t, strings.Trim(time.Now().UTC().Format(dateTimeFormat), " ")[0], strings.Trim(GetDateTimeNowFormat(), " ")[0])
}
