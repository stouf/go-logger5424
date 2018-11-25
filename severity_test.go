package logger5424

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringReturnsExpectedStringForEmergency(t *testing.T) {
	assert.Equal(t, Emergency.String(), "emergency")
}

func TestStringReturnsExpectedStringForAlert(t *testing.T) {
	assert.Equal(t, Alert.String(), "alert")
}

func TestStringReturnsExpectedStringForCritical(t *testing.T) {
	assert.Equal(t, Critical.String(), "critical")
}

func TestStringReturnsExpectedStringForError(t *testing.T) {
	assert.Equal(t, Error.String(), "error")
}

func TestStringReturnsExpectedStringForWarning(t *testing.T) {
	assert.Equal(t, Warning.String(), "warning")
}

func TestStringReturnsExpectedStringForNotice(t *testing.T) {
	assert.Equal(t, Notice.String(), "notice")
}

func TestStringReturnsExpectedStringForInfo(t *testing.T) {
	assert.Equal(t, Info.String(), "info")
}

func TestStringReturnsExpectedStringForDebug(t *testing.T) {
	assert.Equal(t, Debug.String(), "debug")
}

func TestStringPanicsForInvalidValues(t *testing.T) {
	var x Severity = -5
	f := func() {
		_ = x.String()
	}
	assert.Panics(t, f)
}
