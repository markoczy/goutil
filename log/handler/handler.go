package handler

import
(
	"github.com/markoczy/goutil/log/handler/filewriter"
)

type LogFormat func(message string, level string, file string, method string, 
	line int) string

type LogWrite func(text string)

type LogHandler struct {
	Level int
	Format LogFormat
	Write LogWrite
}

func NewLogHandler(level int, format LogFormat, 
	write LogWrite) (*LogHandler, error) {

	return &LogHandler{ level, format, write }, nil
}

func NewFileWriter(level int, format LogFormat, fileName string, 
	bytesMax int, filesMax int) (*LogHandler, error) {

	// Create and add LogHandler
	w, err := filewriter.New(fileName, bytesMax, filesMax)
	if err!= nil { return nil, err }
	return &LogHandler{ level, format, w.LogWriter }, nil
}