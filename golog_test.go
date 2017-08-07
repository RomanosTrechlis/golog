package golog

import (
	"fmt"
	"os"
	"testing"
	"log"
)

func TestCreateLogger(t *testing.T) {
	myLogger := New()
	defer myLogger.Terminate()
	err := myLogger.New(os.Stdout, INFO, 0)
	if err != nil {
		t.Errorf("couldn't create logger")
	}
	if len(myLogger.loggers) != 1 {
		t.Errorf("There should be one logger in the array")
	}

	err = myLogger.New(os.Stdout, 15, 0)
	if err == nil {
		t.Errorf("Expected error type 'LevelTypeErr'", )
	}
	if typeof(err) != typeof(*new(LevelTypeErr)) {
		t.Errorf("Expected error type 'LevelTypeErr'", )
	}
}

func TestTerminate(t *testing.T) {
	myLogger := New()
	defer myLogger.Terminate()
	myLogger.New(os.Stdout, INFO, 0)
	if len(myLogger.loggers) != 1 {
		t.Errorf("There should be one logger in the array")
	}
	myLogger.Terminate()
	if len(myLogger.loggers) != 0 {
		t.Errorf("There should be no logger in the array")
	}
}

func TestMultipleLogWrappers(t *testing.T) {
	myLog := New()
	defer myLog.Terminate()
	myLog.New(os.Stdout, TRACE, 0)
	myLog.Error("this is a test")

	myLog2 := New()
	defer myLog2.Terminate()
	myLog2.New(os.Stdout, TRACE, log.Ldate|log.Ltime|log.Llongfile)
	myLog2.New(os.Stdout, TRACE, 0)
	myLog.Error("this is a test 2")

	if len(myLog.loggers) != 1 {
		t.Errorf("expected one logger in 'myLog' array")
	}

	if len(myLog2.loggers) != 2 {
		t.Errorf("expected one logger in 'myLog2' array")
	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
