package golog

import (
	"fmt"
	"io"
	"log"
)

// Level signifies the severity or intensity of the log event.
type Level int

const (
	// TRACE is used for when we are "tracing" the code
	// and trying to find one part of a function specifically.
	TRACE Level = iota
	// DEBUG is used to log information that is diagnostically
	// helpful to people more than just developers.
	DEBUG
	// INFO is used for generally useful information.
	INFO
	// WARN is used for anything that can potentially cause application oddities.
	WARN
	// ERROR is used for any error which is fatal to the operation,
	// but not the service or application.
	ERROR
	// FATAL is used for any error that is forcing a shutdown
	// of the service or application to prevent data loss.
	FATAL
)

type message struct {
	Level Level
	Body  string
}

var formats = map[Level]string{
	TRACE: "[TRACE] ",
	INFO:  "[INFO] ",
	WARN:  "[WARN] ",
	ERROR: "[ERROR] ",
	FATAL: "[FATAL] ",
}

// New creates a new logger, appends it to the loggers array,
// and begins a goroutine running the logger.
func New(w io.Writer, minLevel Level, flag int) CreateLoggerErr {
	if flag == 0 {
		flag = log.Ldate|log.Ltime|log.Lshortfile
	}
	if minLevel < TRACE || minLevel > FATAL {
		return LevelTypeErr{level: minLevel}
	}
	l := newLogger(w, minLevel, flag)
	loggers = append(loggers, l)
	go l.Start()
	return nil
}

func (l *Level) String() string {
	return fmt.Sprintf("%s", l)
}

// Terminate stops all logger goroutines and deletes them from loggers array
func Terminate() {
	for _, l := range loggers {
		logger := l
		logger.deleteLogger()
		logger.Stop()
	}
}

func write(level Level, format string, v ...interface{}) {
	m := &message{
		Level: level,
		Body:  formats[level] + fmt.Sprintf(format, v...),
	}
	for i := range loggers {
		l := loggers[i]
		if l.Level() > level {
			continue
		}
		l.inChan <- m
	}
}

// Trace sends write request to loggers.
func Trace(format string, v ...interface{}) {
	write(TRACE, format, v...)
}

// Info sends write request to loggers.
func Info(format string, v ...interface{}) {
	write(INFO, format, v...)
}

// Warn sends write request to loggers.
func Warn(format string, v ...interface{}) {
	write(WARN, format, v...)
}

// Error sends write request to loggers.
func Error(format string, v ...interface{}) {
	write(ERROR, format, v...)
}

// Fatal sends write request to loggers.
func Fatal(format string, v ...interface{}) {
	write(FATAL, format, v...)
}
