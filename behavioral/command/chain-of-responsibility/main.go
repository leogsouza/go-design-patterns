package main

import (
	"fmt"
	"time"
)

// Command interface shows the info
type Command interface {
	Info() string
}

// TimePassed holds the time passed of creation of command
type TimePassed struct {
	start time.Time
}

// Info returns how many time has passed
func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

// ChainLogger is a interface
type ChainLogger interface {
	Next(Command)
}

// Logger is a struct which represents a Chain of loggers
type Logger struct {
	NextChain ChainLogger
}

// Next sends the flow message to the next chain
func (f *Logger) Next(c Command) {
	time.Sleep(time.Second)

	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}

func main() {
	second := new(Logger)
	first := Logger{NextChain: second}

	command := &TimePassed{start: time.Now()}

	first.Next(command)
}
