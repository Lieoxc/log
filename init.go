package logs

import (
	"sync"

	"go.uber.org/zap"
)

type logger struct {
	logWrite *zap.Logger
}

var once sync.Once

var log logger

func init() {
	once.Do(func() {
		log.logWrite = NewZapWriter()
	})
}
