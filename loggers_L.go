package logger

import (
	"fmt"

	"go.uber.org/zap"
)

/*
	对于 Log 的方法重写，在 loggerInstance.InitLogger(true) 时，返回一个新实例时使用
*/

// Info handler takes any input returns unformatted output to infoLogger writer
func (e *Log) Info(msg string, args ...zap.Field) {
	e.L.Info(msg, args...)
}

// Infof handler takes any input infoLogger returns formatted output to infoLogger writer
func (e *Log) Infof(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	e.L.Info(logMsg)
}

// Debug handler takes any input returns unformatted output to infoLogger writer
func (e *Log) Debug(msg string, args ...zap.Field) {
	e.L.Debug(msg, args...)
}

// Debugf handler takes any input infoLogger returns formatted output to infoLogger writer
func (e *Log) Debugf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	e.L.Debug(logMsg)
}

// Panic handler takes any input returns unformatted output to infoLogger writer
func (e *Log) Panic(msg string, args ...zap.Field) {
	e.L.Panic(msg, args...)
}

// Panicf handler takes any input infoLogger returns formatted output to infoLogger writer
func (e *Log) Panicf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	e.L.Panic(logMsg)
}

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// log.With(
// 	zap.String("hello", "world"),
// 	zap.Any("failure", errors.New("oh no")),
// 	zap.Any("user", User{Name: "alice"}),
// ).Warn("111")
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func (e *Log) With(fields ...zap.Field) *zap.Logger {
	return e.L.With(fields...)
	// return zap.Any(k, v)

}

// WithError constructs a field that lazily stores err.Error() under the
// provided key. Errors which also implement fmt.Formatter (like those produced
// by github.com/pkg/errors) will also have their verbose representation stored
// under key+"Verbose". If passed a nil error, the field is a no-op.
//
// For the common case in which the key is simply "error", the Error function
// is shorter and less repetitive.
func (e *Log) WithError(err error) *zap.Logger {
	return With(zap.NamedError("error", err))
}

// Warn handlerLogger takes any input returns unformatted output to warnLogger writer
func (e *Log) Warn(msg string, args ...zap.Field) {
	e.L.Warn(msg, args...)
}

// Warnf handler takes any input returns unformatted output to warnLogger writer
func (e *Log) Warnf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	e.L.Warn(logMsg)
}

// Error handler takes any input returns unformatted output to errorLogger writer
func (e *Log) Error(msg string, args ...zap.Field) {
	e.L.Error(msg, args...)
}

// Errorf handler takes any input returns unformatted output to errorLogger writer
func (e *Log) Errorf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	e.L.Error(logMsg)
}

// Fatal handler takes any input returns unformatted output to fatalLogger writer
func (e *Log) Fatal(msg string, args ...zap.Field) {
	e.L.Fatal(msg, args...)
}

// Fatalf handler takes any input returns unformatted output to fatalLogger writer
func (e *Log) Fatalf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	e.L.Fatal(logMsg)
}
