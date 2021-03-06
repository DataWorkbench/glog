package glog

// Level defines log levels.
type Level int8

const (
	// NoLevel is the lowest level.
	NoLevel Level = iota
	// DebugLevel defines debug log level.
	DebugLevel
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	//// PanicLevel defines panic log level.
	//PanicLevel
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	//case PanicLevel:
	//	return "panic"
	default:
		return ""
	}
}
