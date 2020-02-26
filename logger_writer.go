package logger

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/idoall/TokenExchangeCommon/commonutils"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gopkg.in/natefinch/lumberjack.v2"
)

func (e *LogOptions) getLumberjackWriter(fileName string) io.Writer {
	// 获取当前工作目录
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 设置 logs 输出文件目录
	logFolderPath := filepath.Join(workPath, "logs")

	// 如果目录不存在，则创建
	if !commonutils.PathExists(logFolderPath) {
		if err = os.Mkdir(logFolderPath, os.ModePerm); err != nil {
			panic(err)
		}
	}

	// 拼接日志目录文件
	baseLogPath := filepath.Join(logFolderPath, fileName)

	hook := lumberjack.Logger{
		Filename:   baseLogPath,  // 日志文件路径
		MaxSize:    e.MaxSize,    // MB
		MaxBackups: e.MaxBackups, // 最多保留3个备份
		Compress:   true,         // 是否压缩 disabled by default
		MaxAge:     e.MaxAge,     // days
	}
	return &hook
}

// getWriter 日志文件切割
func (e *LogOptions) getWriter(fileName string) io.Writer {
	// 获取当前工作目录
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 设置 logs 输出文件目录
	logFolderPath := filepath.Join(workPath, "logs")

	// 如果目录不存在，则创建
	if !commonutils.PathExists(logFolderPath) {
		if err = os.Mkdir(logFolderPath, os.ModePerm); err != nil {
			panic(err)
		}
	}

	// 拼接日志目录文件
	baseLogPath := filepath.Join(logFolderPath, fileName)
	newLogName := baseLogPath[:strings.LastIndex(baseLogPath, ".")]
	fileExe := baseLogPath[strings.LastIndex(baseLogPath, ".")+1:]

	// 保存30天内的日志，每24小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		newLogName+".%F."+fileExe,
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(time.Duration(int64(24*time.Hour)*int64(e.MaxAge))),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithClock(rotatelogs.Local),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
