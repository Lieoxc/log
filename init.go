package log

import (
	"sync"

	"go.uber.org/zap"
)

var once sync.Once

var logWriter *zap.Logger

func init() {
	once.Do(func() {
		logWriter = NewZapWriter()
	})
}
