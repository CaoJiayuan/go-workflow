package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"os"
	"path/filepath"
	"time"
)

var (
	maxAge = rotatelogs.WithMaxAge(time.Duration(30*86400) * time.Second)
	// 每小时
	rotationTime = rotatelogs.WithRotationTime(time.Hour)

	DefaultLogger = GetLogger()
)

func GetLogger(name ...string) *logrus.Logger {
	l := logrus.New()

	n := "runtime"
	if len(name) > 0 {
		l.WithField("name", name[0])
		n = name[0]
	}
	if w, e := getWriter(n); e == nil {
		l.AddHook(&writer.Hook{
			Writer: w,
			LogLevels: []logrus.Level{
				logrus.InfoLevel,
				logrus.DebugLevel,
				logrus.TraceLevel,
				logrus.WarnLevel,
				logrus.ErrorLevel,
				logrus.PanicLevel,
			},
		})
		l.SetOutput(nilWriter{})
	} else {
		fmt.Print(e)
	}

	return l
}

func root() string {
	binaryRootPath, _ := filepath.Abs(os.Args[0])
	return filepath.Dir(binaryRootPath)
}

func getWriter(name string) (*rotatelogs.RotateLogs, error) {
	path := fmt.Sprintf("%s/%s-", root()+"/logs", name) + "%Y%m%d.log"

	return rotatelogs.New(path, maxAge, rotationTime)
}

type nilWriter struct {
}

func (nilWriter) Write(_ []byte) (n int, err error) {
	return 0, nil
}

func Field(key string, value interface{}) *logrus.Entry {
	return DefaultLogger.WithField(key, value)
}

func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}

func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}
