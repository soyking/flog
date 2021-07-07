package flog

var global = NewLogger("global")

func GetLogger(name string) *Logger {
	return global.GetLogger(name)
}

func Setup(setup func(*Logger), names ...string) {
	global.Setup(setup, names...)
}
