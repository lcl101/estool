package mylog

import "github.com/yandaren/go-slog/slog"

var logger = slog.NewStdoutLoggerSt("stdout_logger")

// const (
// 	DEBUG int = iota
// 	INFO
// 	WARN
// 	ERROR
// )

// type Logger struct {
// 	// File     string
// 	category string
// 	level    int
// }

// var logger Logger

// func SetLevel(level int) {
// 	logger = Logger{
// 		level:    level,
// 		category: "debug",
// 	}
// }

// func (logger *Logger) write(msg ...interface{}) {
// 	// 创建一个日志对象
// 	l := log.New(os.Stdout, "logger.category", log.LstdFlags)

// 	s := make([]interface{}, 1)
// 	s[0] = "[" + logger.category + "]"
// 	msg = append(s, msg)

// 	l.Println(msg...)
// 	logger.category = ""
// }

// func (logger *Logger) Debug(msg ...interface{}) {
// 	if logger.level > DEBUG {
// 		return
// 	}
// 	logger.category = "debug"
// 	logger.write(msg...)
// }

// func (logger *Logger) Info(msg ...interface{}) {
// 	if logger.level > INFO {
// 		return
// 	}
// 	logger.category = "info"
// 	logger.write(msg...)
// }

// func (logger *Logger) Warn(msg ...interface{}) {
// 	if logger.level > WARN {
// 		return
// 	}
// 	logger.category = "warn"
// 	logger.write(msg...)
// }

// func (logger *Logger) Error(msg ...interface{}) {
// 	logger.category = "warn"
// 	logger.write(msg...)
// }

func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}
