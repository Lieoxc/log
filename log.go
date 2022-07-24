package logs

import (
	"fmt"
)

func Info(format string, v ...interface{}) {
	log.logWrite.Info(fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	log.logWrite.Error(fmt.Sprintf(format, v...))

}
