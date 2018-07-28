package flog

import "io"

type Logger struct {
	writer io.Writer
	level  uint32
}

func (l *Logger) Debug(...interface{}) {
	panic("implement me")
}

func (l *Logger) Debugln(...interface{}) {
	panic("implement me")
}

func (l *Logger) Debugf(string, ...interface{}) {
	panic("implement me")
}

func (l *Logger) Info(...interface{}) {
	panic("implement me")
}

func (l *Logger) Infoln(...interface{}) {
	panic("implement me")
}

func (l *Logger) Infof(string, ...interface{}) {
	panic("implement me")
}

func (l *Logger) Warn(...interface{}) {
	panic("implement me")
}

func (l *Logger) Warnln(...interface{}) {
	panic("implement me")
}

func (l *Logger) Warnf(string, ...interface{}) {
	panic("implement me")
}

func (l *Logger) Error(...interface{}) {
	panic("implement me")
}

func (l *Logger) Errorln(...interface{}) {
	panic("implement me")
}

func (l *Logger) Errorf(string, ...interface{}) {
	panic("implement me")
}

func (l *Logger) Fatal(...interface{}) {
	panic("implement me")
}

func (l *Logger) Fatalln(...interface{}) {
	panic("implement me")
}

func (l *Logger) Fatalf(string, ...interface{}) {
	panic("implement me")
}

func NewLogger()*Logger {
	return &Logger{}
}
