package logcomm

import (
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

var (
	logFatal = logrus.New()
	logError = logrus.New()
	logWarn  = logrus.New()
	logInfo  = logrus.New()
	logDebug = logrus.New()
	logTrace = logrus.New()
)

var loggerMap = map[string]*logrus.Logger{
	"fatal": logFatal,
	"error": logError,
	"warn":  logWarn,
	"info":  logInfo,
	"debug": logDebug,
	"trace": logTrace,
}

func initCommonLog(logPath string) {
	// clear logger's default output
	logFatal.SetOutput(ioutil.Discard)
	logError.SetOutput(ioutil.Discard)
	logWarn.SetOutput(ioutil.Discard)
	logInfo.SetOutput(ioutil.Discard)
	logDebug.SetOutput(ioutil.Discard)
	logTrace.SetOutput(ioutil.Discard)

	// make new flshook for each logger
	for l, lg := range loggerMap {
		hk := newHook(l, logPath)
		lg.AddHook(hk)
	}

	// set formatter
	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "15:04:05"
	formatter.DisableTimestamp = false
	logFatal.SetFormatter(formatter)
	logError.SetFormatter(formatter)
	logWarn.SetFormatter(formatter)
	logInfo.SetFormatter(formatter)
	logDebug.SetFormatter(formatter)
	logTrace.SetFormatter(formatter)
}

/******************************/
// 非常规命名的日志文件
/******************************/
//func InitLogger(files []string, logPath string) {
//	logMap := make(map[string]*logrus.Logger)
//	for _, file := range files {
//		NewLogger(logMap, file, logPath)
//	}
//}
//
//func NewLogger(m map[string]*logrus.Logger, level string, logPath string) {
//	logger := logrus.New()
//	hk := newHook(level, logPath)
//	logger.AddHook(hk)
//	m[level] = logger
//}

// initiate watcher of logs
func initFileWatcher(logPath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logFatal.Fatal(err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)

	err = watcher.Add(logPath)
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`^(fatal|error|warn|info|debug|trace)_\d{4}-\d{2}-\d{2}\.log$`)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// whether event is Removing
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					// whether removed file is log
					if reg.Match([]byte(path.Base(event.Name))) {
						updateFileHandle(event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err.Error())
			}
		}
	}()

	<-done
}

// update log's file handle by using logrus' ReplaceHooks when file's removed
func updateFileHandle(filePath string) {
	fileName := path.Base(filePath)
	fileName = strings.Split(fileName, "_")[0]
	logPath := path.Dir(filePath)
	neoHooks := make(logrus.LevelHooks)
	hk := newHook(fileName, logPath)
	neoHooks.Add(hk)
	loggerMap[fileName].ReplaceHooks(neoHooks)
}

// change log's level dynamically
func changeLevel(level int) {
	l := logrus.Level(level)
	logFatal.SetLevel(l)
	logError.SetLevel(l)
	logWarn.SetLevel(l)
	logInfo.SetLevel(l)
	logDebug.SetLevel(l)
	logTrace.SetLevel(l)
}
