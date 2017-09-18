package handler

import
(
	"fmt"
	"time"
)

type LogFormat func(message string, level string, file string, method string, 
	line int) string

type LogWrite func(text string)

type LogHandler struct {
	Format LogFormat
	Write LogWrite
}

var DEFAULT_LOGHANDLER = LogHandler{Format: DEFAULT_LOGFORMAT,
	Write: DEFAULT_LOGWRITE}

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

func NewLogHandler(format LogFormat, write LogWrite) (*LogHandler, error) {
	return &LogHandler{ format, write }, nil
}