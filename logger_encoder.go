package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// initEncoder 初始化 Encoder
func (e *LogOptions) initEncoder() zapcore.EncoderConfig {
	var encoderConfig zapcore.EncoderConfig

	if e.runEnv == RunEnvDebug {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.MessageKey = "msg"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.CallerKey = "file" // 是否输出文件调用位置
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) { //一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
		enc.AppendInt64(int64(d) / 1000000)
	}

	// 是否输出颜色
	if e.ColourOutput && !e.JSONFormat {
		//基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	return encoderConfig
}
