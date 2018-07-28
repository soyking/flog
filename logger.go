package flog

import (
	"io"
	"os"
	"sync"
)

const (
	AllLoggers = ""
)

type Logger struct {
	sync.Mutex
	loggers map[string]*Logger

	output io.Writer // TODO: output lock
	level  Level
}

func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
	l.SetLoggerOutput(AllLoggers, output)
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
	l.SetLoggerLevel(AllLoggers, level)
}

func (l *Logger) GetLogger(name string) *Logger {
	l.Lock()
	defer l.Unlock()

	if name == AllLoggers {
		panic("invalid logger name")
	}

	if logger, ok := l.loggers[name]; ok {
		return logger
	} else {
		logger := NewLogger()
		logger.SetOutput(l.output)
		logger.SetLevel(l.level)
		l.loggers[name] = logger
		return logger
	}
}

func (l *Logger) SetLoggerOutput(name string, output io.Writer) {
	l.Lock()
	defer l.Unlock()

	if name == AllLoggers {
		for _, logger := range l.loggers {
			logger.SetOutput(output)
		}
	} else if logger, ok := l.loggers[name]; ok {
		logger.SetOutput(output)
	}
}

func (l *Logger) SetLoggerLevel(name string, level Level) {
	l.Lock()
	defer l.Unlock()

	if name == AllLoggers {
		for _, logger := range l.loggers {
			logger.SetLevel(level)
		}
	} else if logger, ok := l.loggers[name]; ok {
		logger.SetLevel(level)
	}
}

func NewLogger() *Logger {
	return &Logger{
		loggers: make(map[string]*Logger),
		output:  os.Stdout,
		level:   InfoLevel,
	}
}
