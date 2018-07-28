package flog

import (
	"fmt"
	"os"
	"time"
)

type Level uint8

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	}
	return ""
}

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func (l *Logger) log(level Level, s string) {
	l.output.Write([]byte(fmt.Sprintf("[%s] %s %s\n", level.String(), time.Now().Format(time.RFC3339), s)))
}

func (l *Logger) Debug(a ...interface{}) {
	if l.level <= DebugLevel {
		l.log(DebugLevel, fmt.Sprint(a...))
	}
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Debug(fmt.Sprintf(format, a...))
}

func (l *Logger) Info(a ...interface{}) {
	if l.level <= InfoLevel {
		l.log(InfoLevel, fmt.Sprint(a...))
	}
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.Info(fmt.Sprintf(format, a...))
}

func (l *Logger) Warn(a ...interface{}) {
	if l.level <= WarnLevel {
		l.log(WarnLevel, fmt.Sprint(a...))
	}
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Warn(fmt.Sprintf(format, a...))
}

func (l *Logger) Error(a ...interface{}) {
	if l.level <= ErrorLevel {
		l.log(ErrorLevel, fmt.Sprint(a...))
	}
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Error(fmt.Sprintf(format, a...))
}

func (l *Logger) Fatal(a ...interface{}) {
	if l.level <= FatalLevel {
		l.log(ErrorLevel, fmt.Sprint(a...))
	}
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Fatal(fmt.Sprintf(format, a...))
}
