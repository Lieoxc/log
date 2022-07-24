package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const callerSkipOffset = 1

func NewZapWriter(opts ...zap.Option) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		zap.InfoLevel,
	)
	// 开启开发模式，堆栈跟踪
	opts = append(opts, zap.AddCaller())
	// 设置初始化字段
	opts = append(opts, zap.AddCallerSkip(callerSkipOffset))
	// 开启文件及行号
	opts = append(opts, zap.Development())

	logger := zap.New(core, opts...)

	return logger
}
