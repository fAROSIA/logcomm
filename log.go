package logcomm

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logLevel int

// InitPac Initiate logger
func InitPac(logPath string) {
	initFolder(logPath)
	initCommonLog(logPath)
	go initFileWatcher(logPath)
}

// ChangeLevel change logs' level
func ChangeLevel(level int) {
	logLevel = level
	changeLevel(level)
}

// CError log error
func CError(msg string, args ...interface{}) {
	if isAllowed(Error) {
		clog(Error, msg, args...)
	}
}

// CWarn log warn
func CWarn(msg string, args ...interface{}) {
	if isAllowed(Warn) {
		clog(Warn, msg, args...)
	}
}

// CInfo log info
func CInfo(msg string, args ...interface{}) {
	if isAllowed(Info) {
		clog(Info, msg, args...)
	}
}

// CDebug log debug
func CDebug(msg string, args ...interface{}) {
	if isAllowed(Debug) {
		clog(Debug, msg, args...)
	}
}

// CTrace log trace
func CTrace(msg string, args ...interface{}) {
	if isAllowed(Trace) {
		clog(Trace, msg, args...)
	}
}

func clog(level int, msg string, args ...interface{}) {
	str := fmt.Sprintf(msg, args...)
	switch level {
	case Fatal:
		logFatal.Log(logrus.WarnLevel, str)
	case Error:
		logError.Log(logrus.WarnLevel, str)
	case Warn:
		logWarn.Log(logrus.WarnLevel, str)
	case Info:
		logInfo.Log(logrus.WarnLevel, str)
	case Debug:
		logDebug.Log(logrus.WarnLevel, str)
	case Trace:
		logTrace.Log(logrus.WarnLevel, str)
	}
}

func initFolder(logPath string) {
	_ = os.MkdirAll(logPath, 0755)
}

func isAllowed(level int) bool {
	return level <= logLevel
}
