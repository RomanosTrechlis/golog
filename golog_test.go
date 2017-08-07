package golog

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateLogger(t *testing.T) {
	defer Terminate()
	err := New(os.Stdout, INFO, 0)
	if err != nil {
		t.Errorf("couldn't create logger")
	}
	if len(loggers) != 1 {
		t.Errorf("There should be one logger in the array")
	}

	err = New(os.Stdout, 15, 0)
	if err == nil {
		t.Errorf("Expected error type 'LevelTypeErr'", )
	}
	if typeof(err) != typeof(*new(LevelTypeErr)) {
		t.Errorf("Expected error type 'LevelTypeErr'", )
	}
}

func TestTerminate(t *testing.T) {
	defer Terminate()
	New(os.Stdout, INFO, 0)
	if len(loggers) != 1 {
		t.Errorf("There should be one logger in the array")
	}
	Terminate()
	if len(loggers) != 0 {
		t.Errorf("There should be no logger in the array")
	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
