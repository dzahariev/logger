package logger

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Logger usage
//
// Global initialisation with configuration
//
//  logger.InitLogger(appConfig.LogConfig)
//
// Get logger from context - the context prepares the logger format, level and correlation ID:
//
//  Context ctx
//  ctx.getLogger("main logger")
//
// Create logger in place:
//
//  mainLogger := logger.NewNamedLogger("main logger")
//  mainLogger.GenTraceID()
//
// To log a message with Fatal severity:
//
//  mainLogger.Fatal("Fatal entry")
//
// Output depend on assigned formatter.
//
// For string formatter it look like this:
//
//  2020-04-04T17:22:17+03:00 [fatal] e30375a6-4dc4-40be-8c5f-8656996298ce [main logger]: [[[Fatal entry]]]
//
// For JSON formatter:
//
//  {"time":"2020-04-04T17:15:12+03:00","level":"fatal","corelation_id":"ace08634-1245-4557-9754-d015f8d0235a","logger":"main logger","msg":"[[[Fatal entry]]]"}
type Logger struct {
	// Name of the logger
	Name string

	// TraceID of the logger
	TraceID uuid.UUID
}

// InitLogger initialise the global logger from configuration
func InitLogger(config Config) error {
	if _, ok := AllFormatters()[config.Format]; ok {
		LogFormatter = AllFormatters()[config.Format]
	} else {
		return fmt.Errorf("format value %s, should be one of %v", config.Format, AllFormatters())
	}
	if _, ok := AllLevels()[config.Level]; ok {
		LogLevel = AllLevels()[config.Level]
	} else {
		return fmt.Errorf("level value %s, should be one of %v", config.Level, AllLevels())
	}
	return nil
}

// NewNamedLogger creates new named logger
func NewNamedLogger(name string) Logger {
	result := Logger{
		Name: name,
	}
	result.GenTraceID()
	return result
}

// NewAnnonymousLogger creates new logger without name
func NewAnnonymousLogger() Logger {
	return NewNamedLogger("-")
}

// GenTraceID is used to generate and set new Correlation ID
func (l *Logger) GenTraceID() {
	UUID, err := uuid.NewV4()
	if err != nil {
		UUID = uuid.Nil
	}
	l.TraceID = UUID
}

// SetTraceID is used to set Correlation ID
func (l *Logger) SetTraceID(traceID uuid.UUID) {
	l.TraceID = traceID
}

// SetName is used to set Name
func (l *Logger) SetName(newName string) {
	l.Name = newName
}

// SetName is used to set Name
func (l *Logger) isEnabled(level Level) bool {
	return level.ID <= LogLevel.ID
}

func (l *Logger) prepareMsg(level Level, v ...interface{}) string {
	return LogFormatter.Parse(time.Now(), level, l.TraceID.String(), l.Name, v)
}

// Panicf ently logged
func (l *Logger) Panicf(format string, v ...interface{}) {
	if l.isEnabled(Panic) {
		fmt.Println(l.prepareMsg(Panic, fmt.Sprintf(format, v...)))
	}
}

// Panic ently logged
func (l *Logger) Panic(v ...interface{}) {
	if l.isEnabled(Panic) {
		fmt.Println(l.prepareMsg(Panic, v...))
	}
}

// Fatalf ently logged
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.isEnabled(Fatal) {
		fmt.Println(l.prepareMsg(Fatal, fmt.Sprintf(format, v...)))
	}
}

// Fatal ently logged
func (l *Logger) Fatal(v ...interface{}) {
	if l.isEnabled(Fatal) {
		fmt.Println(l.prepareMsg(Fatal, v...))
	}
}

// Errorf ently logged
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.isEnabled(Error) {
		fmt.Println(l.prepareMsg(Error, fmt.Sprintf(format, v...)))
	}
}

// Error ently logged
func (l *Logger) Error(v ...interface{}) {
	if l.isEnabled(Error) {
		fmt.Println(l.prepareMsg(Error, v...))
	}
}

// Warnf ently logged
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.isEnabled(Warn) {
		fmt.Println(l.prepareMsg(Warn, fmt.Sprintf(format, v...)))
	}
}

// Warn ently logged
func (l *Logger) Warn(v ...interface{}) {
	if l.isEnabled(Warn) {
		fmt.Println(l.prepareMsg(Warn, v...))
	}
}

// Infof ently logged
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.isEnabled(Info) {
		fmt.Println(l.prepareMsg(Info, fmt.Sprintf(format, v...)))
	}
}

// Info ently logged
func (l *Logger) Info(v ...interface{}) {
	if l.isEnabled(Info) {
		fmt.Println(l.prepareMsg(Info, v...))
	}
}

// Debugf ently logged
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.isEnabled(Debug) {
		fmt.Println(l.prepareMsg(Debug, fmt.Sprintf(format, v...)))
	}
}

// Debug ently logged
func (l *Logger) Debug(v ...interface{}) {
	if l.isEnabled(Debug) {
		fmt.Println(l.prepareMsg(Debug, v...))
	}
}

// Tracef ently logged
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.isEnabled(Trace) {
		fmt.Println(l.prepareMsg(Trace, fmt.Sprintf(format, v...)))
	}
}

// Trace ently logged
func (l *Logger) Trace(v ...interface{}) {
	if l.isEnabled(Trace) {
		fmt.Println(l.prepareMsg(Trace, v...))
	}
}
