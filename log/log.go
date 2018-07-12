package log

import (
	"fmt"
	"github.com/markoczy/goutil/log/handler"
	"github.com/markoczy/goutil/log/logconfig"
	"github.com/markoczy/goutil/log/trace"
)

var levels = []string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG"}

// Debug Log message with loglevel debug
func Debug(aMessage string) {
	logWrite(logconfig.DEBUG, false, aMessage)
}

// Info Log message with loglevel info
func Info(aMessage string) {
	logWrite(logconfig.INFO, false, aMessage)
}

// Warn Log message with loglevel warning
func Warn(aMessage string) {
	logWrite(logconfig.WARN, false, aMessage)
}

// Error Log message with loglevel error
func Error(aMessage string) {
	logWrite(logconfig.ERROR, false, aMessage)
}

// Fatal Log message with loglevel fatal
func Fatal(aMessage string) {
	logWrite(logconfig.FATAL, false, aMessage)
}

// Format

// Debugf Log formatted message with loglevel debug
func Debugf(aMessage string, a ...interface{}) {
	logWrite(logconfig.DEBUG, false, fmt.Sprintf(aMessage, a...))
}

// Infof Log formatted message with loglevel info
func Infof(aMessage string, a ...interface{}) {
	logWrite(logconfig.INFO, false, fmt.Sprintf(aMessage, a...))
}

// Warnf Log formatted message with loglevel warning
func Warnf(aMessage string, a ...interface{}) {
	logWrite(logconfig.WARN, false, fmt.Sprintf(aMessage, a...))
}

// Errorf Log formatted message with loglevel error
func Errorf(aMessage string, a ...interface{}) {
	logWrite(logconfig.ERROR, false, fmt.Sprintf(aMessage, a...))
}

// Fatalf Log formatted message with loglevel fatal
func Fatalf(aMessage string, a ...interface{}) {
	logWrite(logconfig.FATAL, false, fmt.Sprintf(aMessage, a...))
}

func logWrite(level int, format bool, aMessage string) {

	// Check if any has log lv high enough
	doWrite := false
	for _, hndl := range logconfig.LogHandlers() {
		if levelOrDefault(hndl) >= level {
			doWrite = true
		}
	}
	if !doWrite {
		return
	}

	// Get Stack trace
	t, err := trace.Trace(3)
	if err != nil {
		fmt.Println("Logging failed, stack trace not retreived:", err)
	}

	for _, hndl := range logconfig.LogHandlers() {
		txt, _ := hndl.Format(aMessage, levels[level], t.File, t.Method, t.Line)
		if levelOrDefault(hndl) >= level {
			hndl.Write(txt)
		}
	}
}

func levelOrDefault(hndl handler.LogHandler) int {
	lv, _ := hndl.Level()
	if lv == logconfig.DEFAULT {
		return logconfig.DefaultLogLevel()
	}
	return lv
}
