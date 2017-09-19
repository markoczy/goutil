package log

import
(
	"fmt"
	"github.com/markoczy/goutil/log/trace"
	"github.com/markoczy/goutil/log/handler"
	"github.com/markoczy/goutil/log/logconfig"
)

var levels = []string{"FATAL","ERROR","WARN","INFO","DEBUG"}

func Debug(aMessage string) {
	logWrite(logconfig.DEBUG, false, aMessage)
}

func Info(aMessage string) {
	logWrite(logconfig.INFO, false, aMessage)
}

func Warn(aMessage string) {
	logWrite(logconfig.INFO, false, aMessage)
}

func Error(aMessage string) {
	logWrite(logconfig.INFO, false, aMessage)
}

func Fatal(aMessage string) {
	logWrite(logconfig.INFO, false, aMessage)
}

// Format

func Debugf(aMessage string, a ...interface{}) {
	logWrite(logconfig.DEBUG, false, fmt.Sprintf(aMessage, a...))
}

func Infof(aMessage string, a ...interface{}) {
	logWrite(logconfig.INFO, false, fmt.Sprintf(aMessage, a...))
}

func Warnf(aMessage string, a ...interface{}) {
	logWrite(logconfig.INFO, false, fmt.Sprintf(aMessage, a...))
}

func Errorf(aMessage string, a ...interface{}) {
	logWrite(logconfig.INFO, false, fmt.Sprintf(aMessage, a...))
}

func Fatalf(aMessage string, a ...interface{}) {
	logWrite(logconfig.INFO, false, fmt.Sprintf(aMessage, a...))
}

func logWrite(level int, format bool, aMessage string) {

	// Check if any has log lv high enough
	doWrite := false
	for _, hndl := range logconfig.LogHandlers() {
		if levelOrDefault(hndl) >= level { doWrite = true }
	}
	if !doWrite { return }

	// Get Stack trace
	t, err := trace.Trace(3)
	if err != nil {
		fmt.Println("Logging failed, stack trace not retreived:", err)
	}

	for _, hndl := range logconfig.LogHandlers() {
		txt, _ := hndl.Format(aMessage, levels[level],t.File,t.Method,t.Line)
		if levelOrDefault(hndl)>=level { hndl.Write(txt) }
	}
}

func levelOrDefault(hndl handler.LogHandler) int {
	lv, _ := hndl.Level()
	if lv == logconfig.DEFAULT { return logconfig.DefaultLogLevel() }
	return lv
}