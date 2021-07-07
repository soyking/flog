package flog

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	tLog := NewLogger("tLog")

	out := bytes.NewBuffer(nil)
	assertOut := func(f func(outS string) bool) {
		outS := out.String()
		out.Reset()
		t.Logf("out log: %s", outS)
		if !f(outS) {
			t.Fatal("check failed")
		}
	}

	tLog.Setup(func(l *Logger) {
		l.SetOutput(out)
		l.SetLevel(logrus.ErrorLevel)
	})
	tLog.Error("log in global")
	assertOut(func(outS string) bool { return strings.Contains(outS, "log in global") })

	module1Log := tLog.GetLogger("module1")
	// inherit log level from parent
	module1Log.Info("log in module1")
	assertOut(func(outS string) bool { return outS == "" })
	module1Log.Error("log in module1")
	assertOut(func(outS string) bool {
		return strings.Contains(outS, "log in module1") && strings.Contains(outS, "logger_name=tLog.module1")
	})
	// change log level from parent
	tLog.Setup(func(l *Logger) {
		l.SetLevel(logrus.InfoLevel)
	}, "module1")
	module1Log.Info("log in module1")
	assertOut(func(outS string) bool { return strings.Contains(outS, "log in module1") })
	module1Log.Debug("log in module1")
	assertOut(func(outS string) bool { return outS == "" })
	// change own log level
	module1Log.Setup(func(l *Logger) { l.SetLevel(logrus.DebugLevel) })
	module1Log.Debug("log in module1")
	assertOut(func(outS string) bool { return strings.Contains(outS, "log in module1") })

	// change all loggers
	tLog.Setup(func(l *Logger) {
		l.SetLevel(logrus.PanicLevel)
	}, AllLoggers)
	tLog.Error("log in global")
	assertOut(func(outS string) bool { return outS == "" })
	module1Log.Debug("log in module1")
	assertOut(func(outS string) bool { return outS == "" })
}
