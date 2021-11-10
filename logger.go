package logcomm

import (
	"log"
	"path"
	"regexp"

	"github.com/sirupsen/logrus"

	"github.com/fsnotify/fsnotify"
)

// initiate watcher of logs
func initFileWatcher(logger *logrus.Logger, logPath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("create watcher failed: " + err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)

	err = watcher.Add(logPath)
	if err != nil {
		panic("add file to watcher failed: " + err.Error())
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
						updateFileHandle(logger, logPath)
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

func updateFileHandle(logger *logrus.Logger, logPath string) {
	hk := newSeperatedHook(logPath)
	neoHooks := make(logrus.LevelHooks)
	neoHooks.Add(hk)
	logger.ReplaceHooks(neoHooks)
}
