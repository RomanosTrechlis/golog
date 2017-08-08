package golog

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateLogger(t *testing.T) {
	Convey("Given a new logger wrapper", t, func() {
		myLogger := New()
		defer myLogger.Terminate()
		err := myLogger.New(os.Stdout, INFO, 0)
		l := len(myLogger.loggers)
		So(err, ShouldBeEmpty)
		So(l, ShouldEqual, 1)
	})

	Convey("Given a logger wrapper with unacceptable logging level", t, func() {
		myLogger := New()
		defer myLogger.Terminate()
		err := myLogger.New(os.Stdout, 15, 0)
		So(err, ShouldNotBeEmpty)
		So(typeof(err), ShouldEqual, typeof(*new(LevelTypeErr)))
	})

	Convey("Given a logger wrapper with multiple loggers attache", t, func() {
		myLogger := New()
		defer myLogger.Terminate()
		myLogger.New(os.Stdout, INFO, 0)
		err := myLogger.New(os.Stdout, ERROR, 0)
		So(err, ShouldBeEmpty)
		l := len(myLogger.loggers)
		So(l, ShouldEqual, 2)
	})
}

func TestTerminate(t *testing.T) {
	Convey("Given a logger wrapper with a logger attached", t, func() {
		myLogger := New()
		myLogger.New(os.Stdout, INFO, 0)
		Convey("When logger wrapper is active", func() {
			l := len(myLogger.loggers)
			So(l, ShouldEqual, 1)
		})

		Convey("When logger wrapper is terminated", func() {
			myLogger.Terminate()
			l := len(myLogger.loggers)
			So(l, ShouldEqual, 0)
		})
	})
}

func TestMultipleLogWrappers(t *testing.T) {
	Convey("Given two logger wrappers", t, func() {
		myLog := New()
		defer myLog.Terminate()
		myLog2 := New()
		defer myLog2.Terminate()

		Convey("When the first wrapper has one logger attached", func() {
			myLog.New(os.Stdout, TRACE, 0)
			So(len(myLog.loggers), ShouldEqual, 1)
		})
		Convey("When the seccond wrapper has two loggers attached", func() {
			myLog2.New(os.Stdout, TRACE, log.Ldate|log.Ltime|log.Llongfile)
			myLog2.New(os.Stdout, TRACE, 0)
			So(len(myLog2.loggers), ShouldEqual, 2)
		})
	})
}

func TestWrite(t *testing.T) {
	Convey("Given a logger wrapper", t, func() {
		v := os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
		lw := &LogWrapper{
			loggers: make([]Logger, 1),
		}
		lg := &logger{
			w:        v,
			inChan:   make(chan *message, 2),
			quitChan: make(chan struct{}),
			minLevel: TRACE,
			Logger:   log.New(v, "", 0),
			mainLog:  lw,
		}
		lw.loggers[0] = lg

		Convey("When a logger logs a trace", func() {
			lw.Trace("Test TRACE")
			res := <-lg.inChan
			So(res.Body, ShouldEndWith, "Test TRACE")
		})
		Convey("When a logger logs a debug", func() {
			lw.Debug("Test DEBUG")
			res := <-lg.inChan
			So(res.Body, ShouldEndWith, "Test DEBUG")
		})
		Convey("When a logger logs an info", func() {
			lw.Info("Test INFO")
			res := <-lg.inChan
			So(res.Body, ShouldEndWith, "Test INFO")
		})
		Convey("When a logger logs a warning", func() {
			lw.Warn("Test WARN")
			res := <-lg.inChan
			So(res.Body, ShouldEndWith, "Test WARN")
		})
		Convey("When a logger logs an error", func() {
			lw.Error("Test ERROR")
			res := <-lg.inChan
			So(res.Body, ShouldEndWith, "Test ERROR")
		})
		Convey("When a logger logs a fatal error", func() {
			lw.Fatal("Test FATAL")
			res := <-lg.inChan
			So(res.Body, ShouldEndWith, "Test FATAL")
		})
		lg.deleteLogger()
	})
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
