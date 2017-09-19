package handler

// Formatter ...
type Formatter func(message string, level string, file string, method string,
	line int) (string, error)

// LogHandler ...
type LogHandler interface {
	Level() (int, error)
	Format(message string, level string, file string,
		method string, line int) (string, error)
	Write(content string) error
	Unlink() error
}