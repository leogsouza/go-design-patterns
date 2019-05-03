package main

import "fmt"

// Command represents a command which returns a value
type Command interface {
	GetValue() interface{}
}

// Volume represents a volume of audio mixer
type Volume byte

// GetValue return the value of Volume
func (v Volume) GetValue() interface{} {
	return v
}

// Mute informs if the mixer is mute or not
type Mute bool

// GetValue returns the value of Mute type
func (m Mute) GetValue() interface{} {
	return m
}

// Memento stores a command which pointers to a Mute of Volume type
type Memento struct {
	memento Command
}

type originator struct {
	Command Command
}

// NewMemento create a new instance of memento
func (o *originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

// ExtractAndStoreCommand stores the memento into the originator
func (o *originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

type careTaker struct {
	mementoStack []Memento
}

func (c *careTaker) Add(m Memento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *careTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		tempMemento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[:len(c.mementoStack)-1]
		return tempMemento
	}
	return Memento{}
}

// MementoFacade holds the originator and careTaker
type MementoFacade struct {
	originator originator
	careTaker  careTaker
}

// SaveSettings saves the type into hour memento and add into mementoStack
func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Add(m.originator.NewMemento())
}

// RestoreSettings gets the settings according to index into stack
func (m *MementoFacade) RestoreSettings(i int) Command {
	m.originator.ExtractAndStoreCommand(m.careTaker.Pop())
	return m.originator.Command
}

func main() {
	m := MementoFacade{}

	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))

	assertAndPrint(m.RestoreSettings(0))
	assertAndPrint(m.RestoreSettings(1))
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
