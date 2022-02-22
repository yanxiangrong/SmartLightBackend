package logging

import (
	"fmt"
	"github.com/TwiN/go-color"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{color.Ize(color.Cyan, "DEBUG"), "INFO",
		color.Ize(color.Yellow, "WARN"),
		color.Ize(color.Red, "ERROR"),
		color.Ize(color.Purple, "FATAL")}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	logger = log.New(os.Stdout, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
	logger.Println(color.Cyan, v, color.Reset)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(color.Yellow, v, color.Reset)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(color.Red, v, color.Reset)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(color.Purple, v, color.Reset)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	//logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])

	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d] ", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
