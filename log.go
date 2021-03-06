package log

import (
	"fmt"
)

func Debug(format string, v ...interface{}) {
	logWriter.Debug(fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	logWriter.Info(fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	logWriter.Error(fmt.Sprintf(format, v...))
}
