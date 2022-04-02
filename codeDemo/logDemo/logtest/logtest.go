package logtest

import (
	"bytes"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"sync"
	"time"
)

var logInstance *logrus.Logger

var once = &sync.Once{}

func Instance() *logrus.Logger {
	once.Do(func() {
		logInstance = newLogger()
	})
	return logInstance
}

type MyFormatter struct{}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d %s] %s\n",
			timestamp, entry.Level, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

//New第一个参数是文件名后缀命名格式
//WithLinkName：软连接
//WithMaxAge：设置文件清理前最长保存时间
//WithRotationTime：设置日志分割的时间，隔多久分割一次
//WithRotationCount: 设置文件清理前最多保存的个数
//WithMaxAge 和 WithRotationCount 二者只能设置一个
//以下配置日志每隔一天秒轮转一个新文件，保留近7天的日志文件
//默认至少记录1分钟的内容，
func newLogger() *logrus.Logger {
	var log *logrus.Logger = logrus.New()
	filepaths := "E:/gowork/src/logDemo/log/project.log"
	writer, _ := rotatelogs.New(
		filepaths+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(filepaths),
		rotatelogs.WithMaxAge(time.Duration(60*60*24*7)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60*60*24)*time.Second),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}
	log.SetReportCaller(true)
	lfHook := lfshook.NewHook(writeMap, &MyFormatter{})
	log.AddHook(lfHook)
	return log
}
