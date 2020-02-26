package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zap 日志中的级别转化
// zap.NewAtomicLevelAt(zap.DebugLevel)
// zap.DebugLevel
// zap.InfoLevel
// zap.WarnLevel
// zap.ErrorLevel

// infoLevel Log 级别
func infoLevel() zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})
}

// warnLevel Log 级别
func warnLevel() zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
}

// allLevel Log 级别
func allLevel() zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.ErrorLevel
	})
}
