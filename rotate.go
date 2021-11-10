package logcomm

import (
	"path"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"github.com/rifflock/lfshook"

	"github.com/sirupsen/logrus"
)

func newMergedHook(logName string, logPath string) logrus.Hook {
	logName = filepath.Join(logPath, logName)
	file := logName + ".log"

	lfsHook := lfshook.NewHook(lfshook.PathMap{
		logrus.PanicLevel: file,
		logrus.FatalLevel: file,
		logrus.WarnLevel:  file,
		logrus.ErrorLevel: file,
		logrus.InfoLevel:  file,
		logrus.DebugLevel: file,
		logrus.TraceLevel: file,
	}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	return lfsHook
}

func newSeperatedHook(logPath string) logrus.Hook {
	pWriter, err := rotatelogs.New(
		path.Join(logPath, "panic")+"_%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		panic("initiate log failed :" + err.Error())
	}
	fWriter, err := rotatelogs.New(
		path.Join(logPath, "fatal")+"_%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		panic("initiate log failed :" + err.Error())
	}
	eWriter, err := rotatelogs.New(
		path.Join(logPath, "error")+"_%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		panic("initiate log failed :" + err.Error())
	}
	wWriter, err := rotatelogs.New(
		path.Join(logPath, "warn")+"_%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		panic("initiate log failed :" + err.Error())
	}
	iWriter, err := rotatelogs.New(
		path.Join(logPath, "info")+"_%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		panic("initiate log failed :" + err.Error())
	}
	dWriter, err := rotatelogs.New(
		path.Join(logPath, "debug")+"_%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		panic("initiate log failed :" + err.Error())
	}
	tWriter, err := rotatelogs.New(
		path.Join(logPath, "trace")+"_%Y-%m-%d"+".log",
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		panic("initiate log failed :" + err.Error())
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.PanicLevel: pWriter,
		logrus.FatalLevel: fWriter,
		logrus.WarnLevel:  wWriter,
		logrus.ErrorLevel: eWriter,
		logrus.InfoLevel:  iWriter,
		logrus.DebugLevel: dWriter,
		logrus.TraceLevel: tWriter,
	}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	return lfsHook
}
