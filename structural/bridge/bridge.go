package bridge

import (
	"errors"
	"fmt"
	"io"
)

// PrinterAPI provides a method which prints a message
type PrinterAPI interface {
	PrintMessage(string) error
}

// PrinterImpl1 implements PrinterAPI interface
type PrinterImpl1 struct{}

// PrintMessage Prints a message
func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

// PrinterImpl2 implements a Writer interface
type PrinterImpl2 struct {
	Writer io.Writer
}

// PrintMessage prints a message
func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

// PrinterAbstraction is an abstraction of PrinterAPI
type PrinterAbstraction interface {
	Print() error
}

// NormalPrinter is a struct which will implement the abstraction
type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print prints a message from printer api
func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

// PacktPrinter is a struct which will implement the abstraction
type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print prints a message from printer api
func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}
