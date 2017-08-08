package golog

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLevelTypeError(t *testing.T) {
	Convey("Given a new LevelTypeErr", t, func() {
		s := LevelTypeErr{level: TRACE}.Error()
		So(s, ShouldEqual, "[TRACE] is not an acceptable value")
	})

	Convey("Given a new LevelTypeErr", t, func() {
		s := LevelTypeErr{level: INFO}.Error()
		So(s, ShouldEqual, "[INFO] is not an acceptable value")
	})
}
