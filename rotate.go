package logcomm

import (
	"path/filepath"
	"time"

	"github.com/rifflock/lfshook"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func newHook(logName string, logPath string) logrus.Hook {
	logName = filepath.Join(logPath, logName)
	writer, err := rotatelogs.New(
		logName+"_%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(time.Duration(10)*time.Second),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err.Error())
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	return lfsHook
}
