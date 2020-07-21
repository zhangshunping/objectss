package utils

import (

	"github.com/sirupsen/logrus"
	"os"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"time"
)

var Log = logrus.New()

func init() {
	// 为当前logrus实例设置消息的输出,同样地,
	// 可以设置logrus实例的输出到任意io.writer
	logfile,_:=os.OpenFile("logs/objectss.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	Log.SetOutput(logfile)
	// 为当前logrus实例设置消息输出格式为json格式.
	// 同样地,也可以单独为某个logrus实例设置日志级别和hook,这里不详细叙述.
//	Log.Formatter = &logrus.JSONFormatter{}
	//Log.Formatter = &logrus.TextFormatter{}
	Log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		TimestampFormat: time.RFC3339,
		TrimMessages: true,
		NoColors: true,
	})
	// 日志等级
	Log.SetLevel(logrus.InfoLevel)
	// 日志输出 执行的程序函数名和路径
	Log.SetReportCaller(true)
}








