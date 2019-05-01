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

// HelloMessage is a struct which represents a greeting message
type HelloMessage struct{}

// Info just returns a Hello World message
func (h HelloMessage) Info() string {
	return "Hello world!"
}

func main() {
	var timeCommand Command
	timeCommand = &TimePassed{time.Now()}

	var helloCommand Command
	helloCommand = &HelloMessage{}

	time.Sleep(time.Second)
	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}
