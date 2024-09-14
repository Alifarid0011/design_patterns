package main

import "io"

type ChainLogger interface {
	Next(string)
}
type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {}

type SecondLogger struct {
	NextChain ChainLogger
}

func (f *SecondLogger) Next(s string) {}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {}

type myTestWriter struct {
	receivedMessage string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	m.receivedMessage += string(p)
	return len(p), nil
}
func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}
