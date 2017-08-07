package golog

import "fmt"

// CreateLoggerErr defines errors while creating a logger
type CreateLoggerErr error

func (err LevelTypeErr) Error() string {
	return fmt.Sprintf("%sis not an acceptable value", formats[err.level])
}

// LevelTypeErr defines log level outside boundaries.
type LevelTypeErr struct {
	level Level
}
