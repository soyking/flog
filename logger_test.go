package flog

import (
	"os"
	"testing"
)

func TestSubLogger(t *testing.T) {
	l := NewLogger()
	logger1 := l.GetLogger("1")
	logger2 := l.GetLogger("2")

	l.SetLoggerLevel("1", FatalLevel)
	if logger1.level != FatalLevel || logger2.level != InfoLevel {
		t.Fatal("level err")
	}
	l.SetLevel(DebugLevel)
	if logger1.level != DebugLevel || logger2.level != DebugLevel {
		t.Fatal("level err")
	}

	l.SetLoggerOutput("1", os.Stderr)
	if logger1.output != os.Stderr || logger2.output != os.Stdout {
		t.Fatal("output err")
	}
	l.SetOutput(os.Stdin)
	if logger1.output != os.Stdin || logger2.output != os.Stdin {
		t.Fatal("output err")
	}
}
