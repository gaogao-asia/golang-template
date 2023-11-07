package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetEncoderConfig(isHandler bool) zapcore.EncoderConfig {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = SyslogTimeEncoder
	encoderConfig.EncodeLevel = CustomLevelEncoder
	encoderConfig.StacktraceKey = ""
	encoderConfig.CallerKey = "caller"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	if isHandler {
		encoderConfig.EncodeCaller = ShortCallerEncoderForHandler
		return encoderConfig
	}

	encoderConfig.EncodeCaller = ShortCallerEncoder
	return encoderConfig
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000Z"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func ShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// TODO: consider using a byte-oriented API to save an allocation.
	call := GetFunctionNameAtRuntime(8)
	enc.AppendString(fmt.Sprintf("%s:%d", call.FilePath, call.Line))
}

func ShortCallerEncoderForHandler(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// TODO: consider using a byte-oriented API to save an allocation.
	call := GetFunctionNameAtRuntime(9)
	enc.AppendString(fmt.Sprintf("%s:%d", call.FilePath, call.Line))
}
