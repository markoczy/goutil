package handler

import 
(
	"bufio"
	"fmt"
	"os"
	"github.com/markoczy/goutil/files"
)

// FileLogger ...
type fileLogger struct {
	level int
	format Formatter
	file  string
	bytes int
	files int

	curFile   *os.File
	curWriter *bufio.Writer
}

// NewFileLogger ...
func NewFileLogger(aLevel int, aFormat Formatter, aFilePath string, 
	aBytesPerFile int, aMaxFiles int) LogHandler {

	return &fileLogger{aLevel, aFormat, aFilePath, aBytesPerFile,
		aMaxFiles, nil, nil}
}

// Write ...
func (w *fileLogger) Write(aMessage string) error {
	writer, err := w.initWrite(aMessage)
	//defer file.Close()
	if err != nil {
		return err
	}

	//log.Println("Before write..")
	_, err = writer.WriteString(aMessage + "\n")
	//log.Println("After write..")
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// Unlink ...
func (w *fileLogger) Unlink() error {
	return w.closeCurFile()
}

// Level ...
func (w *fileLogger) Level() (int, error) {
	return w.level, nil
}

// Format ...
func (w *fileLogger) Format(message string, level string, file string,
	method string, line int) (string, error) {

	return w.format(message, level, file, method, line)
}

func (w *fileLogger) initWrite(aMessage string) (*bufio.Writer, error) {
	// 1. Check if size exceeded

	// a. exists?
	file, err := w.currentFile()
	if err != nil {
		return nil, err
	}
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if int(stats.Size())+len(aMessage) <= w.bytes {
		//log.Println("Size small enough..")
		return w.currentWriter()
	}

	////////////////////////////////////////////////////////////////////////////

	// 2. Copy files to upper and delete last if necessary
	//log.Println("Size too big..")
	w.closeCurFile()
	if err != nil {
		return nil, err
	}

	// a. Check if last exists and delete
	path, err := getFileName(w.file, w.files)
	if err != nil {
		return nil, err
	}

	// xxx.log.3 -> X
	exists, err := files.Exists(path)
	if err != nil {
		return nil, err
	}
	if exists {
		err = os.Remove(path)
		if err != nil {
			return nil, err
		}
	}

	//
	// xxx.log.2 -> xxx.log.3
	// ...
	// xxx.log -> xxx.log.1
	//
	oldPath := path
	for i := w.files - 1; i >= 0; i-- {
		path, err = getFileName(w.file, i)
		if err != nil {
			return nil, err
		}

		exists, err = files.Exists(path)
		if err != nil {
			return nil, err
		}
		if !exists {
			oldPath = path
			continue
		}

		//fmt.Printf("Renaming %s to %s, cycle: %d\n",path,oldPath,i)

		err = os.Rename(path, oldPath)
		if err != nil {
			return nil, err
		}
		oldPath = path
	}

	return w.currentWriter()
}

// Lazy init current file
func (w *fileLogger) currentFile() (*os.File, error) {
	if w.curFile == nil {
		//log.Println("Reopening file")
		file, err := os.OpenFile(w.file, os.O_APPEND|os.O_CREATE, 0777)
		//defer file.Close()
		if err != nil {
			return nil, err
		}
		w.curFile = file
	}
	return w.curFile, nil
}

func (w *fileLogger) closeCurFile() error {
	if w.curFile == nil {
		//		log.Println("cur file nil")
		return nil
	}
	//	log.Println("Closing current file")
	err := w.curFile.Close()
	w.curFile = nil
	w.curWriter = nil
	return err
}

func (w *fileLogger) currentWriter() (*bufio.Writer, error) {
	if w.curWriter == nil {
		file, err := w.currentFile()
		if err != nil {
			return nil, err
		}
		w.curWriter = bufio.NewWriter(file)
	}
	return w.curWriter, nil
}

func getFileName(aFile string, aOffset int) (string, error) {
	if aOffset == 0 {
		return aFile, nil
	}

	// TODO rem suffix and append to new file
	return fmt.Sprintf("%s.%d", aFile, aOffset), nil
}
