package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	yellow = "\033[33m"
	green  = "\033[32m"
	blue   = "\033[34m"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func newLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)

	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, green+"DEBUG: ", logger.Flags()),
		info:    log.New(writer, blue+"INFO: ", logger.Flags()),
		warning: log.New(writer, yellow+"WARNING: ", logger.Flags()),
		err:     log.New(writer, red+"ERR: ", logger.Flags()),
		writer:  writer,
	}
}

func GetLogger(prefix string) *Logger {
	return newLogger(prefix)
}

// Create non-formatted logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(fmt.Sprint(v...) + reset)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(fmt.Sprint(v...) + reset)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(fmt.Sprint(v...) + reset)
}

func (l *Logger) Err(v ...interface{}) {
	l.err.Println(fmt.Sprint(v...) + reset)
}

// Create formate enabled logs

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format+reset, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format+reset, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format+reset, v...)
}

func (l *Logger) Errf(format string, v ...interface{}) {
	l.err.Printf(format+reset, v...)
}
