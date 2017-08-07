package golog

import (
	"testing"
)

func TestLevelTypeErr_Error(t *testing.T) {
	s := LevelTypeErr{level: TRACE}.Error()
	if s != "[TRACE] is not an acceptable value" {
		t.Error("expected '[TRACE] is not an acceptable value'")
	}
}
