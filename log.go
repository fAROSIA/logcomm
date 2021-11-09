package logcomm

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logLevel int

// NewLogger return new logger, if args giv
//func NewLogger(logPath string, args ...string) {
//
//}

// InitLogger Initiate logger
func InitLogger(logPath string, args ...string) {
	initFolder(logPath)
	ChangeLevel(Info)
	// select log mode: merged or discrete
	if len(args) == 0 {
		// discrete
		initCommonLog(logPath)
		go initFileWatcher(logPath)
	} else {
		// merged
		initMergedLog(args[0], logPath)
	}
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

func MError(msg string, args ...interface{}) {
	if isAllowed(Error) {
		mlog(msg, args...)
	}
}

func MWarn(msg string, args ...interface{}) {
	if isAllowed(Warn) {
		mlog(msg, args...)

	}
}

func MInfo(msg string, args ...interface{}) {
	if isAllowed(Info) {
		mlog(msg, args...)

	}
}

func MDebug(msg string, args ...interface{}) {
	if isAllowed(Debug) {
		mlog(msg, args...)

	}
}

func MTrace(msg string, args ...interface{}) {
	if isAllowed(Trace) {
		mlog(msg, args...)

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

func mlog(msg string, args ...interface{}) {
	str := fmt.Sprintf(msg, args...)
	logMerge.Log(logrus.WarnLevel, str)
}
func initFolder(logPath string) {
	_ = os.MkdirAll(logPath, 0755)
}

func isAllowed(level int) bool {
	return level <= logLevel
}
