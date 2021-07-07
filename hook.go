package flog

import "github.com/sirupsen/logrus"

var (
	NameHookFieldKey = "logger_name"
)

// NameHook
type NameHook struct {
	name string
}

func (h *NameHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *NameHook) Fire(entry *logrus.Entry) error {
	entry.Data[NameHookFieldKey] = h.name
	return nil
}

func NewNameHook(name string) *NameHook {
	return &NameHook{name: name}
}
