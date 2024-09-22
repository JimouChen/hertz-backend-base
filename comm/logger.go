package comm

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strconv"
)

// Logger 以后追加新的日志logger在此添加即可
var Logger zerolog.Logger
var MysqlLogger zerolog.Logger

// InitLogger 设置多个日志在此处
func InitLogger() error {
	err := SetGlobalLogger()
	if err != nil {
		return err
	}
	err = SetMysqlLogger()
	if err != nil {
		return err
	}
	return nil
}

func formatOutputPath(pc uintptr, file string, line int) string {
	wd, err := os.Getwd()
	if err != nil {
		return "unknown"
	}
	// 将文件路径转换为相对路径
	relativePath, err := filepath.Rel(wd, file)
	if err != nil {
		return "unknown"
	}
	return relativePath + ":" + strconv.Itoa(line)
}

func SetLogFile(logFileName string) (*lumberjack.Logger, error) {
	// 自定义CallerMarshalFunc来只显示相对路径
	zerolog.CallerMarshalFunc = formatOutputPath
	// set default global log level(debug)
	level, err := zerolog.ParseLevel(CfgLoader.GetString("log.level"))
	if err != nil {
		return nil, err
	}
	zerolog.SetGlobalLevel(level)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	logFile := &lumberjack.Logger{
		Filename:   logFileName,                         // 日志文件名
		MaxSize:    CfgLoader.GetInt("log.max_size"),    // 最大大小（MB）
		MaxBackups: CfgLoader.GetInt("log.max_backups"), // 最大备份数
		MaxAge:     CfgLoader.GetInt("log.max_age"),     // 最大保留天数
		Compress:   true,                                // 是否压缩
	}
	return logFile, nil
}

func SetGlobalLogger() error {
	logFile, err := SetLogFile(CfgLoader.GetString("log.filename"))
	if err != nil {
		return err
	}
	Logger = zerolog.New(logFile).With().Timestamp().Caller().Logger().Level(zerolog.DebugLevel)
	Logger.Info().Str("初始化", "Global Init").Msg("初始化全局日志")
	return nil
}

func SetMysqlLogger() error {
	logFile, err := SetLogFile(CfgLoader.GetString("log.mysql.filename"))
	if err != nil {
		return err
	}
	MysqlLogger = zerolog.New(logFile).With().Timestamp().Caller().Logger().Level(zerolog.DebugLevel)
	MysqlLogger.Info().Str("初始化", "Mysql Init").Msg("初始化mysql日志")
	return nil
}
