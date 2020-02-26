package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New return LogOptions
func New() *LogOptions {
	t := func(t bool) *bool { return &t }(true)
	return &LogOptions{
		Enabled:       t,
		LogInConsole:  true,
		ColourOutput:  true,
		runEnv:        RunEnvDebug,
		LevelSeparate: false,
		MaxSize:       100,
		MaxAge:        7, // day
		MaxBackups:    10,
		Compress:      true,
	}
}

// InitLogger 初始化 Logger
// isNew bool 是否返回一个新实例，不更改默认的Logger输出
// loga := log.New(false)
// loga.InfoFileName = "info.xxx"
// loga.ErrorFileName = "error.xxx"
// loga.JSONFormat = true
// loga.LevelSeparate = true
// loga.SetDivision(log.DivisionTime)
// loga.InitLogger()
// log.Info("SetupLogger")
// log.Error("error level test")
// log.Warn("warn level test")
// log.Debug("debug level test")
// log.Infof("info level test: %s", "111")
// log.Errorf("error level test: %s", "111")
// log.Warnf("warn level test: %s", "111")
// log.Debugf("debug level test: %s", "111")
// log.With(zap.String("Trace", "12345677")).Info("this is a log")
// log.WithError(errors.New("this is a new error")).Info("this is a log")
func (e *LogOptions) InitLogger(isNew bool) *Log {
	var (
		core    zapcore.Core
		encoder zapcore.Encoder
		wsInfo  []zapcore.WriteSyncer
		wsWarn  []zapcore.WriteSyncer
	)

	// 是否显示在控制台
	if e.LogInConsole {
		wsInfo = append(wsInfo, zapcore.AddSync(os.Stdout))
		wsWarn = append(wsWarn, zapcore.AddSync(os.Stdout))
	}

	// zapcore WriteSyncer setting
	if e.isOutput() {
		infoFileName := e.InfoFileName
		errorFileName := e.ErrorFileName
		// 日志文件切割类型
		if e.division == DivisionTime {
			// 是否输出文件根据日期输出文件
			wsInfo = append(wsInfo, zapcore.AddSync(e.getWriter(infoFileName)))
			if e.LevelSeparate {
				wsWarn = append(wsWarn, zapcore.AddSync(e.getWriter(errorFileName)))
			}

		} else if e.division == DivisionSize {
			// 是否输出文件限制大小
			wsInfo = append(wsInfo, zapcore.AddSync(e.getLumberjackWriter(infoFileName)))
			if e.LevelSeparate {
				wsWarn = append(wsWarn, zapcore.AddSync(e.getLumberjackWriter(errorFileName)))
			}
		}
	}

	// 是否 JSON 格式化
	encoderConfig := e.initEncoder()
	if e.JSONFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 是否分享日志文件
	if e.LevelSeparate {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(wsInfo...), infoLevel()),
			zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(wsWarn...), warnLevel()),
		)
	} else {
		core = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(wsInfo...), allLevel())
	}

	// 是否显示 Caller
	var logger *zap.Logger
	if e.caller {
		logger = zap.New(core, zap.AddCaller())
	} else {
		logger = zap.New(core)
	}

	// 重置标准库输出到os.Stderr。
	zap.RedirectStdLog(logger)

	// 实例化 Loger 对象
	if !isNew {
		Logger = &Log{logger}
		return Logger
	}
	return &Log{logger}
}

// isOutput whether set output file
func (e *LogOptions) isOutput() bool {
	return e.InfoFileName != ""
}

// SetCaller 设置是否显示输出
func (e *LogOptions) SetCaller(b bool) {
	e.caller = b
}

// SetDivision 设置是切割日志的方式
func (e *LogOptions) SetDivision(division DivisionType) {
	e.division = division
}

// SetRunDev 设置运行时环境
func (e *LogOptions) SetRunDev(runDev RunEnvType) {
	e.runEnv = runDev
}
