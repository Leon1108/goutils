package logger

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	"github.com/Leon1108/goutils"
	"github.com/astaxie/beego/logs"
	"os"
)

const (
	//DefConfigFile 默认日志配置文件名
	DefConfigFile = "conf/logger.conf"
	//DefLoggerChannelLen 当异步输出时，队列中的消息数
	DefLoggerChannelLen = 1000
	//TestingLoggerName 单元测试时使用的logger name
	TestingLoggerName = "testlogger"
	//DefLoggerName 默认日志配置名称
	DefLoggerName = "default"
)

var loggers map[string]*logs.BeeLogger

type loggerConf struct {
	Name          string
	ChanLen       int64
	Async         bool
	FuncCallDepth *funcCallDepth
	Engine        *loggerEngine
}

type funcCallDepth struct {
	Enable bool
	Depth  int
}

type loggerEngine struct {
	Name string
	Conf map[string]interface{}
}

func init() {
	if goutils.IsTesting() {
		//单元测试，使用默认配置
		conf := loggerConf{
			Name:    TestingLoggerName,
			ChanLen: 1,
			Async:   false,
			FuncCallDepth: &funcCallDepth{
				Enable: true,
				Depth:  2,
			},
			Engine: &loggerEngine{
				Name: "console",
				Conf: map[string]interface{}{
					"level": 7,
				},
			},
		}
		if err := initLoggers([]loggerConf{conf}); err != nil {
			panic(err)
		}
		return
	}

	//正常运行加载配置文件
	var confs []loggerConf
	var err error
	if confs, err = loadConfigFile(DefConfigFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[Warn] config file '%v' not found!\n", DefConfigFile)
			//默认配置，当配置文件不可用时，加载此配置
			confs = []loggerConf{
				loggerConf{
					Name:    DefLoggerName,
					ChanLen: 10,
					Async:   false,
					FuncCallDepth: &funcCallDepth{
						Enable: true,
						Depth:  2,
					},
					Engine: &loggerEngine{
						Name: "console",
						Conf: map[string]interface{}{
							"level": 7,
						},
					},
				},
			}
		} else {
			panic(err)
		}
	}

	if err = initLoggers(confs); err != nil {
		panic(err)
	}

}

//loadConfigFile 加载配置文件
func loadConfigFile(name string) (conf []loggerConf, err error) {
	var bytes []byte
	if bytes, err = ioutil.ReadFile(name); err != nil {
		return
	}
	var loggersConf []loggerConf
	err = json.Unmarshal(bytes, &loggersConf)
	return loggersConf, err
}

//initLoggers 根据日志配置初始化所有日志记录器
func initLoggers(confs []loggerConf) (err error) {
	loggers = map[string]*logs.BeeLogger{}

	for _, conf := range confs {
		logger := logs.NewLogger(conf.ChanLen)
		if conf.Async {
			logger.Async()
		}
		if conf.FuncCallDepth != nil {
			logger.EnableFuncCallDepth(conf.FuncCallDepth.Enable)
			logger.SetLogFuncCallDepth(conf.FuncCallDepth.Depth)
		}
		if conf.Engine != nil {
			var bytes []byte
			if bytes, err = json.Marshal(conf.Engine.Conf); err != nil {
				return
			}
			logger.SetLogger(conf.Engine.Name, string(bytes))
		}
		loggers[conf.Name] = logger
	}

	return
}

//GetLogger 根据名称获取指定的日志记录器
func GetLogger(name string) (*logs.BeeLogger, error) {
	if goutils.IsTesting() {
		name = TestingLoggerName
	}
	if v, ok := loggers[name]; ok {
		return v, nil
	} else if v, ok = loggers[DefLoggerName]; ok { //尝试使用默认日志
		fmt.Printf("[Warn] Logger '%v' not found! \n", name)
		return v, nil
	}

	return nil, fmt.Errorf("Logger '%v' not found!", name)
}
