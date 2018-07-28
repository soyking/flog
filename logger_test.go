package flog

import (
	"bytes"
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

func TestLevel(t *testing.T) {
	l := NewLogger()
	output := bytes.NewBuffer(nil)
	l.SetOutput(output)

	l.Debug("Debug")
	if output.Len() > 0 {
		t.Fatal("should not output")
	}
	l.Info("Info")
	if output.Len() == 0 {
		t.Fatal("should output")
	} else {
		t.Log(output.String())
	}

	l.SetLevel(ErrorLevel)
	output.Reset()
	l.Info("Info")
	if output.Len() > 0 {
		t.Fatal("should not output")
	}
	l.Error("Error")
	if output.Len() == 0 {
		t.Fatal("should output")
	} else {
		t.Log(output.String())
	}
}

func TestLoggerLevel(t *testing.T) {
	l := NewLogger()
	logger1 := l.GetLogger("1")

	l.SetLevel(ErrorLevel)
	l.SetLoggerLevel("1", InfoLevel)

	output := bytes.NewBuffer(nil)
	l.SetOutput(output)

	l.Info("Info")
	if output.Len() > 0 {
		t.Fatal("should not output")
	}
	l.Error("Error")
	if output.Len() == 0 {
		t.Fatal("should output")
	} else {
		t.Log(output.String())
	}

	output.Reset()
	logger1.Info("Info")
	if output.Len() == 0 {
		t.Fatal("should output")
	} else {
		t.Log(output.String())
	}
}
