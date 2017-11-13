package golog_test

import (
	"fmt"

	"log"

	"github.com/RomanosTrechlis/golog"
)

// In the example we can't be certain of the order that the writes
// will be executed and so the dummy implementation for the io.Writer
// do not contain any indication of the underlying struct.
func ExampleNewWriterWrapper() {
	w := golog.NewWriterWrapper(writer1{}, writer2{})
	_, _ = w.Write([]byte("test"))

	// Output: Writer: testWriter: test
}

func ExampleNewLogger() {
	l := golog.NewLogger("TRACE ", log.Lshortfile, golog.NewWriterWrapper(writer1{}, writer2{}))
	l.Print("test")
	// Output: Writer: TRACE golog_test.go:23: test
	// Writer: TRACE golog_test.go:23: test
}

type writer1 struct{}
type writer2 struct{}

func (w writer1) Write(p []byte) (n int, err error) {
	fmt.Print("Writer: " + string(p))
	return len(p), nil
}
func (w writer2) Write(p []byte) (n int, err error) {
	fmt.Print("Writer: " + string(p))
	return len(p), nil
}
