package logger

import (
	"context"
	"fmt"
	"github.com/sinameshkini/microkit/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type Logs struct {
	ctx    context.Context
	level  Level
	fields map[Field]interface{}
	logger *zap.Logger
}

func (l *Logs) addField(key Field, value interface{}) {
	l.fields[key] = value
}

func (l *Logs) zapFields() (zf []zapcore.Field) {
	for k, v := range l.fields {
		zf = append(zf, zap.Any(k.String(), v))
	}

	return
}

func (l *Logs) Commit(message string) *Logs {
	l.addField(FieldCaller, getCaller(2))
	l.addField(FieldSourceCaller, getCaller(3))

	switch l.level {
	case InfoLevel:
		l.logger.Info(message, l.zapFields()...)
	case FatalLevel:
		l.logger.DPanic(message, l.zapFields()...)
	case WarnLevel:
		l.logger.Warn(message, l.zapFields()...)
	case DebugLevel:
		l.logger.Debug(message, l.zapFields()...)
	case ErrorLevel:
		l.logger.Error(message, l.zapFields()...)
	}

	defer func() {
		if err := l.logger.Sync(); err != nil {
			log.Println(err.Error())
		}
	}()
	return l
}

func (l *Logs) With(key Field, value any) *Logs {
	l.addField(key, value)

	return l
}

func (l *Logger) With(key Field, value any) *Logger {
	l.addField(key, value)

	return l
}

func (l *Logs) WithDBQuery(maskedQuery string) *Logs {
	l.addField(FieldDBQuery, maskedQuery)
	return l
}

func (l *Logs) WithDBRowsAffected(rowsCount int) *Logs {
	l.addField(FieldDBRowAffected, rowsCount)
	return l
}

func (l *Logs) WithDBDelay(delayMilliSec int64) *Logs {
	l.addField(FieldDBDelay, delayMilliSec)
	return l
}

func (l *Logs) WithError(err error) *Logs {
	if err != nil {
		l.addField(FieldError, err.Error())
	}

	return l
}

func (l *Logs) WithMethod(method string) *Logs {
	l.With(FieldMethodName, method)

	return l
}

func (l *Logs) WithInput(input any) *Logs {
	l.With(FieldMethodInput, utils.JsonIndent(input))

	return l
}

func (l *Logs) WithUserError(err string) *Logs {
	if err != "" {
		l.addField(FieldUserError, err)
	}

	return l
}

func (l *Logs) WithErrorCode(err string) *Logs {
	if err != "" {
		l.addField(FieldErrorCode, err)
	}

	return l
}

func getCaller(skip int) string {
	pc, file, line, ok := runtime.Caller(skip) // 1 indicates to get the information of the calling function
	if !ok {
		log.Println("failed to get runtime.Caller, skip 1")
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	return fmt.Sprintf("%s: %d %s\n", file, line, fnName)
}
