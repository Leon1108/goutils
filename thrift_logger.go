package goutils

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
)

type ThriftLoggerRotateType int8

const (
	THRIFT_LOGGER_ROTATE_DAILY    ThriftLoggerRotateType = 1
	THRIFT_LOGGER_ROTATE_HOURLY   ThriftLoggerRotateType = 2
	THRIFT_LOGGER_ROTATE_MINUTELY ThriftLoggerRotateType = 3
)

type ThriftLogger struct {
	tType      reflect.Type           // 每个记录器只支持记录一种类型的Thrift Struct
	fileName   string                 // 日志文件名
	fileDir    string                 // 日志文件存储的目录
	rotateType ThriftLoggerRotateType // 日志文件滚动方式，按天or按小时or按分钟
}

// NewThriftLogger 实例化Thrift日志文件记录器
// @param file 日志文件名
func NewThriftLogger(tType reflect.Type, fileName, fileDir string, rotateType ThriftLoggerRotateType) *ThriftLogger {
	return &ThriftLogger{
		tType:      tType,
		fileName:   fileName,
		fileDir:    fileDir,
		rotateType: rotateType,
	}
}

// LogNow 使用当前时间记录日志
func (this *ThriftLogger) LogNow(r thrift.TStruct) error {
	return this.Log(time.Now().Unix(), r)
}

// Log 记录日志
func (this *ThriftLogger) Log(time int64, r thrift.TStruct) error {

	// TODO 判断对象类型

	// 出于安全起见，每次记录都要保证信息已持久化到磁盘，所以每次都要Open，Close
	// TODO 先看看效果如何，再看看有没有折中的方案

	logfile := fmt.Sprintf("%v%c%v", this.fileDir, os.PathSeparator, this.getLogFileName(time))
	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	transport := thrift.NewStreamTransportW(f)
	protocol := thrift.NewTBinaryProtocolTransport(transport)
	transport.Open()
	r.Write(protocol)
	transport.Flush()
	transport.Close()
	f.Close()
	return nil
}

//getLogFilePostfix 根据日志产生的时间，获取日志文件的时间戳后缀
func (this *ThriftLogger) getLogFileName(logSecTime int64) string {
	ts := time.Unix(logSecTime, 0)
	year, month, day := ts.Date()
	hr, min, _ := ts.Clock()
	timeTag := ""
	switch this.rotateType {
	case THRIFT_LOGGER_ROTATE_DAILY:
		timeTag = fmt.Sprintf("%d%02d%02d", year, month, day)
	case THRIFT_LOGGER_ROTATE_MINUTELY:
		timeTag = fmt.Sprintf("%d%02d%02d%02d%02d", year, month, day, hr, min)
	case THRIFT_LOGGER_ROTATE_HOURLY:
		fallthrough
	default:
		timeTag = fmt.Sprintf("%d%02d%02d%02d", year, month, day, hr)
	}

	return fmt.Sprintf("%v.%v", this.fileName, timeTag)
}

type ThriftLogReader struct {
	file string
}

func NewThriftLogReader(file string) *ThriftLogReader {
	return &ThriftLogReader{
		file: file,
	}
}

func (this *ThriftLogReader) Read(processor func(protocol *thrift.TBinaryProtocol) error) error {

	f, err := os.OpenFile(this.file, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	transport := thrift.NewStreamTransportR(f)
	protocol := thrift.NewTBinaryProtocolTransport(transport)
	transport.Open()

	for {
		if err := processor(protocol); err != nil {
			break
		}
	}

	transport.Close()
	f.Close()
	return nil
}
