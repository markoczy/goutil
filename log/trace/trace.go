package trace

import
(
	//"fmt"
	"runtime"
	"errors"
	"strings"
	"strconv"
)

type StackTraceElement struct {
	Method string
	File string
	Line int
}


func Trace(aOffset int) (*StackTraceElement, error) {
	// Get Stack trace and trim
	buf := make([]byte, 0xFFFF)
	str := string(buf[:runtime.Stack(buf, false)])
	// Extract relevant lines
	lines := strings.Split(str,"\n")
	if len(lines) < aOffset*2 + 3 {
		return nil, errors.New("Offset too high")
	}
	// C:/GOPATH/src/sandbox/stacktrace/main.go:41 +0x27
	method := lines[aOffset*2 + 1]
	paths := strings.Split(lines[aOffset*2 + 2], "/src/")
	// stacktrace/main.go:41 +0x27
	file := paths[len(paths)-1];
	// stacktrace/main.go:41
	file = file[:strings.Index(file," ")]
	lnFile := strings.Split(file, ":")
	// stacktrace.main.go
	if len(lnFile) != 2 {
		return nil, errors.New("Parsing failed: expected ':'")
	}
	file = lnFile[0]
	// 41
	line, err := strconv.Atoi(lnFile[1])
	if err != nil {
		return nil, err
	}

	// Preview example:
	//fmt.Println("DEBUG   "+file+"[",line,"]::"+lines[aOffset*2 + 1]+":")
	return &StackTraceElement{method, file, line}, nil
}

func TraceDefault() (*StackTraceElement, error) {
	return Trace(2)
}

