package apilog

import "fmt"

type stdoutWriter struct {
}

func NewStdoutWriter() *stdoutWriter {
	return &stdoutWriter{}
}

func (self *stdoutWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}

func (self *stdoutWriter) Sync() error {
	return nil
}
