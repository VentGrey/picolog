package picolog

import (
	"fmt"
	"sync"
	"time"
)

// LogLevel is an enum that represents the different log levels.
type LogLevel int

// The different log levels.
const (
	Info LogLevel = iota
	Debug
	Warning
	Error
	Ok
)

var logLevelColours = map[LogLevel]string{
	Info:    "\033[34m",
	Debug:   "\033[35m",
	Warning: "\033[33m",
	Error:   "\033[31m",
	Ok:      "\033[32m",
}

var logLevelNames = map[LogLevel]string{
	Info:    "INFO",
	Debug:   "DEBUG",
	Warning: "WARNING",
	Error:   "ERROR",
	Ok:      "OK",
}

// Logger is the main struct for the logger.
type Logger struct {
	// PkgName defines the package name for the logger. This package name is
	// used in the postfix of the log message that reads "At package: <pkgName>"
	// This helps identifying where the log message comes from and makes things
	// easier to debug.
	PkgName string

	// MinLogLevel defines the minimum log level that will be printed. If the
	// log level is lower than this, it won't be printed.
	MinLogLevel LogLevel

	// EnableColours defines if the logger should print the log levels in
	// colours or not. This is useful for when the output is not a terminal or
	// if the terminal doesn't support ANSI escape codes.
	EnableColours bool

	// Mutex lock for thread safety.
	mu sync.Mutex
}

// NewLogger creates a new logger with the provided package name. This helps identifying where the log message comes from.
// It also makes it easier to identify id an error is ours or from the user.
// Optionally, you can provide a minimum log level and if the logger should print the log levels in colours or not.
func NewLogger(pkgName string, minLogLevel LogLevel, enableColours bool) *Logger {
	return &Logger{PkgName: pkgName, MinLogLevel: minLogLevel, EnableColours: enableColours}
}

// Log logs a message with the provided log level.
// Optionally, multiple errors can be provided to be logged as well.
func (l *Logger) Log(level LogLevel, message string, errs ...error) error {

	// Try making this thread safe. If used in a multithreaded environment,
	// this will prevent the logs from being mixed up or become corrupted.
	l.mu.Lock()
	defer l.mu.Unlock()

	if level < l.MinLogLevel {
		return nil
	}

	if level < Info || level > Ok {
		fmt.Println("Picolog Error: Invalid log level provided, defaulting to INFO")
		level = Info
	}

	levelName := logLevelNames[level]
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var logMessage string

	colorPrefix := ""
	colorSuffix := ""

	if l.EnableColours {
		colorPrefix = logLevelColours[level]
		colorSuffix = "\033[0m"
	}

	if len(errs) > 0 {
		for _, err := range errs {
			if err != nil {
				logMessage = fmt.Sprintf("%s[%s] - %s : %s %v - At package: %s%s", colorPrefix, levelName, currentTime, message, err, l.PkgName, colorSuffix)
			}
		}
	} else {
		logMessage = fmt.Sprintf("%s[%s] - %s : %s - At package: %s%s", colorPrefix, levelName, currentTime, message, l.PkgName, colorSuffix)
	}

	fmt.Println(logMessage)
	return nil
}
