# logcomm

Logcomm is  a logger for GO, based on Logrus, File-rotatelogs and Lfshook.



## Install

```sh
$ go get ...
```



## Usage

It provides two ways of logging:

* print in single file by using NewMergedLogger(logPath string, logName string)
* print in seperated files of each level with rotation by using NewSeperatedLogger(logPath string)

Logger that function returns is a logrus.Logger, you can using logrus' API to do everything you want to.

```go
const (
	Panic = iota
	Fatal
	Error
	Warn
	Info
	Debug
	Trace
)

path := "/data/test/log"
// return a mergedLogger
logMerged := logcomm.NewMergedLogger(path, "merge")
// return a seperatedLogger
logSeperated := logcomm.NewSeperatedLogger(path)

logMerged.Logf(Error, "this is fatal %d:%d:%d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
logMerged.Logf(Fatal, "this is fatal %d:%d:%d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
// 
```



## Maintainers

[@fAROSIA](https://github.com/fAROSIA)