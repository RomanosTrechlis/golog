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
	mainLog  *LogWrapper
}

func (l *LogWrapper) newLogger(w io.Writer, minLevel Level, flag int) *logger {
	lg := &logger{
		w:        w,
		inChan:   make(chan *message),
		quitChan: make(chan struct{}),
		minLevel: minLevel,
		Logger:   log.New(w, "", flag),
		mainLog:  l,
	}
	return lg
}

func (l *logger) Level() Level {
	return l.minLevel
}

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
	var lArray []Logger
	for i, logger := range l.mainLog.loggers {
		if logger == l.mainLog.loggers[i] {
			continue
		}
		lArray = append(lArray, l)
	}
	l.mainLog.loggers = lArray
}

func (l *logger) Stop() {
	l.quitChan <- struct{}{}
	//<-l.quitChan

	close(l.inChan)
	close(l.quitChan)
}

// InChan is exported to interface in order to send messages
func (l *logger) InChan() chan *message {
	return l.inChan
}

// QuitChan is exported to interface in order to send messages
func (l *logger) QuitChan() chan struct{} {
	return l.quitChan
}
