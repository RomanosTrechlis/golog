package golog

import (
	"io"
	"log"
)

type logger struct {
	*log.Logger
	w        io.Writer
	minLevel Level
	flags    int
	inChan   chan *message
	quitChan chan struct{}
}

func (l *logger) Level() Level {
	return l.minLevel
}

var loggers []*logger

func (l *logger) Start() {
	for {
		select {
		case message := <-l.inChan:
			l.Logger.Print(message.Body)
		case <-l.quitChan:
			l.deleteLogger()
			return
		}
	}
}

func (l *logger) deleteLogger() {
	var lArray []*logger
	for i, l := range loggers {
		if l == loggers[i] {
			continue
		}
		lArray = append(lArray, l)
	}
	loggers = lArray
}

func (l *logger) Stop() {
	l.quitChan <- struct{}{}
	//<-l.quitChan

	close(l.inChan)
	close(l.quitChan)
}
