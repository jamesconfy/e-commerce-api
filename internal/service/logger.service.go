package service

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type LogSrv interface {
	Info(message string)
	Debug(message string)
	Warning(message string)
	Error(message string)
	Fatal(message string)
}

type logSrv struct {
	logger *log.Logger
}

func (l logSrv) Info(message string) {
	l.logger.Log(log.InfoLevel, getSource(), fmt.Sprintf(" %s", message))
}

func (l logSrv) Debug(message string) {
	l.logger.Log(log.DebugLevel, getSource(), fmt.Sprintf(" %s", message))
}

func (l logSrv) Warning(message string) {
	l.logger.Log(log.WarnLevel, getSource(), fmt.Sprintf(" %s", message))
}

func (l logSrv) Error(message string) {
	l.logger.Log(log.ErrorLevel, getSource(), fmt.Sprintf(" %s", message))
}

func (l logSrv) Fatal(message string) {
	l.logger.Log(log.FatalLevel, getSource(), fmt.Sprintf(" %s", message))
}

func NewLoggerService(fileName string) LogSrv {
	logger := log.New()
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// log.SetOutput(os.Stdout)
	if fileName == "" {
		fileName = "./logs/logrus.log"
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	return &logSrv{logger: logger}
}

// Auxillary Functions
func getSource() (source string) {
	if pc, _, _, ok := runtime.Caller(2); ok {
		str := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		source = fmt.Sprintf("%s():", str[len(str)-1])
	}
	return
}
