// Package logger5424 lets you generate messages with a severity, a message and
// optional contextual information. The current time will automatically be
// attached to the generated log for you.
package logger5424

import (
	"encoding/json"
	"io"
	"time"
)

// Logger is the structure that exposes the main API.
type Logger struct {
	writer io.Writer
}

type log struct {
	Time     string                  `json:"time"`
	Severity string                  `json:"severity"`
	Message  string                  `json:"message"`
	Info     *map[string]interface{} `json:"info"`
}

// New creates a new instance of Logger. The resulting Logger instance will
// write logs to the writer you provide.
func New(writer io.Writer) Logger {
	return Logger{writer}
}

// Json writes logs using the JSON format. Here is an example of the expected
// format:
// {
//   "time": "2006-01-02 15:04:05.999999999 -0700 MST",
//   "severity": "info",
//   "message": "my message",
//   "info": {
//     "foo": "bar"
//   }
// }
func (l Logger) Json(
	severity Severity,
	message string,
	info *map[string]interface{},
) error {
	now := time.Now().UTC().String()
	logAsJson, err := json.Marshal(log{
		now,
		severity.String(),
		message,
		info,
	})
	if err != nil {
		return err
	}
	_, err = l.writer.Write(logAsJson)
	if err != nil {
		return err
	}
	return nil
}
