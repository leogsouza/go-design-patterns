package main

import (
	"fmt"
)

// Command is a generic interface which exposes Execute method
type Command interface {
	Execute()
}

// ConsoleOutput implements Command interface
// And print to the console the member called message
type ConsoleOutput struct {
	message string
}

// Execute prints the ConsoleOutput message to the console
func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

// CreateCommand is a Command constructor
func CreateCommand(s string) Command {
	fmt.Println("Creating command")

	return &ConsoleOutput{
		message: s,
	}
}

// CommandQueue store in a queue a sequence of commands
// to be executed when the queue reach the required length.
type CommandQueue struct {
	queue []Command
}

// AddCommand adds the command to the queue
// The commands will be executed when the queue reach the limit
func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}

		p.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First Message"))
	queue.AddCommand(CreateCommand("Second Message"))
	queue.AddCommand(CreateCommand("Third Message"))

	queue.AddCommand(CreateCommand("Fourth Message"))
	queue.AddCommand(CreateCommand("Fifth Message"))
}
