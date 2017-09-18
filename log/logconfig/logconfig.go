package logconfig

import
(
	"github.com/markoczy/goutil/log/handler"
)


const FATAL = 0
const ERROR = 1
const WARN = 2
const INFO = 3
const DEBUG = 4

var STRLEVEL = []string{"FATAL","ERROR","WARN","INFO","DEBUG"}

var LogLevel int = DEBUG
var LogHandlers []handler.LogHandler = []handler.LogHandler {
	handler.DEFAULT_LOGHANDLER }

func AddLogHandler(aHandler *handler.LogHandler) {
	LogHandlers = append(LogHandlers,*aHandler)
}