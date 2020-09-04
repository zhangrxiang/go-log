package log

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	DefaultFileName         = "log-%Y-%m-%d.log"
	DefaultFileMaxAge       = time.Hour * 24 * 30
	DefaultFileRotationTime = time.Hour * 24
)

func LoadFileLog() {
	hook, err := DefaultFile()
	if err != nil {
		logrus.Fatal(err)
	}
	logger.AddHook(hook)
}

func DefaultFile() (*lfshook.LfsHook, error) {
	logf, err := rotatelogs.New(
		DefaultFileName,
		rotatelogs.WithMaxAge(DefaultFileMaxAge),
		rotatelogs.WithRotationTime(DefaultFileRotationTime),
	)
	if err != nil {
		return nil, err
	}
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  logf,
		logrus.WarnLevel:  logf,
		logrus.ErrorLevel: logf,
		logrus.FatalLevel: logf,
		logrus.PanicLevel: logf,
	}, &logrus.TextFormatter{
		DisableTimestamp: false,
		FullTimestamp:    true,
		ForceQuote:       true,
		TimestampFormat:  DefaultLocalTimeDateFormat,
		ForceColors:      true,
	}), nil
}
