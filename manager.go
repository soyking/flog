package flog

import "sync"

type Manager struct {
	sync.Mutex
	loggers map[string]*Logger
}

func (m *Manager) GetLogger(name string) *Logger {
	m.Lock()
	defer m.Unlock()

	if l, ok := m.loggers[name]; ok {
		return l
	} else {
		logger := NewLogger()
		m.loggers[name] = logger
		return logger
	}
}

func NewManager() *Manager {
	return &Manager{
		loggers: make(map[string]*Logger),
	}
}
