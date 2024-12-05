package logger

import (
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net"
	"sync"
	"time"
)

var (
	i    *impl
	s    *Logger
	once sync.Once
)

type impl struct {
	conn net.Conn
}

func (l *impl) Write(p []byte) (int, error) {
	return l.conn.Write(p)
}

type Logger struct {
	logger         *zap.Logger
	constantFields map[Field]interface{}
}

func Log() *Logger {
	return s
}

func New(endpoint string, timeout time.Duration, level Level) (l *Logger, err error) {
	once.Do(func() {

		i = &impl{}
		i.conn, err = net.DialTimeout("udp", endpoint, timeout)
		if err != nil {
			logrus.Warnln("can not connect to elk, error:", err.Error())
			s, err = NewStdOut()
			if err != nil {
				log.Fatalln(err.Error())
			}

			return
		}

		encoderConfig := ecszap.NewDefaultEncoderConfig()
		prodEC := zap.NewProductionEncoderConfig()
		prodEC.EncodeTime = zapcore.RFC3339TimeEncoder
		core := zapcore.NewTee(ecszap.NewCore(encoderConfig, zapcore.AddSync(i), level.toZapCoreLevel()))
		logger := zap.New(core)

		s = &Logger{
			logger: logger,
		}
	})

	return s, nil
}

func NewStdOut() (l *Logger, err error) {
	once.Do(func() {

		logger, err := zap.NewDevelopment()
		if err != nil {
			return
		}

		s = &Logger{
			logger: logger,
		}
	})

	return s, nil
}

func (l *Logger) addField(key Field, value interface{}) {
	l.constantFields[key] = value
}

func (l *Logger) addFieldsTo(logs *Logs) {
	for k, v := range l.constantFields {
		logs.addField(k, v)
	}
}
