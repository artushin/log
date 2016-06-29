package log

import (
	"bytes"
	"github.com/Sirupsen/logrus"
	"runtime"
	"strconv"
)

type CallersHook struct {
	LogLevels []logrus.Level
	CallDepth int
}

func (ch *CallersHook) Levels() []logrus.Level {
	return ch.LogLevels
}

func (ch *CallersHook) Fire(e *logrus.Entry) error {
	calldepth := ch.CallDepth
	if calldepth == 0 {
		calldepth = 4
	}
	buf := &bytes.Buffer{}
	for i := 0; i < calldepth; i++ {
		_, file, line, ok := runtime.Caller(i + 4)
		if ok {
			buf.WriteString(file)
			buf.WriteRune(':')
			buf.WriteString(strconv.Itoa(line))
			buf.WriteRune(' ')
		}
	}
	e.Data["stack"] = buf.String()
	return nil
}
