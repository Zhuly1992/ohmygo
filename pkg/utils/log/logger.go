package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"mygo/pkg/utils"
	"os"
	"runtime"
)

type Log struct {
	Logger *logrus.Logger `json:"logger"`
}

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
}

func NewLogService(logger utils.Logger) *Log {
	switch logger.StdOut {
	case "file":
		lfsHook := NewFileSplitHook(logger.FileName)
		log.AddHook(lfsHook)

		_, err := os.Stat(logger.SavePath)
		if err != nil {
			err = os.MkdirAll(logger.SavePath, os.ModePerm)
		}
		file, err := os.OpenFile(logger.FileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err == nil {
			log.Out = file
		}
	case "redis":
		NewLogRedis(log)
	}
	return &Log{
		Logger: log,
	}
}

func (log *Log) Info(msg string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	log.Logger.WithFields(logrus.Fields{
		"file":    files,
		"message": args,
	}).Info(msg)
}

func (log *Log) Error(msg string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	files := fmt.Sprintf("%s (%d)", file, line)
	log.Logger.WithFields(logrus.Fields{
		"file":    files,
		"message": args,
	}).Error(msg)
}
