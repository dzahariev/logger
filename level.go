package logger

// Level represents the logging level
type Level struct {
	ID   int
	Name string
}

// LogLevel Global level for the application
var LogLevel = Error

func newLevel(id int, name string) Level {
	return Level{
		ID:   id,
		Name: name,
	}
}

const (
	panicID = 0
	fatalID = 1
	errorID = 2
	warnID  = 3
	infoID  = 4
	debugID = 5
	traceID = 6
)

var (
	// Panic level
	Panic = newLevel(panicID, "panic")

	// Fatal level
	Fatal = newLevel(fatalID, "fatal")

	// Error level
	Error = newLevel(errorID, "error")

	// Warn level
	Warn = newLevel(warnID, "warn")

	// Info level
	Info = newLevel(infoID, "info")

	// Debug level
	Debug = newLevel(debugID, "debug")

	// Trace level
	Trace = newLevel(traceID, "trace")
)

// AllLevels logging levels
func AllLevels() map[string]Level {
	return map[string]Level{
		Panic.Name: Panic,
		Fatal.Name: Fatal,
		Error.Name: Error,
		Warn.Name:  Warn,
		Info.Name:  Info,
		Debug.Name: Debug,
		Trace.Name: Trace,
	}
}
