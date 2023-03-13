package loggerservice

import (
	loggermodel "e-commerce/internal/models/loggerModels"
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	logs "github.com/jeanphorn/log4go"
)

type LogSrv interface {
	Info(arg0 any, args ...any)
	Debug(arg0 any, args ...any)
	Warning(arg0 any, args ...any)
	Error(arg0 any, args ...any)
	Fatal(arg0 any, args ...any)
	Audit(record *loggermodel.AuditLog)
}

type logSrv struct {
	logger *logs.Filter
}

func (l logSrv) Info(arg0 any, args ...any) {
	l.logger.Log(logs.INFO, getSource(), fmt.Sprintf(arg0.(string), args...))
}

func (l logSrv) Debug(arg0 any, args ...any) {
	l.logger.Log(logs.DEBUG, getSource(), fmt.Sprintf(arg0.(string), args...))
}

func (l logSrv) Warning(arg0 any, args ...any) {
	l.logger.Log(logs.WARNING, getSource(), fmt.Sprintf(arg0.(string), args...))
}

func (l logSrv) Error(arg0 any, args ...any) {
	l.logger.Log(logs.ERROR, getSource(), fmt.Sprintf(arg0.(string), args...))
}

func (l logSrv) Fatal(arg0 any, args ...any) {
	l.logger.Log(logs.CRITICAL, getSource(), fmt.Sprintf(arg0.(string), args...))
	l.logger.Close()
	os.Exit(1)
}

func (l logSrv) Audit(record *loggermodel.AuditLog) {
	js, _ := json.Marshal(record)
	l.logger.Log(logs.INFO, getSource(), string(js))
}

func NewLogger() LogSrv {
	return &logSrv{logger: logs.LOGGER("fileLogs")}
}

// Auxillary Functions
func getSource() (source string) {
	if pc, _, line, ok := runtime.Caller(2); ok {
		source = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), line)
	}
	return
}
