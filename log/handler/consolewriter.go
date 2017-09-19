package handler

import
(
	"fmt"
)
	

// ConsoleLogger ...
type consoleLogger struct {
	level  int
	format Formatter
}

// NewConsoleLogger ...
func NewConsoleLogger(level int, format Formatter) LogHandler {
	return &consoleLogger{ level, format }
}

// Level ...
func (l *consoleLogger) Level() (int, error) {
	return l.level, nil
}

// Unlink ...
func (l *consoleLogger) Unlink() error { return nil }

// Format ...
func (l *consoleLogger) Format(message string, level string, file string,
	method string, line int) (string, error) {

	return l.format(message, level, file, method, line)
}

// Write ...
func (l *consoleLogger) Write(content string) error {
	_, err := fmt.Println(content)
	return err
}