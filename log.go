package log

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"os"
	"strings"
)

const contextKey = "logasaurusrex"

func init() {
	logrus.SetOutput(os.Stderr)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.AddHook(&CallersHook{
		LogLevels: []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel},
		CallDepth: 4,
	})
}

func newLogger() logrus.FieldLogger {
	log := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.TextFormatter{},
		Hooks:     logrus.LevelHooks{},
		Level:     logrus.DebugLevel,
	}

	log.Hooks.Add(&CallersHook{
		LogLevels: []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel},
		CallDepth: 4,
	})

	return log
}

func Gin() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := newLogger()
		log := l.WithField("handler", c.HandlerName())
		for _, param := range c.Params {
			log.Data[fmt.Sprintf("urlParam-%s", param.Key)] = param.Value
		}
		c.Request.ParseForm()
		for key, params := range c.Request.Form {
			log.Data[fmt.Sprintf("queryParam-%s", key)] = strings.Join(params, ",")
		}
		c.Set(contextKey, log)
		c.Next()
	}
}

func Context(ctx context.Context) logrus.FieldLogger {
	li := ctx.Value(contextKey)
	if li == nil {
		return newLogger()
	}
	log, ok := li.(logrus.FieldLogger)
	if !ok {
		return newLogger()
	}
	return log
}

type Fields map[string]interface{}

// Implement logrus FieldLogger
func WithField(key string, value interface{}) *logrus.Entry { return logrus.WithField(key, value) }
func WithFields(fields Fields) *logrus.Entry                { return logrus.WithFields(logrus.Fields(fields)) }
func WithError(err error) *logrus.Entry                     { return logrus.WithError(err) }
func Debugf(format string, args ...interface{})             { logrus.Debugf(format, args...) }
func Infof(format string, args ...interface{})              { logrus.Infof(format, args...) }
func Printf(format string, args ...interface{})             { logrus.Printf(format, args...) }
func Warnf(format string, args ...interface{})              { logrus.Warnf(format, args...) }
func Warningf(format string, args ...interface{})           { logrus.Warningf(format, args...) }
func Errorf(format string, args ...interface{})             { logrus.Errorf(format, args...) }
func Fatalf(format string, args ...interface{})             { logrus.Fatalf(format, args...) }
func Panicf(format string, args ...interface{})             { logrus.Panicf(format, args...) }
func Debug(args ...interface{})                             { logrus.Debug(args...) }
func Info(args ...interface{})                              { logrus.Info(args...) }
func Print(args ...interface{})                             { logrus.Print(args...) }
func Warn(args ...interface{})                              { logrus.Warn(args...) }
func Warning(args ...interface{})                           { logrus.Warning(args...) }
func Error(args ...interface{})                             { logrus.Error(args...) }
func Fatal(args ...interface{})                             { logrus.Fatal(args...) }
func Panic(args ...interface{})                             { logrus.Panic(args...) }
func Debugln(args ...interface{})                           { logrus.Debugln(args...) }
func Infoln(args ...interface{})                            { logrus.Infoln(args...) }
func Println(args ...interface{})                           { logrus.Println(args...) }
func Warnln(args ...interface{})                            { logrus.Warnln(args...) }
func Warningln(args ...interface{})                         { logrus.Warningln(args...) }
func Errorln(args ...interface{})                           { logrus.Errorln(args...) }
func Fatalln(args ...interface{})                           { logrus.Fatalln(args...) }
func Panicln(args ...interface{})                           { logrus.Panicln(args...) }
