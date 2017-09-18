package filewriter

import
(
	"fmt"
	"os"
	"bufio"
	//"log"
	"github.com/markoczy/goutil/files"
)


type LogFileWriter struct {
	file string
	bytes int
	files int

	curFile *os.File
	curWriter  *bufio.Writer
}

func New(aFilePath string, aBytesPerFile int, 
	aMaxFiles int) (*LogFileWriter, error) {

	return &LogFileWriter{aFilePath,aBytesPerFile,aMaxFiles, nil, nil}, nil
}

func (w *LogFileWriter)  LogWriter (aMessage string) {
	w.Write(aMessage)
}

func (w *LogFileWriter) Write(aMessage string) error {
	writer, err := initWrite(w, aMessage)
	//defer file.Close()
	if err != nil { return err }

	//log.Println("Before write..")
	_, err = writer.WriteString(aMessage+"\n")
	//log.Println("After write..")
	if err != nil { return err }
	
	err = writer.Flush()
	if err != nil { return err }
	
	return nil
}

func initWrite(aWriter *LogFileWriter, aMessage string) (*bufio.Writer,error) {
	// 1. Check if size exceeded
	
	// a. exists?
	file, err := aWriter.currentFile()
	if err != nil {	return nil, err }
	stats, err := file.Stat()
	if err != nil {	return nil, err }

	if int(stats.Size()) + len(aMessage) <= aWriter.bytes { 
		//log.Println("Size small enough..")
		return aWriter.currentWriter()
	}


	////////////////////////////////////////////////////////////////////////////

	// 2. Copy files to upper and delete last if necessary
	//log.Println("Size too big..")
	aWriter.closeCurFile()
	if err!=nil  { return nil, err }

	// a. Check if last exists and delete
	path, err := getFileName(aWriter.file, aWriter.files)
	if err!=nil  { return nil, err }

	
	// xxx.log.3 -> X
	exists, err := files.Exists(path)
	if err!=nil  { return nil, err }
	if exists {
		err = os.Remove(path) 
		if err!=nil { return nil, err }
	}

	//
	// xxx.log.2 -> xxx.log.3
 	// ...
	// xxx.log -> xxx.log.1
	//
	oldPath := path
	for i:=aWriter.files-1; i>=0; i-- {
		path, err = getFileName(aWriter.file, i)
		if err!=nil  { return nil, err }

		exists, err = files.Exists(path)
		if err!=nil  { return nil, err }
		if !exists { 
			oldPath = path
			continue
		}

		//fmt.Printf("Renaming %s to %s, cycle: %d\n",path,oldPath,i)
		
		err = os.Rename(path, oldPath)
		if err != nil { return nil, err }
		oldPath = path
	}
	
	return aWriter.currentWriter()
}

// Lazy init current file
func (w *LogFileWriter) currentFile() (*os.File, error) {
	if w.curFile == nil {
		//log.Println("Reopening file")
		file, err := os.OpenFile(w.file, os.O_APPEND|os.O_CREATE, 0777)
		//defer file.Close()
		if err != nil { return nil, err }
		w.curFile = file
	}
	return w.curFile, nil
}

func (w *LogFileWriter) closeCurFile() (error) {
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

func (w *LogFileWriter) currentWriter() (*bufio.Writer, error) {
	if w.curWriter == nil {
		file,err := w.currentFile()
		if err != nil { return nil, err }
		w.curWriter = bufio.NewWriter(file)
	}
	return w.curWriter, nil
}

// func openFile(aFile string) (*os.File, error) {
// 	file, err := os.OpenFile(aFile, os.O_APPEND|os.O_CREATE, 0777)
// 	defer file.Close() // failing ?
// 	return file, err
// }

func getFileName(aFile string, aOffset int) (string, error) {
	if aOffset == 0 {
		return aFile, nil
	}

	// TODO rem suffix and append to new file
	return fmt.Sprintf("%s.%d",aFile,aOffset), nil
}