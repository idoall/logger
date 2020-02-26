package main

import (
	"errors"

	"github.com/idoall/logger"
	"go.uber.org/zap"
)

func main() {
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
}
