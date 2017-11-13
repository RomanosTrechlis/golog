package golog_test

import (
	"fmt"

	"github.com/RomanosTrechlis/golog"
)

// In the example we can't be certain of the order that the writes
// will be executed and so the dummy implementation for the io.Writer
// do not contain any indication of the underlying struct.
func ExampleNewWriterWrapper() {
	w := golog.NewWriterWrapper(writer1{}, writer2{})
	_, _ = w.Write([]byte("test"))

	// Output: Writer: test
	// Writer: test
}

type writer1 struct{}
type writer2 struct{}

func (w writer1) Write(p []byte) (n int, err error) {
	fmt.Println("Writer: " + string(p))
	return len(p), nil
}
func (w writer2) Write(p []byte) (n int, err error) {
	fmt.Println("Writer: " + string(p))
	return len(p), nil
}
