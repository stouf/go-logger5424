package logger5424

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testWriter struct {
	logs [][]byte
}

func (w *testWriter) Write(p []byte) (int, error) {
	w.logs = append(w.logs, p)
	return len(p), nil
}

func (w testWriter) getFirstLog() (log, error) {
	if len(w.logs) < 1 {
		return log{}, errors.New("No logs registered")
	}
	var x log
	err := json.Unmarshal(w.logs[0], &x)
	if err != nil {
		return log{}, err
	}
	return x, nil
}

func TestJsonWritesLogWithCurrentUtcTime(t *testing.T) {
	w := testWriter{}
	logger := New(&w)
	err := logger.Json(Info, "Hey!", nil)
	assert.NoError(t, err)
	log, err := w.getFirstLog()
	assert.NoError(t, err)
	date, err := time.Parse(
		"2006-01-02 15:04:05.999999999 -0700 MST",
		log.Time,
	)
	assert.NoError(t, err)
	assert.Equal(t, time.UTC, date.Location())
	now := time.Now().UTC()
	assert.True(t, now.Unix()-date.Unix() < 3)
}

func TestJsonWritesExpectedSeverity(t *testing.T) {
	w := testWriter{}
	logger := New(&w)
	severity := Info
	err := logger.Json(severity, "foo", nil)
	assert.NoError(t, err)
	log, err := w.getFirstLog()
	assert.NoError(t, err)
	assert.Equal(t, severity.String(), log.Severity)
}

func TestJsonWritesExpectedMessage(t *testing.T) {
	w := testWriter{}
	logger := New(&w)
	msg := "test message"
	err := logger.Json(Warning, msg, nil)
	assert.NoError(t, err)
	log, err := w.getFirstLog()
	assert.NoError(t, err)
	assert.Equal(t, msg, log.Message)
}

func TestJsonWritesExpectedInfo(t *testing.T) {
	w := testWriter{}
	logger := New(&w)
	info := map[string]interface{}{
		"foo": "bar",
	}
	err := logger.Json(Warning, "foo", &info)
	assert.NoError(t, err)
	log, err := w.getFirstLog()
	assert.NoError(t, err)
	assert.NotNil(t, log.Info)
	assert.Equal(t, "bar", (*log.Info)["foo"])
}

type brokenWriterError struct{}

func (e brokenWriterError) Error() string {
	return "Broken writer"
}

type brokenWriter struct {
}

func (w brokenWriter) Write(p []byte) (int, error) {
	return 0, brokenWriterError{}
}

func TestJsonReturnsOnErrorFromWriter(t *testing.T) {
	w := brokenWriter{}
	logger := New(w)
	err := logger.Json(Info, "foo", nil)
	assert.Error(t, err)
	_, ok := err.(brokenWriterError)
	assert.True(t, ok)
}
