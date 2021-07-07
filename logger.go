package flog

import (
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	AllLoggers     = "__all__loggers__"
	NameSeparators = "."
)

type Logger struct {
	*logrus.Logger
	name string

	lock     sync.Mutex
	children map[string]*Logger
}

func (l *Logger) GetLogger(name string) *Logger {
	l.lock.Lock()
	defer l.lock.Unlock()

	if name == AllLoggers {
		panic("invalid logger name")
	}

	if child, ok := l.children[name]; ok {
		return child
	} else {
		fullname := strings.Join([]string{l.name, name}, NameSeparators)
		child := NewLogger(fullname)
		// copy features from parent logger
		child.SetOutput(l.Out)
		child.SetFormatter(l.Formatter)
		child.SetReportCaller(l.ReportCaller)
		child.SetLevel(l.Level)
		child.ExitFunc = l.ExitFunc

		// hack for hooks
		hooksMap := make(map[logrus.Hook]struct{})
		for _, hooks := range l.Hooks {
			for _, hook := range hooks {
				if _, ok := hooksMap[hook]; !ok {
					if _, ok := hook.(*NameHook); !ok {
						child.AddHook(hook)
					}
					hooksMap[hook] = struct{}{}
				}
			}
		}
		child.AddHook(NewNameHook(fullname))

		l.children[name] = child
		return child
	}
}

// Use Setup to update Logger and its children's features
func (l *Logger) Setup(setup func(*Logger), names ...string) {
	if len(names) == 1 && names[0] == AllLoggers {
		l.lock.Lock()
		defer l.lock.Unlock()
		setup(l)
		for _, child := range l.children {
			child.Setup(setup, AllLoggers)
		}
	} else {
		child := l
		for _, name := range names {
			child = child.GetLogger(name)
		}

		l.lock.Lock()
		defer l.lock.Unlock()
		setup(child)
	}
}

func NewLogger(name string) *Logger {
	l := logrus.New()
	l.AddHook(NewNameHook(name))
	return &Logger{
		Logger:   l,
		name:     name,
		children: make(map[string]*Logger),
	}
}
