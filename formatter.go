package logger

import (
	"fmt"
	"time"
)

// Formatter represents the logging Formatter
type Formatter struct {
	Name    string
	Pattern string
}

// LogFormatter Global formatter for the application
var LogFormatter = StringFormatter

func newFormatter(name string, pattern string) Formatter {
	return Formatter{
		Name:    name,
		Pattern: pattern,
	}
}

// Parse string
func (f *Formatter) Parse(timestamp time.Time, level Level, correlationID string, logerName string, v ...interface{}) string {
	return fmt.Sprintf(f.Pattern, timestamp.Format(time.RFC3339), level.Name, correlationID, logerName, v)
}

var (
	// StringFormatter instance
	StringFormatter = newFormatter("string", "%s [%s] %s [%s]: %v")

	// JSONFormatter instance
	JSONFormatter = newFormatter("json", "{\"time\":\"%s\",\"level\":\"%s\",\"trace_id\":\"%s\",\"logger\":\"%s\",\"msg\":\"%v\"}")
)

// AllFormatters available formatters
func AllFormatters() map[string]Formatter {
	return map[string]Formatter{
		StringFormatter.Name: StringFormatter,
		JSONFormatter.Name:   JSONFormatter,
	}
}
