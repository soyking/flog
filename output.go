package flog

import (
	"io"
	"sync"
)

type LockOutput struct {
	sync.Mutex
	output io.Writer
}

func (l *LockOutput) Write(p []byte) (int, error) {
	l.Lock()
	defer l.Unlock()

	return l.output.Write(p)
}

func NewLockOutput(output io.Writer) *LockOutput {
	return &LockOutput{
		output: output,
	}
}
