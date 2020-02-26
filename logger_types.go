package logger

import (
	"go.uber.org/zap"
)

// LogOptions struct that holds all user configurable options for the logger
type LogOptions struct {
	Enabled       *bool  `json:"enabled,omitempty"`
	LogInConsole  bool   `json:"logconsole"`                                                 // 是否输出到控制台
	InfoFileName  string `json:"info_filename"`                                              // 输出到 Info Log的文件，如果不配置，不输出到文件
	ErrorFileName string `json:"error_filename" yaml:"error_filename" toml:"error_filename"` // 错误日志名称
	MaxSize       int    `json:"max_size" yaml:"max_size" toml:"max_size"`                   // 文件大小，单位 MB，默认 100MB
	MaxBackups    int    `json:"max_backups" yaml:"max_backups" toml:"max_backups"`          // 保留的最大旧日志文件数。默认是 10
	MaxAge        int    `json:"max_age" yaml:"max_age" toml:"max_age"`                      // 日志保留最大天数,默认7天
	Compress      bool   `json:"compress" yaml:"compress" toml:"compress"`                   // 是否压缩
	JSONFormat    bool   `json:"jsonformat"`                                                 // 是否 JSON 格式化日志格式
	LevelSeparate bool   `json:"level_separate" yaml:"level_separate" toml:"level_separate"` // 不同级别是否分离
	ColourOutput  bool   `json:"colour"`

	caller   bool         // 是否显示调用文件位置
	division DivisionType // 切割日志的形式 "Time" 按时间 "Size" 按大小
	runEnv   RunEnvType   // 运行环境
}

// Log struct
type Log struct {
	L *zap.Logger
}

var (
	// Logger create a pointer to Logging struct for holding data
	Logger *Log
)

const (
	// TimeDivision 按时间切割日志
	TimeDivision = "time"
	// SizeDivision 按大小切割日志
	SizeDivision = "size"
)

// Level type
type Level int8

const (
	// DebugLevel 日志通常很庞大，通常在生产。
	DebugLevel Level = iota - 1

	// InfoLevel 是默认的日志记录优先级。
	InfoLevel

	// WarnLevel 日志比Info更重要，但不需要个人
	//人工审核。
	WarnLevel

	// ErrorLevel 日志是高优先级的。如果应用程序运行顺利，
	//它不应生成任何错误级别的日志
	ErrorLevel

	//DPanicLevel 日志是特别重要的错误。在开发中
	//编写消息后记录器出现紧急情况。
	DPanicLevel

	// PanicLevel 记录一条消息，然后出现紧急情况。
	PanicLevel

	// FatalLevel 记录一条消息，然后调用os.Exit（1）。
	FatalLevel
)

// DivisionType 切割类型
type DivisionType string

const (
	// DivisionTime 按时间
	DivisionTime = DivisionType("Time")
	// DivisionSize 按大小
	DivisionSize = DivisionType("Size")
	// DivisionNone 不写入文件
	DivisionNone = DivisionType("None")
)

// RunEnvType 运行环境
type RunEnvType string

const (
	// RunEnvDebug Debug
	RunEnvDebug = RunEnvType("Debug")
	// RunEnvProd Prod
	RunEnvProd = RunEnvType("Prod")
)
