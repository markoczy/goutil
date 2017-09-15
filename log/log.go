package log

import
(
	"fmt"
	"github.com/markoczy/goutil/log/trace"
)

func Debug(aMessage string) {
	t, err := trace.TraceDefault()
	if err != nil {
		// TODO
	}
	fmt.Println(t.Method)
}

func Info() {

}