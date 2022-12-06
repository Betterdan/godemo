package logger

import (
	"demo/common"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type logHook struct{}

var MyLog zerolog.Logger

func LoadMyLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//定义默认key名称
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "msg"
	zerolog.ErrorFieldName = "errMsg"
	zerolog.CallerFieldName = "currentLogFile"

	//设置全局日志级别
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	curErr := errors.New("测试一个新大错误")

	log.Print("test zeroLog")                                         //{"level":"debug","time":1669890534,"message":"test zeroLog"}
	log.Debug().Str("name", "dan").Float64("age", 28.5).Msg("这是一个测试") //{"level":"debug","name":"dan","age":28.5,"time":1669890534,"message":"这是一个测试"}
	log.Debug().Str("work", "kingdee").Send()                         //{"level":"debug","work":"kingdee","time":1669890534}
	log.Info().Str("version", "2.0.5").Msg("info 日志")                 //{"level":"info","version":"2.0.5","time":1669890534,"message":"info 日志"}
	log.Error().Str("version", "2.0.3").Msg("error 日志")               //{"level":"error","version":"2.0.3","time":1669890534,"message":"error 日志"}
	err := errors.New("seems we have an error here")
	log.Error().Str("version", "2.0.5").Err(err).Msg("msg信息") //{"level":"error","version":"2.0.5","error":"seems we have an error here","time":1669890534,"message":"msg信息"}
	log.Warn().Str("version", "2.0.3").Msg("warn 日志")         //{"level":"warn","version":"2.0.3","time":1669890534,"message":"warn 日志"}
	//log.Fatal().Str("version", "2.0.3").Msg("fatal 日志") //{"level":"fatal","version":"2.0.3","time":1669890534,"message":"fatal 日志"}
	//log.Panic().Str("version", "2.0.3").Msg("panic 日志")

	/** 新增一个父logger 增加一些公共日志记录 */
	logFile := getLogFile("info")
	multi := zerolog.MultiLevelWriter(logFile, os.Stdout)
	logger := zerolog.New(multi).With().Str("version", "v2.0.5").Timestamp().Logger()
	/** 注册钩子 */
	//logger.Hook(logHook{})

	//子log 用于当前项目输出，
	MyLog = logger.With().Str("company_id", "企业id").Str("自定义key", "自定义value").Caller().Logger()

	MyLog.Info().Msg("这是一个错误")

	//{"l":"info","name":"统一加名称","自定义key":"自定义value","errMsg":"测试一个新大错误","t":1669892178,"currentLogFile":"D:/go_work/demo/service/AccountService.go:60","msg":"测试自定义key名称"}
	MyLog.Info().Err(curErr).Msg("测试自定义key名称")

	MyLog.Warn().Msg("这是一个warn错误")
	MyLog.Error().Msg("这是一个error 错误")
}

func (logHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level <= zerolog.InfoLevel {
		e.Str("lev","info")
	} else if level == zerolog.WarnLevel {
		e.Str("lev","warn")
	} else {
		e.Str("lev","error")
	}
}

func getLogFile(logLevel string) *os.File {
	path, err := common.GetCurrentPath()
	if err != nil {
		fmt.Println("加载配置失败!", err)
	}
	logDir := path + "/log/"
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建日志目录失败", err)
	}

	logFileName := logDir + time.Now().Format("2006-01-02") + "_" + logLevel + ".log"
	logFile, _ := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return logFile
}
