package goutils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"path/filepath"
	"time"
)

const (
	DefLogDirName  = "logs"
	DefLogFileName = "app"
	DefPattern     = "%Y%m%d" //"%Y%m%d%H%M"
	DefLogFileExt  = ".log"
	IsDebug        = true // 当前是否为Debug模式
)

var log *logrus.Logger

func init() {
	var err error
	var logDir string
	if logDir, err = MkDirInWd(DefLogDirName); err != nil {
		panic(err)
	}
	if err = initLogger(logDir, DefLogFileName, DefPattern); err != nil {
		panic(err)
	}
}

func initLogger(fileDir, fileName, pattern string) (err error) {
	log = logrus.New()

	var rlog *rotatelogs.RotateLogs
	filePath := filepath.Join(fileDir, fileName)
	if rlog, err = rotatelogs.New(
		fmt.Sprintf("%s.%s%s", filePath, pattern, DefLogFileExt),
		rotatelogs.WithMaxAge(time.Hour*24*365),   // 一年
		rotatelogs.WithRotationTime(time.Hour*24), // 默认就是1天
	); err != nil {
		return
	}

	if IsDebug {
		log.SetLevel(logrus.DebugLevel)
	}

	log.AddHook(lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: rlog,
		logrus.WarnLevel:  rlog,
		logrus.InfoLevel:  rlog,
		logrus.ErrorLevel: rlog,
		logrus.FatalLevel: rlog,
		logrus.PanicLevel: rlog,
	}))
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	return
}

func SetToInfoLevel() {
	log.SetLevel(logrus.InfoLevel)
}

func Info(format string, args ...interface{}) {
	_log(logrus.InfoLevel, nil, format, args...)
}

func Debug(format string, args ...interface{}) {
	_log(logrus.DebugLevel, nil, format, args...)
}

func InfoDict(fields map[string]interface{}, format string, args ...interface{}) {
	if format == "" {
		format = "%s"
		args = []interface{}{""}
	}
	_log(logrus.InfoLevel, fields, format, args);
}

func Error(format string, args ...interface{}) {
	_log(logrus.ErrorLevel, nil, format, args...)
}

func Warn(format string, args ...interface{}) {
	_log(logrus.WarnLevel, nil, format, args...)
}
func ErrorDict(fields map[string]interface{}, format string, args ...interface{}) {
	_log(logrus.ErrorLevel, fields, format, args)
}

func _log(level logrus.Level, fields map[string]interface{}, format string, args ...interface{}) {
	//var method string
	//if pc, _, _, ok := runtime.Caller(2); ok {
	//	method = runtime.FuncForPC(pc).Name()
	//}

	// create fields
	logFields := logrus.Fields{
		//LogFieldNameCategory: category,
	}
	//if category != LogCategoryAccess {
	//	logFields[LogFieldNameFunc] = method
	//}
	for k, v := range fields {
		logFields[k] = v
	}
	logger := log.WithFields(logFields)

	switch level {
	case logrus.InfoLevel:
		logger.Infof(format, args...)
	case logrus.DebugLevel:
		logger.Debugf(format, args...)
	case logrus.ErrorLevel:
		logger.Errorf(format, args...)
	case logrus.WarnLevel:
		logger.Warnf(format, args...)
	case logrus.FatalLevel:
		logger.Fatalf(format, args...)
	case logrus.PanicLevel:
		logger.Panicf(format, args...)

	}
}
