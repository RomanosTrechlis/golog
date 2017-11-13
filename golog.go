package golog

import (
	"io"
	"log"
	"sync"
)

type writerWrapper struct {
	writers []io.Writer
}

// NewWriterWrapper creates a wrapper struct implementing the io.Writer interface,
// containing an array of io.Writers.
func NewWriterWrapper(writers ...io.Writer) io.Writer {
	return writerWrapper{writers}
}

// NewLogger returns a logger that can have multiple io.Writer.
func NewLogger(prefix string, flag int, writers ...io.Writer) *log.Logger {
	return log.New(NewWriterWrapper(writers...), prefix, flag)
}

// Write implements the io.Writer interface for the writerWrapper struct.
//
// For every io.Writer wrapped, it executes a goroutine that Writes the
// passed. Errors are return to a channel. The Write method returns the
// first error passed into the channel.
func (w writerWrapper) Write(p []byte) (n int, err error) {
	n = len(p)

	var wg sync.WaitGroup
	wg.Add(len(w.writers))

	errChan := make(chan error, len(w.writers))
	for _, writer := range w.writers {
		go write(&wg, writer, p, errChan)
	}

	wg.Wait()

	select {
	case err = <-errChan:
		return n, err
	default:
		return n, nil
	}
}

func write(wg *sync.WaitGroup, writer io.Writer, p []byte, errChan chan error) {
	defer wg.Done()
	_, err := writer.Write(p)
	if err != nil {
		errChan <- err
	}
}
