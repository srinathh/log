package log_test

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"runtime"

	"github.com/srinathh/log"
)

// TSVLogWriter demonstrates writing Go log in a tab separated variable format with a custom
// OutputFn. Here we are simply writing to a writer here but we could have any amount
// of complexity - eg. spawning a goroutine to to write to a remote logging system.
type TSVLogWriter struct{ w io.Writer }

func (tl *TSVLogWriter) Output(calldepth int, s string) error {
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	_, err := fmt.Fprintf(tl.w, "%s\t%d\t%s\n", filepath.Base(file), line, s)
	return err
}

func ExampleOutputFn() {
	var buf bytes.Buffer
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("")
	log.SetOutput(&buf)
	log.Print("Standard Logging")
	customlogger := &TSVLogWriter{&buf}
	log.SetOutputFn(customlogger.Output)
	log.Print("Tab Separated Logging")
	log.SetDefOutputFn()
	log.Print("Standard Logging")

	fmt.Print(&buf)
	// Output:
	// customlogger_test.go:33: Standard Logging
	// customlogger_test.go	36	Tab Separated Logging
	// customlogger_test.go:38: Standard Logging
}
