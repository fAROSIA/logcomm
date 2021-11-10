package logcomm

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

// NewMergedLogger return new merged logger which prints logs in single file
func NewMergedLogger(logPath string, logName string) *logrus.Logger {
	initFolder(logPath)
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	hk := newMergedHook(logName, logPath)
	logger.AddHook(hk)
	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "15:04:05"
	formatter.DisableTimestamp = false
	logger.SetFormatter(formatter)
	logger.SetLevel(Trace)

	return logger
}

// NewSeperatedLogger return new merged logger which prints logs in different files
func NewSeperatedLogger(logPath string) *logrus.Logger {
	initFolder(logPath)
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	hk := newSeperatedHook(logPath)
	logger.AddHook(hk)
	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "15:04:05"
	formatter.DisableTimestamp = false
	logger.SetFormatter(formatter)
	logger.SetLevel(Trace)
	go initFileWatcher(logger, logPath)

	return logger
}

func initFolder(logPath string) {
	_ = os.MkdirAll(logPath, 0755)
}
