package log_test

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"runtime"

	"github.com/srinathh/log"
)

// TSVOutputer demonstrates writing Go log in a tab separated variable format with a custom
// LogOutputer. Here we are simply writing to a writer here but we could have any amount
// of complexity - eg. spawning a goroutine to to write to a remote logging system.
type TSVLogOutputer struct{ w io.Writer }

func (tl *TSVLogOutputer) OutputLog(l *log.Logger, calldepth int, s string) error {
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	_, err := fmt.Fprintf(tl.w, "%s\t%d\t%s\n", filepath.Base(file), line, s)
	return err
}

func ExampleCustomlogger() {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.Lshortfile)
	logger.Print("Standard Logging")
	logger.SetLogOutputer(&TSVLogOutputer{&buf})
	logger.Print("Tab Separated Logging")
	logger.SetLogOutputer(nil)
	logger.Print("Standard Logging")

	fmt.Print(&buf)
	// Output:
	// customlogger_test.go:31: Standard Logging
	// customlogger_test.go	33	Tab Separated Logging
	// customlogger_test.go:35: Standard Logging
}
