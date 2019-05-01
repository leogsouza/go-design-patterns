package chainofresponsibility

import (
	"fmt"
	"io"
	"strings"
)

// ChainLogger is a interface
type ChainLogger interface {
	Next(string)
}

// FirstLogger is a logger with impllements the next chain
type FirstLogger struct {
	NextChain ChainLogger
}

// Next sends the flow message to the next chain
func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

// SecondLogger is the second logger where
// will be redirected the flow of messages
type SecondLogger struct {
	NextChain ChainLogger
}

// Next sends the flow message to the next chain
func (se *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if se.NextChain != nil {
			se.NextChain.Next(s)
		}

		return
	}

	fmt.Printf("Finishing in second logging\n\n")
}

// WriterLogger sends the information to the console
type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

// Next sends the flow message to the next chain
func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}

// ClosureChain is a struct which implements a closure
// And use it in a chain of responsibility
type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

// Next sends the flow message to the next chain
func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}

	if c.NextChain != nil {
		c.Next(s)
	}
}
