# logger
基于Zap扩展的日志库

可以实现按天文件切割

## Installation

`go get -u github.com/idoall/logger`

## Quick Start

使用方法

```go
log := logger.New()
log.InfoFileName = "info.log"
log.ErrorFileName = "error.log"
log.JSONFormat = false
log.LevelSeparate = true
log.SetDivision(logger.DivisionTime)
log.InitLogger(false)
logger.Info("SetupLogger")
logger.Error("error level test")
logger.Warn("warn level test")
logger.Debug("debug level test")
logger.Infof("info level test: %s", "111")
logger.Errorf("error level test: %s", "111")
logger.Warnf("warn level test: %s", "111")
logger.Debugf("debug level test: %s", "111")
logger.With(zap.String("Trace", "12345677")).Info("this is a log")
logger.WithError(errors.New("this is a new error")).Info("this is a log")
```
> 如果输出JSON格式，在控制台不会使用彩色模式显示
