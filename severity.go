package logger5424

// Severity is a severity level such as defined by the RFC 5424.
// https://tools.ietf.org/html/rfc5424
type Severity int

const (
	Emergency Severity = iota
	Alert
	Critical
	Error
	Warning
	Notice
	Info
	Debug
)

// String returns the string version of a Severity.
func (s Severity) String() string {
	switch s {
	case Emergency:
		return "emergency"
	case Alert:
		return "alert"
	case Critical:
		return "critical"
	case Error:
		return "error"
	case Warning:
		return "warning"
	case Notice:
		return "notice"
	case Info:
		return "info"
	case Debug:
		return "debug"
	}
	panic("Unexpected value for type Severity")
}
