package logconfig

import
(
	"fmt"
	"time"
	"github.com/markoczy/goutil/log/handler"
)


const FATAL = 0
const ERROR = 1
const WARN = 2
const INFO = 3
const DEBUG = 4

var DefaultLogLevel int = DEBUG

// TODO how not make this public??
var LogHandlers []handler.LogHandler = []handler.LogHandler {
	DEFAULT_LOGHANDLER }

func Register(aHandler *handler.LogHandler) {
	LogHandlers = append(LogHandlers,*aHandler)
}

// func UnregisterAll() {
// 	for _, el range LogHandlers {
		
// 	}
// }

// TODO ptr to defaultLv
var DEFAULT_LOGHANDLER = handler.LogHandler{Level: DefaultLogLevel, 
	Format: DEFAULT_LOGFORMAT, Write: DEFAULT_LOGWRITE}

func DEFAULT_LOGWRITE(message string) {
	fmt.Println(message)
}

func DEFAULT_LOGFORMAT (message string, level string, file string, 
	method string, line int) string {

	t := time.Now()
	tStr := t.Format(time.RFC3339)
	return fmt.Sprintf(log_format, tStr, level, file, method, line, message)
}

const log_format = "[%s] %s %s::%s[%d]: %s"

