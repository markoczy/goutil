package logconfig

import
(
	"fmt"
	"time"
	"github.com/markoczy/goutil/log/handler"
)

// FATAL use log level fatal (0)
const FATAL = 0
// ERROR use log level error (1)
const ERROR = 1
// WARN use log level warn (2)
const WARN = 2
// INFO use log level info (3)
const INFO = 3
// DEBUG use log level debug (4)
const DEBUG = 4
// DEFAULT use default log level (-100)
const DEFAULT = -100

// DefaultLogLevel The default level of logging
var defaultLogLevel = DEBUG

// SetDefaultLogLevel ...
func SetDefaultLogLevel(level int) {
	defaultLogLevel = level
}

// DefaultLogLevel ...
func DefaultLogLevel() int {
	return defaultLogLevel
}

// DefaultLogHandler ...
var DefaultLogHandler = handler.NewConsoleLogger(DEFAULT, DefaultLogFormat)

// DefaultLogFormat ...
func DefaultLogFormat (message string, level string, file string, 
	method string, line int) (string, error) {

	t := time.Now()
	tStr := t.Format(time.RFC3339)
	return fmt.Sprintf(logFormat, tStr, level, file, method, line,
		message), nil
}

var logHandlers = []handler.LogHandler { DefaultLogHandler }

func LogHandlers() []handler.LogHandler {
	return logHandlers
}

func Register(aHandler handler.LogHandler) {
	logHandlers = append(logHandlers,aHandler)
}

func UnregisterAll() {
	for _, el := range logHandlers {
		el.Unlink()
	}
	logHandlers = []handler.LogHandler{}
}


const logFormat = "[%s] %s %s::%s[%d]: %s"

