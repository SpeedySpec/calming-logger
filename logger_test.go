package logger

import (
	"testing"
)

func TestDefaultLoggerLogsWarnings(t *testing.T) {
	logger := DefaultLogger()
	if logger.config.LevelToLog != ERROR {
		t.FailNow()
	}
}
